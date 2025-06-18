# transit

Using 511 transit data to make things

I'm trying to keep this backend as lightweight as possible without importing any packages, so that this can serve either as a really light executable or module, which ever i decide makes sense to turn this into

You'll need a few things to get started
[511 transit API Key](https://511.org/open-data/token)
This is free to sign up for, get, and use. I think tokens get rate limited upon requests when they exceed more than 60 requests per 3600 seconds, so keep that in mind

After you get your API key, you'll want to export it to the env variable `TRANSIT_DATA_API_KEY`, so something like this if you need the help:

```
export TRANSIT_DATA_API_KEY=<your api key>
```

Thats it for now (as of 1-15-2025)

This project depends on using go version 1.24.0 at least. This is for the quality of life improvement for using go tools that was released with this version.
To regenerate the graphql go code, run:

```
go tool gqlgen
```

Nice! (3-11-25)

Syncing data

at some point, we'll want to have the data requested from the 511 API to be cached in some way. I believe the order of latency would likely go internet request > db query > cache.

If we want to update the database that will store the documents, we could create a task queue, with built in priorities, which would be impressive if it was configurable on usage metrics or user preference of transit.

It will be first good to test the performance of a request vs a database query for the data. Then we could implement caching, but only in front of the method that is fastest

//

gql server is started with `go run server.go`
