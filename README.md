Slides for Chicago Ultimate Go Meetup, October 18:

http://go-talks.appspot.com/github.com/kkeuning/cug/slides/gatt.slide

# Go APIs: Just Generate Everything With Goa

aka, "generate all the things!"

![alt tag](https://github.com/kkeuning/cug/blob/master/slides/gatt.jpg)

An Introduction to Goa for Microservices and SPAs

# Get this repo first for the slides and examples:
```
go get github.com/kkeuning/cug
```

## To present the slides locally:
```
Install present
go get golang.org/x/net
go get golang.org/x/tools
go install golang.org/x/tools/cmd/present
cd $GOPATH/src/github.com/kkeuning/cug/slides
```
Point your browser to localhost:3999 to view the slides.

```
cd $GOPATH/src/github.com/kkeuning/cug
$GOPATH/bin/present
```
## Examples

### First, install goa and gorma
```
go get -u github.com/goadesign/goa/...
go get github.com/goadesign/gorma
```
Check out the v1.0.0 tag of goa for best compatibility with my examples.
```
cd $GOPATH/src/github.com/goadesign/goa
git checkout tags/v1.0.0
cd goagen
go install
```

### Next, install the goa examples
```
go get github.com/goadesign/examples
go get github.com/goadesign/goa-cellar
go get github.com/goadesign/gorma-cellar
go get github.com/kkeuning/cug
go get github.com/kkeuning/reduxa
```
### Add my forks as remotes to get my additions to the standard examples
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
### Build the container gorma-cellar container.
```
docker build -t gorma-cellar .
```

### Fire up the React example, and connect it to goa-cellar
```
cd $GOPATH/src/github.com/kkeuning/cug/examples/react-goa-cellar
docker-compose build
docker-compose up
```

### Seed your database
```
cd $GOPATH/src/github.com/gorma-cellar/seeder
go build
bash seeder.sh
```
Point your browser to 0.0.0.0/2015
Click Bottles to connect to the api.  
I suggest doing this with you browser developer tools open recording network activity.
Also, the Chrome Redux devtools are nice and enabled for the example.


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
