# assert

An assertion library for Go.

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
