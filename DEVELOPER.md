## Notes for Devs

If you want to develop on this project, you'll need a few things to get started

### Bay Area Transit Data API Key

[511 transit API Key](https://511.org/open-data/token)
This is free to sign up for, get, and use. I think tokens get rate limited upon requests when they exceed more than 60 requests per 3600 seconds, so keep that in mind

After you get your API key, you'll want to export it to the env variable `TRANSIT_DATA_API_KEY`, so something like this if you need the help:

```
export TRANSIT_DATA_API_KEY=<your api key>
```

### Go dependencies

This project depends on using go version 1.24.0 at least. This is for the quality of life improvement for using go tools that was released with this version.

### Generating graphql boilerplate from the schema

When changes or updates are made to the aspects of graphql in this project, there are a few important parts where this can happen. Please refer to _gqlgen_[repo](https://github.com/99designs/gqlgen) and [docs](https://gqlgen.com) as a place to get started if unfamiliar with teh gqlgen library

#### Updating the schema

When updating the graphql schema in this project, located at [/graph/schema.graphqls](/graph/schema.graphqls), a nice shortcut to generate all of the code necessary from updates is to run the following command from the root of the project:

```

go generate ./...

```

This command is generally used to genarate code generically in go, but for our purposes with the gql library, it works equally as well

#### Fine tuning the resolvers
One thing to keep in mind when extending the gql resolvers is that gqlgen.yml is responsible for configuring how each sub-dependency is resolved. if you don't specify that you want a particular gql object to be manualy resolved, you won't have control over any part of it, which importantly means that you can't really populate the important fields. To extend the resolvers for nested objects, you can look at what is done for [line objects embeded within operator objects](https://github.com/cpreciad/transit/blob/815f575b33a68738881104660fb5968bb5c1b968/gqlgen.yml#L156-L159). After adding, you'll want to re-generate everything again, and implement the new resolvers... etc 

#### Running the gql server

Pretty simple as well for now. Run the following command

```
go run server.go
```
