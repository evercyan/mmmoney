
define check
	command -v $(1) 1>/dev/null || $(2)
endef

help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## release: goreleaser
release:
	@@$(call check,goreleaser,brew install goreleaser)
	goreleaser release --rm-dist

.PHONY: help release