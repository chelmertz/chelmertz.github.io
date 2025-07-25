.PHONY: docs
docs: blog
	@./blog

blog: cmd/blog/blog.go cmd/blog/atom.go
	go build -o blog ./cmd/blog

.PHONY: watch
watch:
	watchexec --restart --filter 'assets/**' --filter 'cmd/**' --filter '_posts/**' --filter 'tables/**' make

.PHONY: serve
serve: docs
	serve -d docs

.PHONY: new
new:
# @ so that the output is only the filename, so that we can do `vim $(make new)`
	@go run cmd/new/new.go

.PHONY: latest
latest:
# so that `vim $(make latest)` saves me from tabbing a lot
	@echo "_posts/$(shell ls -t _posts | head -n1)"

.PHONY: lint
lint:
	rm -rf wget.log 192.168.1.211:7999
	wget --spider --recursive --no-verbose --output-file=wget.log --reject=jpg,png,gif --max-redirect=5 http://192.168.1.211:7999

.PHONY: profile-mem
profile-mem:
	PROFILE=MEM go run cmd/blog/blog.go
	go tool pprof -http=:6060 mem.pprof , localhost:6060

.PHONY: profile-cpu
profile-cpu:
	PROFILE=CPU go run cmd/blog/blog.go
	go tool pprof -http=:6060 cpu.pprof , localhost:6060

