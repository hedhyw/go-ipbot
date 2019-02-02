APP_NAME=ipbot
APP_CMD=./cmd/$(APP_NAME)
APP_DIST=./dist/$(APP_NAME)

DEBUG_ENV=./env/debug.sh
RUN_SCRIPT=./scripts/run.sh $(DEBUG_ENV) $(APP_CMD)
TEST_SCRIPT=./scripts/test.sh $(DEBUG_ENV)

GREEN='\033[0;32m'
BLUE='\033[0;34m'
FMT='\033[0m'
DONE=@echo -e $(GREEN)Done$(FMT)

all: test build
build:
	@echo -e $(BLUE)Building...$(FMT)
	CGO_ENABLED=0 vgo build -o $(APP_DIST) $(APP_CMD)
	$(DONE)
run:
	$(RUN_SCRIPT)
	$(DONE)
test:
	@echo -e $(BLUE)Testing...$(FMT)
	$(TEST_SCRIPT)
	$(DONE)
