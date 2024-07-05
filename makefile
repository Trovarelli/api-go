create_migration:
	migrate create -ext=sql -dir=internal/database/migrations -seq init
gen_tables:
	migrate -path=internal/database/migrations -database "postgresql://postgres:1234@localhost:5432/postgres?sslmode=disable" -verbose up 
drop_datbles:
	migrate -path=internal/database/migrations -database "postgresql://postgres:1234@localhost:5432/postgres?sslmode=disable" -verbose down

.PHONY: create_migration migrate_up migrate_down
