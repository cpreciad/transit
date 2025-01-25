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
