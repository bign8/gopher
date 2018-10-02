# Gopher
Googly eyed gopher

## About

This is a sample web application that demonstrates how to build a simple web application using [GopherJS](https://github.com/gopherjs/gopherjs).


## Licence

The Go gopher was designed by Renee French.  
[http://reneefrench.blogspot.com/](http://reneefrench.blogspot.com/)  

SVGs designed by [keygx](https://github.com/keygx) and licensed under the Creative Commons 3.0 Attributions license.  
[CC BY 3.0](https://creativecommons.org/licenses/by/3.0/)

This code is licensed under (TODO: figure out this license)


## Docker Development

For local development, the serving logic has been encapsulated in a dockerfile.

```sh
docker build -t bign8/gopher .
docker run -it --rm -p 8080:8080 bign8/gopher .
curl localhost:8080
```
