# CrudGo

Este es un ejemplo sencillo de un CRUD en go sobre una base de mongo. Se deben configurar las siguientes variables de entorno:

Variables de entornos.
 ```
MONGO_URI=mongodb:
MONGO_DBNAME=
MONGO_COLLECTION=
API_TOKEN=
PORT=
 ```
 
La estructura utilizada en el mongo es:

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
 
 
Los endpoint:
```
/user/id/:id              
/user/country/:country
```
 
 Los comandos de make son:
 ```
 run
 install
 build
 docker-build
 docker-run
 ```
 
