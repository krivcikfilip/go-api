# GO Api

my first touch with go language

### Technologies
- [**Fiber**](https://docs.gofiber.io/) Web framework
- [**GORM**](https://gorm.io/) ORM for working with database
- [**Air**](https://github.com/cosmtrek/air) Live reload
- [**PostgreSQL**](https://www.postgresql.org/) Database

### How to run

#### Docker
> docker compose up

#### With Air
> air

#### Without Air
> go run main.go

### Endpoints
prefix **/api**

#### Categories
| Method | Route           | Body                 |
|--------|-----------------|----------------------|
| GET    | /categories     |                      |
| GET    | /categories/:id |                      |
| POST   | /categories     | `{"name": "scifi" }` |
| PUT    | /categories/:id | `{"name": "drama" }` |

#### Books
| Method | Route      | Body                                                  |
|--------|------------|-------------------------------------------------------|
| GET    | /books     |                                                       |
| GET    | /books/:id |                                                       |
| POST   | /books     | `{"name": "harry potter"`                             |
| PUT    | /books/:id | `{"name": "harry potter 2", "categories": [1, 2] } }` |