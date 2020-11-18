#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
ADD . /go/src/gonews
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
ENV DB_USER $DB_USER
ENV DB_PASSWORD $DB_PASSWORD
ENV DB_NAME $DB_NAME
ENV API_KEY $API_KEY
ENV SECRET $SECRET
ENV SERVER_PORT $SERVER_PORT
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT ./app
LABEL Name=gonews Version=0.0.1
EXPOSE 3456
