mkdir go-blog
cd go-blog

mkdir -p cmd internal/{appname} pkg/db configs migrations scripts

touch go.mod cmd/main.go configs/config.go scripts/migrate.sh Dockerfile docker-compose.yml .env

go mod init github.com/yourusername/go-blog



