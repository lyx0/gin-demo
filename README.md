GET

```
curl http://localhost:8080/products
```

POST

```
curl http://localhost:8080/products \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": 3, "name": "tea", "description": "nice cuppa tea", "price": 1.29}'

```
