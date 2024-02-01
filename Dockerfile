# BUILD STAGE
# add AS BUILDER for creating build stage
FROM golang:1.20-alpine3.18 AS builder
# declare current working directiry inside the image
WORKDIR /app
# first point mean that we copy all from current dir
COPY . .

RUN go build -o project/main project/main.go


# RUN STAGE
# use alpine linux image
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/project/main .
COPY project/app.env .

EXPOSE 8080
CMD ["/app/main"]
