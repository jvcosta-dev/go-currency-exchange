FROM golang:1.22.2-alpine
RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN CGO_ENABLED=1 go build -o server cmd/web/main.go

COPY internal/config/.env.local /app/internal/config/.env.local
COPY internal/database/migrations /app/internal/database/migrations
COPY main.db /app/main.db

CMD ["./server"]
