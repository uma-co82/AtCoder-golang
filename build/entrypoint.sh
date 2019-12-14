#!/bin/sh
cd /go/src/github.com/uma-co82/AtCoder-golang

# happybirthday
echo 'Happy Birthday! task start...'
go run cmd/happybirthday/main.go 5 4

# ringofavoritenumbers
echo 'Ringos Favorite Numbers	task start...'
go run cmd/ringofavoritenumbers/main.go 0 5

# threeortwo
echo '*3 or /2 task start...'
go run cmd/threeortwo/main.go 3 5 2 4

#patisserie
echo 'Patisserie ABC task start...'
go run cmd/patisserie/main.go 5 3 3 1 4 1 5 9 2 6 5 3 5 8 9 7 9

echo 'All task finished!!'