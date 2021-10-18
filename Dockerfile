FROM golang:1.17.1 AS httpbuilder

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN mkdir bin 
RUN GOOS=linux GOARCH=amd64
RUN go build -o bin/httpserver

FROM alpine AS production
COPY --from=httpbuilder /app/bin/httpserver .
CMD ["./httpserver"]