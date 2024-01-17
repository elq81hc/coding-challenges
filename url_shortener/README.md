# Requirements
- Create a REST API that will allow a client to add a URL with the following format:
```json
# Request body:
{
    "url": "https://www.google.com"
}

# Success response:
{
    "key":"wsf5f",
    "long_url":"https://www.google.com",
    "short_url":"http://localhost/wsf5f"
}

# Error response:
{
    "error": "Missing field: url"
}
```
- The API call to create a shortened URL should be idempotent (this means that if you send multiple identical requests, only the initial request would cause a change).
- Add a check to ensure that a duplicate key that is generated from a different long URL is handled.

# Design
## Step 1 - Understand the problem and establish design scope
1. Clarification questions.
- what is the traffic volume?
- how long is the shortened URL?
- what characters are allowed in the shortened URL? -> [0-9a-zA-Z] (62 characters)
- can shortened URL be deleted or updated? -> no

2. The envelope estimation
- Write operations: 1000 QPS
- Read operations: 1000 * 10 = 10000 QPS
- Assume average URL length is 100 characters
- Storage yearly: 1000 * 86400 * 365 * 100 ~ 3.15TB
- Number of records yearly: 1000 * 86400 * 365 ~ 31.5 million records

# Step 2 - High level design
1. API endpoints
- URL shortening:
POST api/v1/shorten
    - request parameter: {url: longURLString}
    - return shortURLString

- URL redirection:
GET api/v1/{shortURLString}
    - return longURL for HTTP redirection

2. URL redirecting
One thing worth discussing is 301 redirect vs 302 redirect.
- 301 redirect: **permanent redirect**. Since it is permanently redirected, the browser will cache the redirect and will not make a request to the original URL unless the cache is cleared.
- 302 redirect: **temporary redirect** meaning that subsequent requests will still be made to the original URL, then the server will redirect the request to the original URL.
- Use case: if the priority is reduce the server load, then 301 redirect is better. However, if the priority is to track the number of clicks and source of the click, then 302 redirect is better.
