package google

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var redirectURL = "http://localhost:8888/callback"

func tokenCacheFile() (string, error) {
	path := os.Getenv("GOOGLE_TOKEN_CACHE")
	if path == "" {
		return "", fmt.Errorf("GOOGLE_TOKEN_CACHE not set")
	}

	// expand `~` to home directory
	if len(path) >= 2 && path[:2] == "~/" {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		path = filepath.Join(usr.HomeDir, path[2:])
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", err
	}

	return path, nil
}

func saveToken(path string, token *oauth2.Token) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	encoderErr := json.NewEncoder(f).Encode(token)

	err = f.Close()
	if err != nil {
		return err
	}

	return encoderErr
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	var token oauth2.Token
	decoderErr := json.NewDecoder(f).Decode(&token)

	err = f.Close()
	if err != nil {
		return nil, err
	}

	return &token, decoderErr
}

func getOAuthConfig() (*oauth2.Config, error) {
	credPath := os.Getenv("GOOGLE_CREDENTIALS_PATH")
	if credPath == "" {
		return nil, fmt.Errorf("GOOGLE_CREDENTIALS_PATH not set")
	}

	b, err := os.ReadFile(credPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read credentials: %v", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse credentials: %v", err)
	}

	config.RedirectURL = redirectURL
	return config, nil
}

// starts the local OAuth server and waits for token exchange
func waitForWebLogin(config *oauth2.Config) (*oauth2.Token, error) {
	codeCh := make(chan string)
	mux := http.NewServeMux()

	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		pageBytes, err := os.ReadFile("assets/login_successful.html")
		if err != nil {
			http.Error(w, "Login successful, but failed to load page.", http.StatusInternalServerError)
			log.Printf("failed to load login success page: %v", err)
		} else {
			if _, err := w.Write(pageBytes); err != nil {
				log.Printf("failed to write response: %v", err)
			}
		}

		codeCh <- code
	})

	ln, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		return nil, fmt.Errorf("failed to start local server: %v", err)
	}

	server := &http.Server{Handler: mux}

	// serve in background
	go func() {
		if err := server.Serve(ln); err != nil && err != http.ErrServerClosed {
			log.Printf("server error: %v", err)
		}
	}()

	// open browser
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	if err := openBrowser(authURL); err != nil {
		return nil, fmt.Errorf("failed to open browser: %v", err)
	}

	// wait for code
	code := <-codeCh

	// shutdown server
	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("server shutdown error: %v", err)
	}

	// exchange code for token
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("token exchange failed: %v", err)
	}

	return token, nil
}

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "darwin":
		cmd, args = "open", []string{url}
	case "linux":
		cmd, args = "xdg-open", []string{url}
	case "windows":
		cmd, args = "rundll32", []string{"url.dll,FileProtocolHandler", url}
	default:
		return fmt.Errorf("unsupported platform")
	}

	return exec.Command(cmd, args...).Start()
}

func GetClient() (*calendar.Service, error) {
	ctx := context.Background()

	tokFile, err := tokenCacheFile()
	if err != nil {
		return nil, err
	}

	config, err := getOAuthConfig()
	if err != nil {
		return nil, err
	}

	var tok *oauth2.Token
	tok, err = tokenFromFile(tokFile)
	if err != nil {
		// generate new token
		tok, err = waitForWebLogin(config)
		if err != nil {
			return nil, err
		}
		err = saveToken(tokFile, tok)
		if err != nil {
			return nil, err
		}
	}

	client := config.Client(ctx, tok)
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to create calendar service: %v", err)
	}
	return srv, nil
}
