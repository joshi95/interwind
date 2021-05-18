.PHONY: clean

build:
	go build ./interwind/app/api -o ../../build/interwind-app; \
	./interwind-app;
