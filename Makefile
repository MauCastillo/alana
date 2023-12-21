LOGDATE = $(shell date +'%Y_%m_%d')
BUILD_NAME_TEST = test-latest_$(LOGDATE)
BUILD_NAME_MAIN = main-latest_$(LOGDATE)
BUILD_ARM_NAME_MAIN = raspberry-bot_$(LOGDATE)
MAIN_FILE_PATH = operations/scalping/bot/main.go

run: 
	go run $(MAIN_FILE_PATH)
	
build:
	go build -o $(BUILD_NAME_MAIN)
	cp -b $(BUILD_NAME_MAIN)  ../../../build/principal/
	rm -f $(BUILD_NAME_MAIN) 

build_arm:	
	env GOOS=linux GOARCH=arm GOARM=5 go build -o $(BUILD_ARM_NAME_MAIN) $(MAIN_FILE_PATH)

test:
	go build -o $(BUILD_NAME_TEST) $(MAIN_FILE_PATH)

clean:
	rm -f $(BUILD_NAME_TEST)
	rm -f $(BUILD_NAME_MAIN)
	rm -f $(BUILD_ARM_NAME_MAIN)

