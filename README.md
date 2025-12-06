# tui-calendar

A simple terminal calendar app with Google Calendar sync.

<!-- ## Screenshots -->

## Installation

### Using Go CLI

```bash
go install github.com/NStefan002/tui-calendar@latest
```

### Building from Source

```bash
git clone https://github.com/NStefan002/tui-calendar.git
cd tui-calendar
go build -o tui-calendar .
# to run
./tui-calendar
```

## First-time Setup (Google OAuth)

<!-- prettier-ignore -->
> [!NOTE]
> This step is only required the first time you run `tui-calendar`, after that your credentials will be saved
> locally and you will never have to do this again.

`tui-calendar` uses your Google Calendar, so the first time you run it you need to provide your own Google OAuth
credentials. It only takes a minute and you do it once.

1. Create your OAuth credentials
   1. Go to: <https://console.cloud.google.com/apis/credentials>
   2. Create (or select) a project.
   3. On the left, open OAuth consent screen → set it to External → add yourself as a Test User → Save.
   4. Go back to Credentials → Create Credentials → OAuth client ID.
   5. Choose Desktop App and create it.
   6. Download the JSON file.

2. Save the downaloaded JSON file as `credentials.json` in the appropriate config directory for your OS:
   - Linux: `~/.config/tui-calendar/credentials.json`
   - macOS: `~/Library/Application Support/tui-calendar/credentials.json`
   - Windows: `C:\Users\<you>\AppData\Roaming\tui-calendar\credentials.json`
3. Run `tui-calendar`. The app will open a browser window for Google login, then store your token so you won’t have to
   log in again.
