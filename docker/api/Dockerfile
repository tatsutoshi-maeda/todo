FROM golang:1.13 as builder
WORKDIR /todo
COPY todo/. .
ENV GO111MODULE=on
RUN go build main.go 

FROM golang:1.13
WORKDIR /todo
COPY --from=builder /todo/main ./main
EXPOSE 8080
ENTRYPOINT [ "./main"]