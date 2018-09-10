# API Exercise - Schools

Use any language and any frameworks/libraries you wish. Consult the internet as necessary.

1. Generate test data consisting of a list of schools:
    * each school should have an integer `id` and string `name`
2. endpoint to list all schools on `GET /schools`
3. endpoint to get individual school `GET /schools/:schoolId`
4. endpoint to add new school `POST /schools`
5. endpoint to update a school `name` with a given `id`
6. what to do when looking up non-existent school by `id`?
7. return api docs `GET /docs/`
8. discuss
    * optimizing for read heavy use
    * optimizing for write heavy use
    * how to handle concurrent writes/updates
    * how to ensure all users get service when one bad actor may overwhelm requests
    * approaches to authentication
    * approaches to logging, observability
    * how to create client lib / sdk
      * when would you provide one? what are advantages / disadvantages?
    * approaches to testing this api


## Testing Endpoints

The following tests assume port to be 8080.

### Get all schools
```bash
curl http://localhost:8080/schools
```

### Get individual school
```bash
curl http://localhost:8080/schools/1
```

### Add new school
```bash
curl -d '{"id": 92, "name": "School92"}' -H "Content-Type: application/json" -X POST  http://localhost:8080/schools
```

### Update school name
```bash
curl -d '{"name": "School9000"}' -H "Content-Type: application/json" -X PUT  http://localhost:8080/schools/1
```

### Update non-existant school
```
curl -v -d '{"name": "School9000"}' -H "Content-Type: application/json" -X PUT  http://localhost:8080/schools/100
```

### Get non-existent school
```bash
curl -I -X GET http://localhost:8080/schools/999
```

### Get api docs
```bash
curl http://localhost:8080/docs/
```
