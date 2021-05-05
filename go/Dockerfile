FROM golang:alpine

RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
 
ENTRYPOINT [ "top" ]

EXPOSE 8080