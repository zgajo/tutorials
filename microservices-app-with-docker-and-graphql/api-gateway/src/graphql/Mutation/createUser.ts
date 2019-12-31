import { USER_INSERT } from "../../types";
import UsersService from "../../adapters/UsersService";

const createUserResolver = async (_: any, { email, password }: USER_INSERT) => {
  return await UsersService.createUser({ email, password });
};

export default createUserResolver;
