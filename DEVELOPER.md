## Notes for Devs

If you want to develop on this project, you'll need a few things to get started

### Bay Area Transit Data API Key

[511 transit API Key](https://511.org/open-data/token)
This is free to sign up for, get, and use. I think tokens get rate limited upon requests when they exceed more than 60 requests per 3600 seconds, so keep that in mind

After you get your API key, you'll want to export it to the env variable `TRANSIT_DATA_API_KEY`, so something like this if you need the help:

```
export TRANSIT_DATA_API_KEY=<your api key>
```
