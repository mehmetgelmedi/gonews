up: dep go

dep:
	cd frontend/ && npm install && npm run build

go:
	go get && go run main.go