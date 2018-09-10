const express = require("express");
const app = express();

app.use(express.json());

const swaggerUi = require("swagger-ui-express");
const swaggerDocument = require("../swagger/swagger.json");

app.use("/docs", swaggerUi.serve, swaggerUi.setup(swaggerDocument));

let schools = [1, 2, 3].map(id => ({
  id,
  name: `School${id}`
}));

app.get("/schools", function(req, res) {
  return res.send(schools);
});

app.get("/schools/:schoolId", function(req, res) {
  const school = schools.find(school => school.id == req.params.schoolId);
  if (!school) {
    res.status(404);
    res.send();
    return;
  }

  return res.send(school);
});

app.post("/schools", function(req, res) {
  const school = {
    id: req.body.id,
    name: req.body.name
  };
  schools.push(school);
  return res.send(school);
});

app.put("/schools/:schoolId", function(req, res) {
  const school = schools.find(school => school.id == req.params.schoolId);
  if (!school) {
    res.status(404);
    res.send();
    return;
  }

  school.name = req.body.name;
  res.send(school);
});

app.listen(8080);
