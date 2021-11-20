FROM mysql:5.7
ENV MYSQL_ROOT_PASSWORD=password

COPY docker/database/sakila-schema.sql /docker-entrypoint-initdb.d/sakila-01-table.sql

COPY docker/database/sakila-data.sql /docker-entrypoint-initdb.d/salila-02-data.sql
