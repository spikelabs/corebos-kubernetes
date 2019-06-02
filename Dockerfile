FROM golang:1.10

RUN apt-get install -y tzdata bash wget curl git

RUN mkdir -p $$GOPATH/bin && curl https://glide.sh/get | sh
RUN mkdir /go/src/corebos-kubernetes

WORKDIR /go/src/corebos-kubernetes

COPY ./glide.yaml .
COPY ./glide.lock .

RUN glide install

COPY . .

CMD go run main.go

EXPOSE 80

