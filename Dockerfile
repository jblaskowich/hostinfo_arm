FROM golang:1.11.5

LABEL maintainer="jblaskowich@gmail.com"

WORKDIR /

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -a -installsuffix cgo -o app .

FROM scratch  

COPY --from=0 app /

CMD ["/app"]