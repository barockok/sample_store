test_datastore:
	DATABASE_CONFIG="root@tcp(127.0.0.1:3306)/sample_store_test_db?parseTime=true" go test github.com/barockok/sample_store/store/datastore