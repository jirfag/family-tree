FROM alpine
MAINTAINER fredliang

RUN apk add --no-cache ca-certificates
WORKDIR /app
ENV GIN_MODE release
ADD config.deploy.yml /app/config.deploy.yml
ADD family-tree /app/family-tree
RUN chmod +x ./family-tree
CMD ["./family-tree"]