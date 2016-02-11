# go-pngquant

**go-pngquant** is PNG image compressor go-library that wrapping `pngquant` command.
`pngquant` command **is required**.

<https://pngquant.org/>

## Usage

To install, use `go get`:

```
$ go get github.com/yusukebe/go-pngquant
```

### import

```go
import (
    pngquant "github.com/yusukebe/go-pngquant"
)
```

#### func Compress

`func Compress(i image.Image, speed string) image.Image, error`

#### func CompressBytes

`func CompressBytes(b []byte, speed string) []bytes, error`

## License

MIT

## Author

Yusuke Wada <http://github.com/yusukebe>
