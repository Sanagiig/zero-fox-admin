# $1 模块名
echo gen $1 api
goctl api go --api ./api/$1/doc/api/$1.api --dir ./api/$1/ --style go_zero