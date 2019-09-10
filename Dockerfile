FROM golang:1.12

MAINTAINER navygong@gmail.com

ENV TZ=Asia/Shanghai
ENV path /go/src/o2k8s

WORKDIR ${path}
COPY . ${path}

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone \
    && go build -i -v -o o2k8s \
    && cp o2k8s /usr/bin/ \
    && rm -rf /go/src/o2k8s \
    && rm -rf /go/pkg/o2k8s \
    && rm -rf /go/pkg/linux_amd64/o2k8s

CMD ["o2k8s"]
