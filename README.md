# tui-calendar

A simple terminal calendar app with Google Calendar sync.

<!-- ## Screenshots -->

## Installation

### Using Go CLI

```bash
go install github.com/NStefan002/tui-calendar/v2@latest
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

- Run `tui-calendar init` and follow the instructions to set up Google OAuth credentials
- When you finish the setup, your credentials will be saved in one of the following locations depending on your OS:
  - Linux: `~/.config/tui-calendar/credentials.json`
  - macOS: `~/Library/Application Support/tui-calendar/credentials.json`
  - Windows: `C:\Users\<you>\AppData\Roaming\tui-calendar\credentials.json`

- Run `tui-calendar`, login, and start using the app
