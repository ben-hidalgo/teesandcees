
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
