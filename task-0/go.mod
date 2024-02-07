//renamed from "main" to "main-module" to run test with "go test ./"
module main-module

require (
    hello/hello v0.0.0
)

replace hello/hello => ./hello

go 1.21.6
