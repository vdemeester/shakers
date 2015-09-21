FROM golang:1.5

RUN go get github.com/tools/godep
RUN go get golang.org/x/tools/cmd/cover
RUN go get github.com/golang/lint/golint
RUN go get golang.org/x/tools/cmd/vet

WORKDIR /go/src/github.com/vdemeester/shakers

RUN mkdir Godeps
COPY Godeps/Godeps.json Godeps/
RUN godep restore

COPY . /go/src/github.com/vdemeester/shakers
