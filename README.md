# kamus_bahasa_indonesia



## Graphql
### Endpoint
```
https://kbbi-graphql.herokuapp.com/
```
### Request
```graphql
query{penjelasan(kata:"malam")
  {
    kata
    arti
  }
}
```
### Response
```graphql
{
  "data": {
    "penjelasan": {
      "arti": [
        "waktu setelah matahari terbenam hingga matahari terbit",
        "lilin (dipakai untuk membatik): pedagang itu menjual --, nila, dan benang",
        "massa plastis amorf yang berasal dari mineral, tumbuhan, dan hewan"
      ],
      "kata": "malam"
    }
  }
}
```

## RestApi
### Endpoint
```
https://kbbi-graphql.herokuapp.com/api/arti/{kata}
```
### Example
```
https://kbbi-graphql.herokuapp.com/api/arti/malam
```
