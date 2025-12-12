# Changelog

## [2.1.0](https://github.com/NStefan002/tui-calendar/compare/v2.0.1...v2.1.0) (2025-12-12)


### Features

* center events header and fill entire width ([77c8ad2](https://github.com/NStefan002/tui-calendar/commit/77c8ad269d6751619378bd8ff067a41dec5ad327))

## [2.0.1](https://github.com/NStefan002/tui-calendar/compare/v2.0.0...v2.0.1) (2025-12-07)


### Bug Fixes

* update `go.sum` ([935cc9a](https://github.com/NStefan002/tui-calendar/commit/935cc9a0fbae2eec7ac96d6361b318fe3275688c))

## [2.0.0](https://github.com/NStefan002/tui-calendar/compare/v1.0.2...v2.0.0) (2025-12-07)


### ⚠ BREAKING CHANGES

* ask each user to create their own credentials file

### Features

* add `init` subcommand with detailed instructions ([0777e94](https://github.com/NStefan002/tui-calendar/commit/0777e94750d800c792b552d314bc60774ee8fddb))
* ask each user to create their own credentials file ([b0e34dd](https://github.com/NStefan002/tui-calendar/commit/b0e34dd979cb0bd37a352ee173f5c85f6f31ec66))


### Bug Fixes

* print credentials help to `stdout` ([ac534ca](https://github.com/NStefan002/tui-calendar/commit/ac534ca5c99229e1a45223623ad05ba4360eaaa1))

## [1.0.2](https://github.com/NStefan002/tui-calendar/compare/v1.0.1...v1.0.2) (2025-12-06)


### Bug Fixes

* import path ([c8afed5](https://github.com/NStefan002/tui-calendar/commit/c8afed51182ad1e4a4ad8fbeab7ceed281cafb35))

## [1.0.1](https://github.com/NStefan002/tui-calendar/compare/v1.0.0...v1.0.1) (2025-12-06)


### Bug Fixes

* update project name in go.mod ([a0882d3](https://github.com/NStefan002/tui-calendar/commit/a0882d35fe381e6c9bd1dc7e654289b71592a7c9))

## 1.0.0 (2025-12-04)


### Features

* add help text for keymaps ([ed61b80](https://github.com/NStefan002/tui-calendar/commit/ed61b801aee0f74a35c9a0df94779e7c7d4d7adb))
* **am:** add description field to form ([52caaf0](https://github.com/NStefan002/tui-calendar/commit/52caaf06cdf6946421a77fa04e141576355b6071))
* **auth:** add Google authentication support and improve events display ([bfcc91d](https://github.com/NStefan002/tui-calendar/commit/bfcc91d5c34e7ca57f51aaf4c768426558b7a87d))
* **auth:** handle google login on a local server ([42124c0](https://github.com/NStefan002/tui-calendar/commit/42124c0da1f17e4065d7ccb110a1b538cda35dc2))
* **calendar view:** use different style for dates that contain events ([415ec8f](https://github.com/NStefan002/tui-calendar/commit/415ec8f5a698687bdf0f36610be3b92b21fd20d8))
* **calendar:** fetch events in range of +-30 years of current date ([368e6b7](https://github.com/NStefan002/tui-calendar/commit/368e6b717c86b09393795ae788a38dd6d6acfe84))
* delete event functionality ([f6bc082](https://github.com/NStefan002/tui-calendar/commit/f6bc082c3852de7559f4689c7636ec86fff4890a))
* event creation in add event view ([f53eb60](https://github.com/NStefan002/tui-calendar/commit/f53eb60da418ac65542d6f19caeacb17cb317fce))
* **event details view:** update footer ([695fe6f](https://github.com/NStefan002/tui-calendar/commit/695fe6ff2106aaf5c4aea8bd0b8fa1feace889b0))
* **event details:** display event location if available ([40b18f4](https://github.com/NStefan002/tui-calendar/commit/40b18f42008f02c81c235e60175aa96881ef963d))
* implement edit event functionality ([240b11b](https://github.com/NStefan002/tui-calendar/commit/240b11bf78135facd2327eae2a0b8666ba50c0e6))
* redirect logger output to a file ([18aeda8](https://github.com/NStefan002/tui-calendar/commit/18aeda8daaf312dbf2ddefe2029c91db4aa536e0))
* **styles:** new better-looking styles ([d793c78](https://github.com/NStefan002/tui-calendar/commit/d793c783e0833c365fe605d6d43899b7e6ae0693))
* support showing full help ([64b2e53](https://github.com/NStefan002/tui-calendar/commit/64b2e5377daf402b1d9edef573f477f9c6bb043e))
* **ui:** add spinner when loading calendar ([95ff36b](https://github.com/NStefan002/tui-calendar/commit/95ff36b9de2fb40fe19838e599f0396aae77f666))
* **ui:** add submodels for different views ([7ceb282](https://github.com/NStefan002/tui-calendar/commit/7ceb28235a608a7aeb317e2e1eb19bba1185c248))
* **ui:** basic calendar UI ([9939c02](https://github.com/NStefan002/tui-calendar/commit/9939c0271536090c9070833527b8150f7203eade))
* **ui:** center text horizontally depending on terminal width ([9b7037a](https://github.com/NStefan002/tui-calendar/commit/9b7037a59d04bf113ee6cdd2ea5efacdef0b18b3))
* **ui:** extract views and styles into separate files, add details view ([36a9423](https://github.com/NStefan002/tui-calendar/commit/36a942380147a291fb6be75eabc8ed9c5d0973f8))
* **ui:** field alignment in calendar ([b18841a](https://github.com/NStefan002/tui-calendar/commit/b18841aba9826e8a9041ad33a63fbe2afbb6bffe))
* **ui:** make views nicer, prototype for `AddEventsView` ([1142188](https://github.com/NStefan002/tui-calendar/commit/114218833231920ac29947931fdc9a9913206eda))
* **ui:** use `lipgloss.Width` for more accurate width calculation ([0dd7c50](https://github.com/NStefan002/tui-calendar/commit/0dd7c50fdf277469296ed6b3753946b08f92fb42))
* use `key` bubble for keymaps logic ([832a8f4](https://github.com/NStefan002/tui-calendar/commit/832a8f45651535260682813829e293e267c7d1f3))
* use `tea.Key*` constants for some special keys ([d06f980](https://github.com/NStefan002/tui-calendar/commit/d06f98025504482a9f5a49e5cce5a364907dd4ee))


### Bug Fixes

* **add event view:** correct date, reset form when entering add event view from event details ([4a0af2f](https://github.com/NStefan002/tui-calendar/commit/4a0af2f3ec3c774804f35f0f96660faeea8857ff))
* **add event view:** keymaps get passed to active fields correctly ([490b49a](https://github.com/NStefan002/tui-calendar/commit/490b49aba314960a0598bc47c57773b4324920e4))
* correctly create all-day event ([7370921](https://github.com/NStefan002/tui-calendar/commit/737092129f46f3e079ad8142324311507e7d3605))
* correctly handle all-day events ([a9e431b](https://github.com/NStefan002/tui-calendar/commit/a9e431b18f75a33084fc4ad18a9388ea7a99b047))
* **event view:** always return from event details view to calendar view ([d12158a](https://github.com/NStefan002/tui-calendar/commit/d12158affcc51ef826cd734d00c89ac504a16953))
