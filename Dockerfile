
# docker build -t nubol .
# docker rm nu
# docker run -d -p 3000:3000 --name nu nubol
# docker logs nu


FROM node:10.16.0-alpine AS buildui
WORKDIR /usr/src/app
COPY package.json .
COPY package-lock.json .
RUN npm install
COPY client/ client/
RUN npm run build

FROM golang:1.12.6 AS buildserver
WORKDIR /go/src/app
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/gin-contrib/static
COPY server/ ./server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main server/main.go

FROM alpine:3.9.4
WORKDIR /usr/src/app
COPY --from=buildui /usr/src/app/dist ./dist
COPY --from=buildserver /go/src/app/main ./main
ENTRYPOINT ["/usr/src/app/main"]
CMD [ "-root", "/" ]
