migration-users:
	@read -p "input the migration name: " name; \
	migrate create -ext sql -dir services/users/entities/databases/migrations $$name