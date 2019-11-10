var createError = require("http-errors");
var express = require("express");
var cookieParser = require("cookie-parser");
var logger = require("morgan");
const cors = require("cors");
const ytdl = require("ytdl-core");

var app = express();

app.use(express.static(__dirname + "public"));
app.use(cors());

app.use(logger("dev"));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(__dirname + "/public"));

app.get("/", function(req, res) {
  res.sendFile(__dirname + "/public/index.html");
});

app.get("/download", (req, res) => {
  console.log("object");
  var URL = req.query.URL;

  var URL = req.query.URL;
  res.header("Content-Disposition", 'attachment; filename="video.mp4"');

  ytdl(URL, {
    format: "mp4"
  }).pipe(res);
});

// catch 404 and forward to error handler
app.use(function(req, res, next) {
  next(createError(404));
});

// error handler
app.use(function(err, req, res, next) {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = req.app.get("env") === "development" ? err : {};

  // render the error page
  res.status(err.status || 500);
  res.render("error");
});

module.exports = app;
