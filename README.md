# CrudGo

It is a example to Api Crund in go with gin-gonic in hexagonal Architecture. Create .env with this:
```
  MONGO_URI=
  MONGO_DBNAME=
  MONGO_COLLECTION=
  API_TOKEN=
  PORT=
  ELEMENTS_TO_PAGINATE=
```
 
The data mongo structure is:

```json
{
  "id":22,
  "name":"Hong Xiang",
  "email":"hongxiang17@gmail.com",
  "age":27,
  "entryDate":{"$date":"2018-06-25T21:46:44.000Z"},
  "country":"Chile"
}
 ```

Endpoint
 '''
GET    /user?page=x <page default is 0>
GET    /user/id/:id              
GET    /search/:field/:value <page default is 0>
DELETE /user/:id
POST   /user/
 '''
 

Commands in Make
```
  run
  install
  build
  docker-build
  docker-run
  unit-testing
```

For local run project:
```
 Make run
```
And load to data example in folder db
 
