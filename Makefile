GO=go
# output binary name
OUT=go-talk


exec:
	$(GO) build -o $(OUT)
	./$(OUT)

build:
	$(GO) build -o $(OUT)

clean:
	rm $(OUT)
