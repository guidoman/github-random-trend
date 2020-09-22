# Github random trend API

This is a Rest API that will return (or redirect to) a random URL of a trending Github repository.

## Usage
This project uses the [Github Unofficial Trending API](https://github.com/huchenme/github-trending-api), by [Hu Chen](https://github.com/huchenme). All query string parameters of Hu Chen's API are supported:
- `language`
- `since`
- `spoken_language_code`

This API accepts an additional parameter named `redirect`:
- if `true`, the request will be redirected to the resulting random Github repository
- otherwise, the random repository URL is returned as plain text

The API is currently deployed on Google Cloud App Engine.

## Examples
Get the URL of a random repository:  <a href="https://github-random-trend-dot-guidoc-tech.oa.r.appspot.com/repo" target="_blank">https://github-random-trend-dot-guidoc-tech.oa.r.appspot.com/repo?redirect=true</a>

Redirect to random Python repository trending this week:  <a href="https://github-random-trend-dot-guidoc-tech.oa.r.appspot.com/repo?language=python&since=week&redirect=true" target="_blank">https://github-random-trend-dot-guidoc-tech.oa.r.appspot.com/repo?language=python&since=week&redirect=true</a>

Get the URL of a random trending Italian-speaking repository: <a href="https://github-random-trend-dot-guidoc-tech.oa.r.appspot.com/repo?spoken_language_code=it" target="_blank">https://github-random-trend-dot-guidoc-tech.oa.r.appspot.com/repo?spoken_language_code=it</a>

## Usage suggestions
You could, for example, set the home page of your browser to this API URL (with `redirect=true`). Everytime you open a new tab or page you will discover and learn something new:
- frameworks, libraries, applications
- ideas
- company or developer
- ...

## Notes
This API is implemented with the [Go programming language](https://golang.org/).