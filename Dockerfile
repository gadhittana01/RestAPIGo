FROM golang:alpine

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o /app/bin ./

EXPOSE 8000

CMD [ "sh", "-c", "/app/bin/restapi"]