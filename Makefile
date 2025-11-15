SHELL := /bin/bash -o pipefail
# リポジトリのルートディレクトリ
ROOT_DIR := $(shell git rev-parse --show-toplevel 2>/dev/null || dirname $(lastword $(MAKEFILE_LIST)))
# 現在のディレクトリ
CURRENT_DIR := $(shell pwd)
# リポジトリのルートからの相対パスを取得
ROOT_REL_PATH := $(shell python3 -c "import os; print(os.path.relpath('$(CURRENT_DIR)', '$(ROOT_DIR)'))" 2>/dev/null || echo "$(shell basename $(PWD))")
# タグ名
TAG_NAME := $(ROOT_REL_PATH)/$(VERSION)

# バージョンをタグ付けしてプッシュする
# e.g: make push VERSION=v1.0.0
push:
	git tag $(TAG_NAME); \
	git push origin $(TAG_NAME)
	@echo "Pushed tag: $(TAG_NAME)"
	@echo $(VERSION) > $(CURRENT_DIR)/version

# 現在のバージョンを表示する
# e.g: make version
version:
	@echo $(ROOT_REL_PATH)/$$(cat $(CURRENT_DIR)/version)

.PHONY: bump version

.check-version:
# semantic versioningの形式 (v1.0.0) を満たしているかをチェック
ifdef VERSION
	@echo "$(VERSION)" | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$$' > /dev/null || \
		(echo "Error: VERSION must be in semantic versioning format (v1.0.0)"; exit 1)
endif
