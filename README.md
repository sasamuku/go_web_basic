# Go web basic app

### Create
```
curl -i -X POST -H "ContentType: application/json" -d '{"Content":"First Content", "author":"sasamuku"}' http://localhost:8080/post/
```

### Read
```
curl -i -X GET http://localhost:8080/post/1
```

### Update
```
curl -i -X PUT -H "Content-Type: application/json" -d '{"Content":"Modified Content", "author":"sasamuku"}' 
```

### Delete
```
curl -i -X DELETE http://localhost:8080/post/1
```
