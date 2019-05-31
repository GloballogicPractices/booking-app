.PHONY: init deploy
#.DEFAULT_GOAL := help
.DEFAULT_GOAL := init

help:
	# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST)  | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build:
	# Build every service
	@cd eventsmgmt && docker build -t eventmgmt:0.0.01 .
