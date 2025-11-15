ROOT_DIR := $(shell git rev-parse --show-toplevel 2>/dev/null || dirname $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(shell pwd)
DIR_NAME := $(shell python3 -c "import os; print(os.path.relpath('$(CURRENT_DIR)', '$(ROOT_DIR)'))" 2>/dev/null || echo "$(shell basename $(PWD))")

# バージョンをタグ付けしてプッシュする
# e.g: make push VERSION=v1.0.0
push: .check-version
	git tag $(DIR_NAME)/$(VERSION); \
	git push origin $(DIR_NAME)/$(VERSION)
	@echo "Pushed tag: $(DIR_NAME)/$(VERSION)"
	@echo $(VERSION) > $(CURRENT_DIR)/version

# 現在のバージョンを表示する
# e.g: make version
version:
	@echo $(DIR_NAME)/$$(cat $(CURRENT_DIR)/version)

.PHONY: bump version

.check-version:
# semantic versioningの形式 (v1.0.0) を満たしているかをチェック
ifdef VERSION
	@echo "$(VERSION)" | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$$' > /dev/null || \
		(echo "Error: VERSION must be in semantic versioning format (v1.0.0)"; exit 1)
endif
