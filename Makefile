pullpostgres:
	docker pull postgres:12-alpine

postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root amado
	
dropdb:
	docker exec -it postgres12 dropdb --username=root amado

dbup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/amado?sslmode=disable" -verbose up

dbdown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/amado?sslmode=disable" -verbose down