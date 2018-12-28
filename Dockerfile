FROM  golang:1.10.1-alpine

RUN echo "Installing git and bash"
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN echo "Installing glide"

RUN go get -u github.com/Masterminds/glide
