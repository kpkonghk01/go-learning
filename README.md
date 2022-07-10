Personal notes :)

### Installation

[/usr/bin vs /usr/local/bin](https://www.jianshu.com/p/ea6c4758dba4)

[Install Ref](https://willh.gitbook.io/build-web-application-with-golang-zhtw/01.0/01.1#go-biao-zhun-tao-jian-an-zhuang)

[Go version control](https://github.com/moovweb/gvm)

```
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

gvm install go1.18.3 -B
gvm use go1.18.3 --default
```

## Project init

```
mkdir <your project name>
cd <your project name>
go mod init <your git host>/<your git acc>/<your project name>
# e.g. github.com/kpkonghk01/go-learning
```

### VSCode settings

[VSCode setting](https://github.com/golang/vscode-go/issues/971#issuecomment-927666108)

[VSCode syntax highlight](https://code.visualstudio.com/docs/languages/go)

### Local doc

```
# go install golang.org/x/tools/cmd/godoc@latest
godoc -http :8000
```

go to http://localhost:8000/pkg to see all the packages you installed

### Run test

```
go test ./src
```

### Need Research:

- Basic syntax (https://go.dev/doc/)
- Auto fmt on save
- lint? (golangci-lint)
- Dependency management (https://go.dev/blog/using-go-modules)
- Error handling
- How to DI (https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection)
- IOC?
- How to write tests? (https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world)
- How to handle async function call
- How to connect to db
- How to debug using debugger (https://github.com/golang/vscode-go/blob/master/docs/debugging.md#configure)
- How to write parallel program
- Logging best practice
- How to make a http service
- How to make cmd
- `npm run` in go? use makefile? or don't need?
