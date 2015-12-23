# Gorm With Rows

## Quick Started 

    # Get code
    $ go get -u github.com/fengjh/gorm-with-rows

    # Set environment variables
    $ source .env

    # Setup postgres database
    $ postgres=# CREATE USER gorm_with_rows;
    $ postgres=# CREATE DATABASE gorm_with_rows OWNER gorm_with_rows;

    # Run Application
    $ go test -bench=.

## Results

    // SQL
    $ [0.63ms]  SELECT  id, title, description, created_at, updated_at, deleted_at FROM "products"  WHERE ("products".deleted_at IS NULL OR "products".deleted
    $ [177.59ms]  SELECT  * FROM "products"  WHERE ("products".deleted_at IS NULL OR "products".deleted_at <= '0001-01-02')

    // Benchmark log
    $ BenchmarkLoadProductsWithRows-8	      30	  44941920 ns/op
    $ BenchmarkLoadProductsWithFind-8	      10	 176184201 ns/op
    $ ok  	github.com/fengjh/gorm-with-rows	3.847s
