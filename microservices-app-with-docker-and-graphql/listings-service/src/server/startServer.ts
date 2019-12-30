import bodyParser from "body-parser";
import cors from "cors";
import express from "express";
import { createConnection } from "typeorm";

import accessEnv from "../helpers/accessEnv";

const PORT = accessEnv("PORT", 7100);

const app = express();

app.use(bodyParser());

app.use(
  cors({
    origin: (_, cb) => cb(null, true),
    credentials: true
  })
);

app.listen(PORT, "0.0.0.0", async () => {
  await createConnection();
  console.log(`Listings services running on port: ${PORT}`);
});
