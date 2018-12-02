* golang资料参考：https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md
* golang工具安装参考：https://www.jianshu.com/p/35f8f57ab074


```
mkdir -p $GOPATH/src/golang.org/x 
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/tools.git
go get golang.org/x/tools/cmd/godoc
go get -u -v github.com/nsf/gocode
go get -u -v github.com/rogpeppe/godef
go get -u -v github.com/golang/lint/golint
go get -u -v github.com/lukehoban/go-find-references
go get -u -v github.com/lukehoban/go-outline
go get -u -v sourcegraph.com/sqs/goreturns
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v github.com/tpng/gopkgs
go get -u -v github.com/newhook/go-symbols

调试工具：
go get github.com/derekparker/delve/cmd/dlv

如果不行，先进行解析
192.168.123.209 localhost
192.30.253.112 github.com
173.194.75.141 golang.org
173.194.71.141 golang.org
```
