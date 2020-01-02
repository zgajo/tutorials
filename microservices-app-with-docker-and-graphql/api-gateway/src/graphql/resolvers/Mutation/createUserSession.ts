import { USER_INSERT } from "../../../types";
import UsersService from "../../../adapters/UsersService";

const createUserSessionResolver = async (
  _: any,
  { email, password }: USER_INSERT,
  context: any
) => {
  const userSession: { id: string } = await UsersService.createUserSession({
    email,
    password
  });

  context.res.cookie("userSessionId", userSession.id, { httpOnly: true });

  return userSession;
};

export default createUserSessionResolver;
