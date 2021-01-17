LOGFILE=$(LOGPATH) `date +'%A-%b-%d-%Y-%H-%M-%S'`

.PHONY: hp
hp: ## 🌱 This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help

.PHONY: cm
cm: ## 🌱 git commit
	@echo '************👇  run command 👇************'
	git add .
	git commit -m "🌱Go-Multi-Stage-Docker💙-${LOGFILE}"
	git push -u origin main