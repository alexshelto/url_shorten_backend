# url_shorten_backend
Writen in Go using Gin & SQLite

## Shortens URLs with a custom base62 encoding
converts numerical values to 0-9a-zA-Z

## Endpoints 

Note: `:hash` is for the shortened url key, `:id` is for the id of the generated entry in the table

| endpoint           | method | body                         | response success | response body                                                                                                                                                  |
|--------------------|--------|------------------------------|------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| /l                 | POST   | {"original_url": "some_url"} | 200              | ```{     "ID":int, "CreatedAt":datettime,     "UpdatedAt":datettime, "DeletedAt":null, "OriginalUrl":string,    "ShortenedUrl":string,  "VisitCount": int }``` |
| /l/id/:id          | GET    |                              | 200              |                                                                                                                                                                |
| /l/:hash           | GET    |                              | 301              |                                                                                                                                                                |
| /l/analytics/:hash | GET    |                              | 200              |                                                                                                                                                                |
| /l/create/form     | GET    |                              | 200              |                                                                                                                                                                |
| /l/create/form     | POST   | form: ("url": "some_url")    | 200              |                                                                                                                                                                |
