FROM golang:1.23.9-alpine

WORKDIR /app

# Desativa mecanismo automático de toolchain para evitar erro
ENV GOTOOLCHAIN=local

COPY . .

RUN go mod tidy

RUN go build -o ve-backend .

CMD ["./ve-backend"]
