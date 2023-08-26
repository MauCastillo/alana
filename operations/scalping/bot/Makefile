LOGDATE = $(shell date +'%Y_%m_%d')
BUILD_NAME_TEST = test-latest_$(LOGDATE)
BUILD_NAME_MAIN = main-latest_$(LOGDATE)
BUILD_ARM_NAME_MAIN = main-arm64

build:
	go build -o $(BUILD_NAME_MAIN)
	cp -b $(BUILD_NAME_MAIN)  ../../../build/principal/
	rm -f $(BUILD_NAME_MAIN) 

build_arm:
	export CC=aarch64-linux-gnu-gcc
	export GOARCH=arm64
	export GOOS=linux
	
	CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build -o $(BUILD_ARM_NAME_MAIN) main.go
	# cp -b $(BUILD_ARM_NAME_MAIN)  ../../../build/arm/
	# rm -f $(BUILD_ARM_NAME_MAIN) 

test:
	go build -o $(BUILD_NAME_TEST)
	cp -b $(BUILD_NAME_TEST)  ../../../build/test/
	rm -f $(BUILD_NAME_TEST) 

clean:
	rm -f $(BUILD_NAME_TEST)
	rm -f $(BUILD_NAME_MAIN)

	rm -rf ../../../build/test/
	rm -rf ../../../build/principal/