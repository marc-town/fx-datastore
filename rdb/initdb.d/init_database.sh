cat docker-entrypoint-initdb.d/001-create-database.sql | mysql -u root -p
cat docker-entrypoint-initdb.d/002-create-tables.sql | mysql -u root -p
