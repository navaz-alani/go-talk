GO=go

# output binary name
OUT=go-talk

# build and run app
exec:
	$(GO) build -o $(OUT)
	./$(OUT)

build:
	$(GO) build -o $(OUT)

# clean up output binaries
clean:
	rm $(OUT)
