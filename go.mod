module github.com/b-open-io/bap-overlay

go 1.24.2

require (
	github.com/4chain-ag/go-overlay-services v0.0.0-00010101000000-000000000000
	github.com/GorillaPool/go-junglebus v0.2.14
	github.com/b-open-io/overlay v0.0.0-20250409215804-0af5fcd6d003
	github.com/bitcoin-sv/go-templates v0.0.0-00010101000000-000000000000
	github.com/bsv-blockchain/go-sdk v1.1.22
	github.com/gofiber/fiber/v2 v2.52.6
	github.com/joho/godotenv v1.5.1
	github.com/redis/go-redis/v9 v9.7.3
)

require (
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/centrifugal/centrifuge-go v0.10.2 // indirect
	github.com/centrifugal/protocol v0.10.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gookit/color v1.5.4 // indirect
	github.com/gookit/goutil v0.6.18 // indirect
	github.com/gookit/gsr v0.1.0 // indirect
	github.com/gookit/slog v0.5.8 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/philhofer/fwd v1.1.3-0.20240916144458-20a13a1f6b7c // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/segmentio/encoding v0.3.6 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	github.com/tinylib/msgp v1.2.5 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.59.0 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/exp v0.0.0-20220909182711-5c715a9e8561 // indirect
	golang.org/x/sync v0.13.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/term v0.31.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/protobuf v1.36.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/4chain-ag/go-overlay-services => github.com/4chain-ag/go-overlay-services v0.1.1-0.20250415204231-5c9e736aceec

// replace github.com/4chain-ag/go-overlay-services => ../go-overlay-services

replace github.com/bsv-blockchain/go-sdk => github.com/b-open-io/go-sdk v1.1.22-0.20250429153431-0a7620b47c68

// replace github.com/bsv-blockchain/go-sdk => ../go-sdk

replace github.com/bitcoin-sv/go-templates => github.com/b-open-io/go-templates v0.0.0-20250429144338-87ffb904e7f0

// replace github.com/bitcoin-sv/go-templates => ../go-templates

replace github.com/b-open-io/overlay => github.com/b-open-io/overlay v0.0.0-20250429210952-102a532767c6

// replace github.com/b-open-io/overlay => ../overlay
