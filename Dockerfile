FROM golang:1.12.0-stretch

RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list && \
    sed -i 's|security.debian.org/debian-security|mirrors.ustc.edu.cn/debian-security|g' /etc/apt/sources.list && \
    apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 68818C72E52529D4 && \
    echo "deb http://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.0 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-4.0.list

RUN apt-get update
RUN apt-get install -y vim net-tools lsof nghttp2 mongodb

ADD . /var/source
WORKDIR /var/source
RUN GOPROXY=https://goproxy.io go mod download
RUN go build -o /go/bin/grpc-client /var/source/cmd/grpc-client &&\
    go build -o /go/bin/grpc-server /var/source/cmd/grpc-server &&\
    go build -o /go/bin/http-server /var/source/cmd/http-server &&\
    go build -o /go/bin/h2c-server /var/source/cmd/h2c-server
