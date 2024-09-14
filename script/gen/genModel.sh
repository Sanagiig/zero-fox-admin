pwd
echo $1
if [ -n $1 ] 
then
  go run ./rpc/sys/gen/generator.go
else 
  go run ./rpc/$1/gen/generator.go
fi
