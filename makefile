
GOCMD=go
GO_HOT_CMD=fresh
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) mod tidy
KILL_CMD=tskill

GOCTL=goctl ## goctl

all: deps build ## 默认的构建目标

clean: ## 清理目标
	$(GOCLEAN)
	rm -rf target

deps: ## 安装依赖目标
	$(GOGET) -v

format: ## 格式化代码
	$(GOCTL) api format --dir api/admin/doc/api
	$(GOCTL) api format --dir api/front/doc/api
	$(GOCTL) api format --dir api/web/doc/api

copy_config:
	mkdir -p target/sys-rpc && cp rpc/sys/etc/sys.yaml target/sys-rpc/sys-rpc.yaml
	mkdir -p target/admin-api && cp api/admin/etc/admin-api.yaml target/admin-api/admin-api.yaml

gen_proto:
	${GOCMD} run ./script/proto/main.go sys

build: copy_config
	$(GOBUILD) -o target/sys-rpc/sys-rpc -v ./rpc/sys/sys.go
	$(GOBUILD) -o target/admin-api/admin-api -v ./api/admin/admin.go

start: ## 运行目标
	nohup ./target/sys-rpc/sys-rpc -f ./target/sys-rpc/sys-rpc.yaml  > /dev/null 2>&1  &
	nohup ./target/admin-api/admin-api -f ./target/admin-api/admin-api.yaml  > /dev/null 2>&1 &

stop: ## 停止目标
	$(KILL_CMD)  admin-api
	$(KILL_CMD)  sys-rpc
	echo "server is stopped."

restart: stop start ## 重启项目

gen: gen_proto	## 生成所有模块代码
	# api-admin
	$(GOCTL) api go -api ./api/admin/doc/api/admin.api -dir ./api/admin/ --style go_zero
	# sys-rpc
	$(GOCTL) rpc protoc rpc/sys/sys.proto --go_out=./rpc/sys/ --go-grpc_out=./rpc/sys/ --zrpc_out=./rpc/sys/ -m --style go_zero

model: ## 生成model代码
	go run rpc/sys/gen/generator.go