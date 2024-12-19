This is a folder set up to hold the an api necessary to communicate with the 
transit server, whatever it may be. the package name is simply `api`. 

`go get github.com/cpreciad/transit/api` is probably how this will go

This may contain gRPC/protobuffs, of which a server would implement

It may also contain graphQL setup, which hopefully follows similar suit in providing an 
interface
