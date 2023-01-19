⚠️ **This repository has been deprecated. Use [stretchr/testify](https://github.com/stretchr/testify) instead.**

# assert

**assert** is an assertion library for Go.

## Installation

```sh
go get crdx.org/assert
```

## Methods

All methods return a `bool` indicating test success or failure.

### `Equal(t, actual, expected)`

Assert that `actual == expected`.

### `NotEqual(t, actual, expected)`

Assert that `actual != expected`.

### `True(t, condition)`

Assert that `condition` is `true`.

### `Truef(t, condition, format, ...args)`

Assert that `condition` is `true`, and show a custom message if not.

## Contributions

Open an [issue](https://github.com/crdx/assert/issues) or send a [pull request](https://github.com/crdx/assert/pulls).

## Licence

[GPLv3](LICENCE).
