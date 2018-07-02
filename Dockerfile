FROM alpine
LABEL maintainer="fredliang"

RUN apk --no-cache add tzdata  ca-certificates && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app
ENV GIN_MODE release
ADD config/ /app/config/
ADD main /app/main
ADD docs /app/docs/
RUN chmod +x ./main
CMD ["./main"]