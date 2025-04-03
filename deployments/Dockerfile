# Go 빌드 환경
FROM golang:1.24 AS builder

WORKDIR /app

# 먼저 go.mod와 go.sum만 복사하여 의존성 다운로드
COPY go.mod go.sum ./
RUN go mod tidy

# 이후 모든 소스코드를 복사
COPY . .

# 빌드 실행
RUN go build -o main ./cmd/api/main.go

# 실행 환경
FROM golang:1.24
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
