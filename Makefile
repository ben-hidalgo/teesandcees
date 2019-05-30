local-tcapi:
	(cd tcapi && \
		go fmt ./... && \
		LISTEN_ADDRESS=":8000" \
		REDIS_ADDRESS="localhost:6379" \
		REDIS_PASSWORD="" \
		REDIS_DB="0" \
		go run main.go)

docker-tcapi:
	docker-compose up --build tcapi

docker-protos:
	docker-compose up --build protos

docker-redis:
	docker-compose up --build redis

docker-sharedb:
	docker-compose up --build sharedb

docker-mongo:
	docker-compose up --build mongo

docker-testtcapi:
	mkdir -p testtcapi/protos/
	cp protos/tcapi.proto testtcapi/protos/
	docker-compose up --build testtcapi
