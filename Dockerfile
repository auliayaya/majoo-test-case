FROM golang:alpine as builder

WORKDIR /app 

COPY . ./
COPY app.env ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" main.go

FROM scratch

WORKDIR /app

COPY app.env ./app


COPY --from=builder /app/cmd /usr/bin/

ENTRYPOINT ["cmd"]