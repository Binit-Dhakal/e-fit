FROM golang:1.24-alpine     

RUN apk add --no-cache git bash curl && \
    go install github.com/air-verse/air@latest


WORKDIR /app 

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

CMD ["air"]
