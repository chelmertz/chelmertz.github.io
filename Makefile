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

.PHONY: profile-mem
profile-mem:
	PROFILE=MEM go run cmd/blog/blog.go
	go tool pprof -http=:6060 mem.pprof , localhost:6060

.PHONY: profile-cpu
profile-cpu:
	PROFILE=CPU go run cmd/blog/blog.go
	go tool pprof -http=:6060 cpu.pprof , localhost:6060

