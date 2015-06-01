var express = require("express");
var server = express();
function Index(req, res) {
  res.send("ok from node!")
}
server.get("/", Index)
server.listen("8081");
