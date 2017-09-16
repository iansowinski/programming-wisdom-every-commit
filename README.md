# programming wisdom tweet after every commit

[Blogpost](blog.iansowinski.com/2017/09/15/Code-Quality)

...So this is faster version of that small script. Instead of packing [anaconda](https://github.com/ChimeraCoder/anaconda) to binary I decided to just generate array of tweets and harcode it with binary. Will have to make makefile, but for now it works this way:

1. `go get github.com/ChimeraCoder/anaconda && go get github.com/fatih/color` 
2. add your twitter credentials to `project/tweets.go`
3. `cd project && go generate`
4. `cd .. && go build wisdom.go` (`GOOS=platform GOARCH=architecture go build wisdom.go` for cross-compiling)
5. Congrats! Your binary is ready!