base=$(cd `dirname $0`; pwd)

cd $base;

go mod tidy

go run cmd/* ${1} ${2}
