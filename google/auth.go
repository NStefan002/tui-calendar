package google

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func tokenCacheFile() (string, error) {
	path := os.Getenv("GOOGLE_TOKEN_CACHE")
	if path == "" {
		return "", fmt.Errorf("GOOGLE_TOKEN_CACHE not set in .env or environment")
	}

	// Expand ~ to home directory
	if path[:2] == "~/" {
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

	defer func() {
		if cerr := f.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "Error closing file %s: %v\n", path, cerr)
		}
	}()

	return json.NewEncoder(f).Encode(token)
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer func() {
		if cerr := f.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "Error closing file %s: %v\n", file, cerr)
		}
	}()

	token := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(token)
	return token, err
}

func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token",
		oauth2.AccessTypeOffline,
		oauth2.SetAuthURLParam("redirect_uri", "urn:ietf:wg:oauth:2.0:oob"))

	fmt.Printf("Go to the following URL in your browser, authorize the app, and paste the code below:\n%v\n\n", authURL)

	var code string
	fmt.Print("Paste authorization code here: ")
	if _, err := fmt.Scan(&code); err != nil {
		return nil, err
	}

	token, err := config.Exchange(
		context.Background(),
		code,
		oauth2.SetAuthURLParam("redirect_uri", "urn:ietf:wg:oauth:2.0:oob"),
	)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetClient() (*calendar.Service, error) {
	ctx := context.Background()

	credPath := os.Getenv("GOOGLE_CREDENTIALS_PATH")
	if credPath == "" {
		return nil, fmt.Errorf("GOOGLE_CREDENTIALS_PATH not set in .env or environment")
	}

	// load credentials from the specified path
	b, err := os.ReadFile(credPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read client secret file at %s: %v", credPath, err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope) // use full read/write access
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %v", err)
	}

	tokFile, err := tokenCacheFile()
	if err != nil {
		return nil, err
	}

	var tok *oauth2.Token
	tok, err = tokenFromFile(tokFile)
	if err != nil {
		tok, err = getTokenFromWeb(config)
		if err != nil {
			return nil, err
		}
		_ = saveToken(tokFile, tok)
	}

	client := config.Client(ctx, tok)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to create calendar service: %v", err)
	}
	return srv, nil
}
