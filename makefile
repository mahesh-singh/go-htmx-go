build:
	go build -o bin/gohtmxgo.exe
run: build #depends on build command
	bin/gohtmxgo.exe
postgres:
	docker run --name postgres-container -e POSTGRES_PASSWORD=password -d -p 5432:5432 postgres
