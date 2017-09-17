# Boilerplate for REST API in go

## ToDo's

[ ] Write logger wrapper with scope
[ ] Add own migration logic
[ ] Add own seeding logic

### Operations

| Operation | Command      | Description |
| --------- | ------------ | ----------- |
| **serve** | `fresh *`    | Runs the app and restarts if a change happened |
| **run**   | `go run *`   | Runs the app |
| **build** | `go build *` | Builds the app. |

## Library's

- **Router** [mux](http://www.gorillatoolkit.org/pkg/mux)
- **Middlewares** [negroni](https://github.com/urfave/negroni)
- **ORM** [gorm](http://jinzhu.me/gorm/)
- **Logger** [logrus](https://github.com/sirupsen/logrus)
