
## Create user test for mongodb
```
use users_db
db.createUser(
  {
    user: "user_test",
    pwd: "user_password",
    roles: [ { role: "readWrite", db: "users_db" } ]
  }
)

use users_db
db.auth("user_test", "user_password")
```
## Install
* `git clone https://github.com/tinhn/restfull-go-mux.git`


## Start service
* Change config in `mongo-cfg.toml` file
* Run `go run main.go`

## Test service
```
curl -H 'Content-Type: application/json' -L -X GET 'http://localhost:3333/users'
curl -H 'Content-Type: application/json' -L -X POST 'http://localhost:3333/users' --data-raw '{"name":"tinhn","age":18,"email":"tinhn@gmail.org"}'
curl -H 'Content-Type: application/json' -L -X GET 'http://localhost:3333/users/5f227fd610eb695a25e4faad'
curl -H 'Content-Type: application/json' -L -X PUT 'http://localhost:3333/users/5f227fd610eb695a25e4faad' --data-raw '{"name":"tinhnguyen","age":20,"email":"tinhn@gmail.gmail"}'
curl -H 'Content-Type: application/json' -L -X DELETE 'http://localhost:3333/users' --data-raw '{"id":"5f227c7010eb695a25e4faab","name":"tinhnguyen","age":18,"email":"tinhn@sendo.vn"}'

```
