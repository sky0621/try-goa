# try-goa

## Ref

https://goa.design/ja/

## Method

- Make "design" directory

- add design/design.go

- exec goagen

at $GOPATH/src/github.com/sky0621/try-goa
$ goagen bootstrap -d github.com/sky0621/try-goa/design

- build and exec app

go build -o cellar
./cellar

curl -i localhost:1456/bottles/1
