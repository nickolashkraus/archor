help: ## Print help for targets with comments
	@cat $(MAKEFILE_LIST) | \
		grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' | \
		sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-16s\033[0m %s\n", $$1, $$2}'
