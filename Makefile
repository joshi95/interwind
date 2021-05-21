.PHONY: clean

build:
	go build ./interwind/app/interwind-api -o ../../build/interwind-app 
	./interwind/build/interwind-app;
