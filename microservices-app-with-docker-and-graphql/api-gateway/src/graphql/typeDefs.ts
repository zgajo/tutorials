import { gql } from "apollo-server";

const typeDefs = gql`
  scalar Date

  type Listing {
    description: String!
    id: ID!
    title: String!
  }

  type User {
    id: ID!
    email: String!
  }

  type UserSession {
    id: ID!
    createdAt: Date!
    expiresAt: Date!
    user: User!
  }

  type Query {
    listings: [Listing!]!
    userSession(me: Boolean!): UserSession
  }

  type Mutation {
    createUser(email: String!, password: String!): User!
    createUserSession(email: String!, password: String!): UserSession
  }
`;

export default typeDefs;
