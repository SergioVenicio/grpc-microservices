DATA_SOURCE_URL="root:verysecretpass@tcp(127.0.0.1:3306)/order?parseTime=true" \
PAYMENT_SERVICE_URL=localhost:6000 \
APPLICATION_PORT=5000 \
ENV=development \
go run ./cmd/main.go