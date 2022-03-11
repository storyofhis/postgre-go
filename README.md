# postgre-go

installation: 
```shellsession
$ go get github.com/go-sql-driver/mysql github.com/jinzhu/gorm github.com/jinzhu/gorm/dialects/postgres
```

.env
```go
export DIALECT="postgres"
export HOST="localhost"
export DBPORT="5432"
export USER="postgres"
export NAME="book_keeper"
export PASSWORD="@bc123"
```

CREATE TABLE Person: 
```go
type Person struct {
	gorm.Model

	Name string
	Email string `gorm:"typevarchar(100);unique_index"`
	Books []Book
}
```

CREATE TABLE Book: 
```go
type Book struct {
	gorm.Model

	Title string
	Author string
	CallNumber int `gorm:"unique_index"`
	PersonID int
}
```
