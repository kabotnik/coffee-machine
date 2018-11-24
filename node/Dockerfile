# FROM node:11.2-alpine as build
# WORKDIR /build
# ENV GOOS="linux"
# ENV GOARCH="amd64"
# COPY . .
# RUN go build -o coffee

FROM node:11.2-alpine as release

WORKDIR /usr/src/app

COPY package*.json ./
RUN npm install

COPY . .

EXPOSE 9090
CMD [ "npm", "start" ]