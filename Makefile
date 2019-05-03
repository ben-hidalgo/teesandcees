
docker-tcapi:
	docker-compose up --build tcapi

docker-protos:
	docker-compose up --build protos

docker-testtcapi: docker-protos
	cp protos/artifacts/ruby/tcapi_pb.rb testtcapi/lib/
	docker-compose up --build testtcapi
