# sqlc-example

Example webapp using sqlc

## Build

```
$ echo 'DATABASE_URL="mysql://user:password@127.0.0.1:3306/dbname"' > .env
$ dbmate up
$ sqlc generate
$ DSN="user:password@tcp(localhost:3306)/dbname" go build
```

## License

sqlc-example is published under the MIT License, see LICENSE.

## Author
[Noriaki Watanabe@nnabeyang](https://twitter.com/nnabeyang)
