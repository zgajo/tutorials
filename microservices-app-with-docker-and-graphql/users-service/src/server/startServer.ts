import bodyParser from "body-parser";
import cors from "cors";
import express, { ErrorRequestHandler } from "express";
import { createConnection } from "typeorm";
import accessEnv from "../helpers/accessEnv";
import routes from "./routes";

const PORT = accessEnv("PORT", 7101);

const app = express();

app.use(bodyParser.json());

app.use(
  cors({
    origin: (_, cb) => cb(null, true),
    credentials: true
  })
);

app.use("/", routes);

const errorMiddlware: ErrorRequestHandler = (err, _, res, __) => {
  return res.status(500).json({
    message: err.message
  });
};

app.use(errorMiddlware);

app.listen(PORT, "0.0.0.0", async () => {
  await createConnection();

  console.log(`Users services running on port: ${PORT}`);
});
