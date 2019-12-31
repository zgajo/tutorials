import get from "lodash/get";
import { GraphQLFormattedError } from "graphql";

export const formatGraphQLErrors = (error: GraphQLFormattedError) => {
  const errorDetails = get(error, "originalError.response.body");

  try {
    if (error) {
      return JSON.parse(errorDetails);
    }
  } catch (error) {}

  return error;
};
