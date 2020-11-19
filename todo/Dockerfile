FROM golang:1.13
WORKDIR /todo
COPY . .
RUN go build main.go 
ENTRYPOINT [ "./main"]