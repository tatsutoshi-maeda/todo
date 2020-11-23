FROM golang:1.13
WORKDIR /todo
COPY todo/. .
ENV GO111MODULE=on
# RUN go mod download
EXPOSE 8080
RUN go build main.go 
ENTRYPOINT [ "./main"]