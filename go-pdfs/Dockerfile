# Etapa 1: Compilación del código
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Inicializar el módulo si no existe y copiar código
COPY . .
RUN go mod init image-processor || true
RUN go build -o main .

# Etapa 2: Entorno de ejecución ultra ligero
FROM alpine:3.18

WORKDIR /app

# Copiar solo el binario compilado desde la etapa anterior
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]