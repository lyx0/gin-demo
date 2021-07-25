List all products

```
curl http://localhost:8080/products
```

Get a product with id

```
curl localhost:8080/products/2
```

Add a product

```
curl http://localhost:8080/products \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "3", "name": "tea", "description": "nice cuppa tea", "price": 1.29}'

```
