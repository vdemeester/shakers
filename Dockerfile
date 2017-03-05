FROM golang:1.8

RUN go get golang.org/x/tools/cmd/cover
RUN go get github.com/golang/lint/golint
RUN go get github.com/Masterminds/glide

WORKDIR /go/src/github.com/vdemeester/shakers

COPY glide.yaml glide.yaml
RUN glide up

COPY . /go/src/github.com/vdemeester/shakers
