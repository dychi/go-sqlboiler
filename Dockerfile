FROM golang:1.19.5-bullseye

CMD ["go", "run", "cmd/api/main.go"]