NAME := tomo
FLAGS := -trimpath -ldflags="-s -w"
TARGETS := linux windows macos

all: build-app

generate:
	go generate ./...

build-app: build-app-windows-amd64 build-app-linux-amd64

build-app-windows-amd64: generate
	$(eval TARGET=x86_64-windows-gnu)
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC="zig cc -target $(TARGET)" CXX="zig c++ -target $(TARGET)" go build $(FLAGS) -o target/windows/$(NAME).exe ./app/...

build-app-linux-amd64: generate
	$(eval TARGET=x86_64-linux-gnu)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC="zig cc -target $(TARGET)" CXX="zig c++ -target $(TARGET)" go build $(FLAGS) -o target/linux/$(NAME) ./app/...

clean:
	rm -rf target/*
