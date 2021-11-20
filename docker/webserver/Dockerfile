FROM node:14.17 AS ANGULAR_BUILD
RUN npm install -g @angular/cli@12.2.0
COPY webapp /webapp
WORKDIR /webapp
RUN npm install && ng build

FROM golang:1.17.3-alpine3.13 AS GO_BUILD
COPY server /server
WORKDIR /server
RUN go build -o /go/bin/server

FROM alpine:3.14.0
WORKDIR /app
COPY --from=ANGULAR_BUILD /webapp/dist/webapp/* ./webapp/dist/webapp/
COPY --from=GO_BUILD /go/bin/server ./
RUN ls
CMD ./server
