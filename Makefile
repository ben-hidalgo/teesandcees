
docker-tcapi:
	docker-compose up --build tcapi

docker-protos:
	docker-compose up --build protos

docker-testtcapi:
	mkdir -p testtcapi/protos/
	cp protos/tcapi.proto testtcapi/protos/
	docker-compose up --build testtcapi
