FROM golang:1.8

COPY . "$GOPATH/src/github.com/MBControlGroup/MBCG-BE-SMS/"
RUN cd "$GOPATH/src/github.com/MBControlGroup/MBCG-BE-SMS" && go get -v && go install -v

WORKDIR $GOPATH/src/github.com/MBControlGroup/MBCG-BE-SMS

CMD ["go", "run", "main.go"]
