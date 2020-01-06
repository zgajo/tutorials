import { USER_SESSION_DELETE } from "../../../types";
import UsersService from "../../../adapters/UsersService";

const deleteUserSessionResolver = async (
  _: any,
  { sessionId }: USER_SESSION_DELETE,
  context: any
) => {
  await UsersService.deleteUserSession({
    sessionId
  });

  context.res.clearCookie("userSessionId");

  return true;
};

export default deleteUserSessionResolver;
