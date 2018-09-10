# API Exercise - Schools

Use any language and any frameworks/libraries you wish. Consult the internet as necessary.

1. Generate test data consisting of a list of schools:
    * each school will have an integer `id` and string `name`
    * you may keep the data in memory (no db required)
2. Create endpoint to list all schools
3. Create endpoint to get individual school by `id`
4. Create endpoint to add new school
5. Create endpoint to update an existing school (in this case only the `name` will get updated)
6. Handle looking up non-existent school
7. Host API docs available at `/docs/`

## Solution

1. n/a
2. `GET /schools` should return all schools
3. `GET /schools/:schoolId` should return a single school with id = schoolId
    * `404` when `schoolId` not found (after question 6)
4. `POST /schools` should add a new school and return the school added
    * also acceptable `POST /schools/:schoolId` which creates a school with a given `schoolId`
    * bonus if rejecting creating a duplicate school is handled
5. `PUT /schools/:schoolId` should update name associated with single school
    * `404` when `schoolId` not existing (after question 6)
6. Questions 3 and 5 should return `404`
7. Can document using any format (including hand rolled text)
    * bonus if using API documentation like swagger/openapi/raml/api blueprint

## Discussion

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
