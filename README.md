# SafeConversion

SafeConversion provides safe type casting functions for Go.

## Installation

```bash
go get github.com/corentings/safeconversion
```

## Usage

```go
import "github.com/corentings/safeconversion"
```

### CastInt

CastInt provides a safe way to cast an integer from one type to another.

```go
func CastInt[From, To safeconversion.Integer](value From) (To, error)
```

#### Example

```go
result, err := safeconversion.CastInt[int32, int64](math.MaxInt32)
// result is 9223372036854775807 and err is nil

result, err = safeconversion.CastInt[int64, int32](math.MaxInt64)
// result is 0 and err is safeconversion.ErrValueOutOfRange
```

### CastFloat

- TODO

### CastString

- TODO

## Disclaimer

This library is not meant to replace the standard type casting functions. It is meant to provide a safe alternative when the standard functions would silently overflow or underflow. 
It's not recommended to use this library without auditing the code first to be sure it fulfills your needs. 

I am not a security expert, so use this library at your own risk. 

## License

This library is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## Acknowledgements

- [rung/go-safecast](https://github.com/rung/go-safecast)
- [cybergarage/go-safecast](https://github.com/cybergarage/go-safecast)
- [fortio/safecast](https://github.com/fortio/safecast) - I used this as a reference for the CastInt implementation and highly recommend it.
