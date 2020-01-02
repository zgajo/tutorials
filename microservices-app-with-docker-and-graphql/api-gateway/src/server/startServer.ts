import { ApolloServer } from "apollo-server-express";

import cookieParser from "cookie-parser";
import cors from "cors";
import express from "express";

import accessEnv from "../helpers/accessEnv";
import typeDefs from "../graphql/typeDefs";
import resolvers from "../graphql/resolvers";
import { formatGraphQLErrors } from "./formatGraphQLErrors";
import injectSession from "./injectSession";

const PORT = accessEnv("PORT", 7000);

const apolloServer = new ApolloServer({
  context: a => a,
  resolvers,
  typeDefs,
  formatError: formatGraphQLErrors
});

const app = express();

app.use(cookieParser());

app.use(
  cors({
    origin: (_, cb) => cb(null, true),
    credentials: true
  })
);

app.use(injectSession);

apolloServer.applyMiddleware({ app, cors: false, path: "/graphql" });

app.listen(PORT, "0.0.0.0", () => {
  console.log(`API gateway running on port: ${PORT}`);
});
