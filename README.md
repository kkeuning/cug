Slides for Ultimate Go Meetup, October 18:

http://go-talks.appspot.com/github.com/kkeuning/goa-react/slides/justgenerate.slide

# Introduction to Goa for Microservices and SPAs
```
go get github.com/kkeuning/goa-react
```

## Slides
```
Install present
go get golang.org/x/net
go get golang.org/x/tools
go install golang.org/x/tools/cmd/present
```
```
cd $GOPATH/src/github.com/kkeuning/goa-react
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
### Fire up the react example, and connect it to goa-cellar
```
cd examples/react-goa-cellar
npm i
npm start
```
Connect your browser to localhost:3000.

By default will connnect to cellar.goa.design, modify src/index.js to connect to your local instance.

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
