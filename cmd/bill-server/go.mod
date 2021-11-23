module github.com/hades300/bill-center/cmd/bill-server

require (
	github.com/gogf/gf/v2 v2.0.0-beta
	github.com/hades300/bill-center v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
)

go 1.14

replace github.com/willf/bitset v1.2.1 => github.com/bits-and-blooms/bitset v1.2.1

replace github.com/hades300/bill-center => ../../.
