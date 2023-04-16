# Tree URL Crawler

Enter a URL and it will crawl, explore, and collect the URLs for you!

Built using Go. I created this project as a way to practice using Go and its concurrency features.

The crawler includes rate-limiting and basic caching capabilities, and uses a simple Depth-First Search (DFS) algorithm for crawling.

Some important links:

* View the app: [crawler-fe.vercel.app](https://crawler-fe.vercel.app/)
* Front-end source code: [https://github.com/tsaqifammar/crawler-fe](https://github.com/tsaqifammar/crawler-fe)

## View

![page](https://user-images.githubusercontent.com/54428874/232195481-f32b8fee-8eea-4280-a6f0-331afa8c1529.png)

## End-point

### <code>GET</code> <code><b>/crawl</b></code>

To crawl for URLs, given a root URL, a depth, and a max URL to be crawled.

#### Parameters

| name   | description                              |
| ------ | ---------------------------------------- |
| url    | The root url.                            |
| depth  | The maximum depth to be crawled.         |
| maxUrl | The maximum number of URLs to be crawled |

#### Example

```js
GET /crawl?url=https://someurl.com&depth=3&maxUrl=12

Response:
{
  "url": "https://someurl.com",
  "childUrls": [
    {
      "url": "https://url11.com",
      "childUrls": [
        {
          "url": "https://url111.com",
          "childUrls": [],
        },
        {
          "url": "https://url112.com",
          "childUrls": [],
        },
      ],
    },
    {
      "url": "https://url12.com",
      "childUrls": [
        {
          "url": "https://url121.com",
          "childUrls": [],
        },
        {
          "url": "https://url122.com",
          "childUrls": [],
        },
        {
          "url": "https://url123.com",
          "childUrls": [],
        },
      ],
    },
    {
      "url": "https://url13.com",
      "childUrls": [],
    },
    {
      "url": "https://url14.com",
      "childUrls": [
        {
          "url": "https://url141.com",
          "childUrls": [],
        },
      ],
    },
  ],
}
```
