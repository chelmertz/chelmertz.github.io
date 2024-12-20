.PHONY: docs
docs:
	go run cmd/blog/blog.go

.PHONY: watch
watch:
	watchexec --restart --filter 'assets/**' --filter 'cmd/**' --filter '_posts/**' --filter 'tables/**' make

.PHONY: serve
serve: docs
	serve -d docs

.PHONY: new
new:
	go run cmd/new/new.go

.PHONY: lint
lint:
	rm -rf wget.log 192.168.1.211:7999
	wget --spider --recursive --no-verbose --output-file=wget.log --reject=jpg,png,gif --max-redirect=5 http://192.168.1.211:7999
