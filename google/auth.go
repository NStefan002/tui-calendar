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
	"path/filepath"
	"runtime"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const redirectURL = "http://localhost:8888/callback"

func AppConfigDir() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(dir, "tui-calendar")

	if err := os.MkdirAll(path, 0700); err != nil {
		return "", err
	}

	return path, nil
}

func tokenCacheFile() (string, error) {
	cfgDir, err := AppConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(cfgDir, "token.json"), nil
}

func saveToken(path string, token *oauth2.Token) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(token)
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var token oauth2.Token
	if err := json.NewDecoder(f).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}

func getOAuthConfig() (*oauth2.Config, error) {
	cfgDir, err := AppConfigDir()
	if err != nil {
		return nil, err
	}

	credPath := filepath.Join(cfgDir, "credentials.json")

	b, err := os.ReadFile(credPath)
	if err != nil {
		fmt.Println("First time running tui-calendar? Run `tui-calendar init`.")
		return nil, fmt.Errorf("unable to read credentials file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse credentials: %v", err)
	}

	config.RedirectURL = redirectURL
	return config, nil
}

func waitForWebLogin(config *oauth2.Config) (*oauth2.Token, error) {
	codeCh := make(chan string)
	mux := http.NewServeMux()

	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		pageBytes, err := os.ReadFile("assets/login_successful.html")
		if err != nil {
			http.Error(w, "Login successful, but failed to load local HTML.", http.StatusInternalServerError)
		} else {
			w.Write(pageBytes)
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
	fmt.Printf("Opening browser for authentication. If it does not open, please visit the following URL:\n%s\n", authURL)
	if err := openBrowser(authURL); err != nil {
		return nil, fmt.Errorf("failed to open browser: %v", err)
	}

	// wait for code
	code := <-codeCh

	// shutdown server
	if err := server.Shutdown(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to shutdown server: %v", err)
	}

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

	if tok, err = tokenFromFile(tokFile); err != nil {
		tok, err = waitForWebLogin(config)
		if err != nil {
			return nil, err
		}
		if err := saveToken(tokFile, tok); err != nil {
			return nil, err
		}
	}

	client := config.Client(ctx, tok)

	return calendar.NewService(ctx, option.WithHTTPClient(client))
}
