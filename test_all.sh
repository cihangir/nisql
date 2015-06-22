#!/bin/sh
set -e

echo "Testing against mysql"
export NISQL_TEST_DSN=nisqltest:nisqltest@/nisqltest
export NISQL_TEST_DIALECT=mysql
go test ./...

echo "Testing against postgres"
export NISQL_TEST_DSN="user=nisqltest password=nisqltest dbname=nisqltest sslmode=disable"
export NISQL_TEST_DIALECT=postgres
go test ./...

echo "Testing against sqlite"
export NISQL_TEST_DSN=/tmp/nisqltest.bin
export NISQL_TEST_DIALECT=sqlite
go test ./...
