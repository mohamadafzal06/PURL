dbup:
	docker-compose up -d
build:
	go build .
run:
	./purl

# run it in repository/postgres
# #################
#migration-status:
# sql-migrate status -env="development" -config=dbconfig.yml
#
#migration-up:
# sql-migrate up -env="development" -config=dbconfig.yml
## if dont use limit, all migration down called
#migration-down:
# sql-migrate down -env="development" -config=dbconfig.yml -limit=<n>

################
