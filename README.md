# go-htmx-go
Go and HTMX 

## Install protobuf and grpc

`brew install pprotobuf`

`go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`


# Test grpc via grpc curl

Install `brew install grpcurl`

List ` grpcurl --plaintext localhost:9092 list`


lits `grpcurl --plaintext localhost:9092 list currency.Currency`

describe `grpcurl --plaintext localhost:9092 describe currency.Currency.GetRate`

describe `grpcurl --plaintext localhost:9092 describe .currency.RateRequest`

Invoke `grpcurl --plaintext -d '{"Base": "GPB", "Destination": "USD"}' localhost:9092 currency.Currency.GetRate`

