FROM golang:alpine3.7 as build
WORKDIR /build
ENV GOOS="linux"
ENV GOARCH="amd64"
COPY . .
RUN go build -o coffee

FROM alpine:3.7 as release
COPY --from=build ./build/coffee /coffee
ENV PRODUCT="kcompact"

EXPOSE 9090
CMD ["/coffee"]