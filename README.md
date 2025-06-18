## transit

see the DEVELOPER.md docs for a bit more meat. this readme is aspirations for this project

### Syncing data

at some point, we'll want to have the data requested from the 511 API to be cached in some way. I believe the order of latency would likely go internet request > db query > cache.

If we want to update the database that will store the documents, we could create a task queue, with built in priorities, which would be impressive if it was configurable on usage metrics or user preference of transit.

It will be first good to test the performance of a request vs a database query for the data. Then we could implement caching, but only in front of the method that is fastest
