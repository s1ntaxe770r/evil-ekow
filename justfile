
run:
   go run main.go

# start a redis container
db:
   docker run -p 6379:6379 --name reedis -d redis:alpine
