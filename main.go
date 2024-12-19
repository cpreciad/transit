package main

import (
    grpc "github.com/cpreciad/transit/api/grpc"
    graphQL "github.com/cpreciad/transit/api/graphql"
)

func main(){
    grpc.ApiCall()
    graphQL.ApiCall()
}
