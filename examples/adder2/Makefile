build-race:
	@echo "building with race detection"; \
	go build -race

clean:
	rm -rf app/
	rm -rf client/
	rm -rf schema/
	rm -rf swagger/

lint:
	golint

test:

prep:
	npm i -g bootprint
	npm i -g bootprint-openapi
	GOPATH=$(GOPATH) go get -u github.com/kardianos/govendor

gen-main:
	@if [ -a _scaffolds ]; then rm _scaffolds/*; fi;
	goagen main -d github.com/kkeuning/goa-react/examples/adder2/design -o _scaffolds
	
gen:
	goagen app -d github.com/kkeuning/goa-react/examples/adder2/design
	@if [ -a _scaffolds ]; then rm _scaffolds/*; fi;
	goagen main -d github.com/kkeuning/goa-react/examples/adder2/design -o _scaffolds
	goagen client -d github.com/kkeuning/goa-react/examples/adder2/design
	goagen js -d github.com/kkeuning/goa-react/examples/adder2/design
	goagen swagger -d github.com/kkeuning/goa-react/examples/adder2/design
	goagen gen -d github.com/kkeuning/goa-react/examples/adder2/design --pkg-path=github.com/goadesign/gorma
	bootprint openapi swagger/swagger.json api-doc

