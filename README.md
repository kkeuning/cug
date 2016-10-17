Slides for Ultimate Go Meetup, October 18:

http://go-talks.appspot.com/github.com/kkeuning/cug/slides/gatt.slide

# Introduction to Goa for Microservices and SPAs
```
go get github.com/kkeuning/cug
```

## Slides
```
Install present
go get golang.org/x/net
go get golang.org/x/tools
go install golang.org/x/tools/cmd/present
```
```
cd $GOPATH/src/github.com/kkeuning/cug
$GOPATH/bin/present
```
## Examples

### First, install goa and gorma
```
go get github.com/goadesign/goa
go get github.com/goadesign/goa/goagen
go get github.com/goadesign/gorma
```

### Next, install the goa examples
```
go get github.com/goadesign/examples
go get github.com/goadesign/goa-cellar
go get github.com/goadesign/gorma-cellar
```
### Add my forks as remotes
My fork of goa-cellar adds support for cors
```
cd $GOPATH/src/github.com/goadesign/goa-cellar
git remote add kkeuning https://github.com/kkeuning/goa-cellar
git remote update
git checkout -t kkeuning/cors
```
My fork of gorma cellar adds several extras for the demo...
```
cd $GOPATH/src/github.com/goadesign/gorma-cellar
git remote add kkeuning https://github.com/kkeuning/gorma-cellar
git remote update
git checkout -t kkeuning/chicago
```
If you have docker installed, build the container.
```
docker build -t gorma cellar .
```

### Fire up the react example, and connect it to goa-cellar
```
cd examples/react-goa-cellar
docker-compose build
docker-compose up
```
Point your browser to 0.0.0.0/2015

### Extras
Swagger UI:
https://github.com/swagger-api/swagger-ui

Bootprint (Static Swagger UI alternative)
https://github.com/bootprint/bootprint

Vegeta:
https://github.com/tsenart/vegeta

Baloo (end to end api testing):
https://github.com/h2non/baloo

Chrome Redux devtools:
https://github.com/zalmoxisus/redux-devtools-extension

Restful API Guidelines:
http://zalando.github.io/restful-api-guidelines/

JSON API Specification:
http://jsonapi.org/
