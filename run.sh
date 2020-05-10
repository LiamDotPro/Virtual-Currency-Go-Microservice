#!/usr/bin/env bash

#exec docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=currencies -v pgvolume:/pgdata postgres
#exec docker run -d -p 4000:80 -e PGADMIN_DEFAULT_EMAIL=test@test.com -e PGADMIN_DEFAULT_PASSWORD=yoursecurepassword dpage/pgadmin4

exec go build
exec ./Virtual-Currency-Go-Microservice
