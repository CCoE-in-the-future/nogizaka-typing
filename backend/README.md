## Nogizaka Typing Backend

### Run

```bash
cd backend
go mod tidy
GOOS=linux GOARCH=amd64 go build -o main main.go
sam local start-api
```
