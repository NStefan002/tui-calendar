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

Run `tui-calendar init` and follow the instructions to set up Google OAuth credentials. When you finish the setup, your
credentials will be saved in `~/.config/tui-calendar/credentials.json`. Then you can run `tui-calendar` to login and
start using the app.
