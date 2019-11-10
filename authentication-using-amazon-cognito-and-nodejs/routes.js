var express = require("express");

var router = express.Router();

var authController = require("./AuthController");
var helloController = require("./HelloController");
var authMiddleware = require("./AuthMiddleware");

router.post("/auth/register", authController.register);
router.post("/auth/login", authController.login);
router.post("/auth/validate", authController.validate_token);
router.get("/hello", authMiddleware.Validate, helloController.simple_hello);

module.exports = router;
