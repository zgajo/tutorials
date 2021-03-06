import { gql } from "apollo-server";

const typeDefs = gql`
  scalar Date

  type Listing {
    id: ID!
    title: String!
    description: String!
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
    createListing(title: String!, description: String!): Listing!
    createUserSession(email: String!, password: String!): UserSession
    deleteUserSession(sessionId: ID!): Boolean
  }
`;

export default typeDefs;
