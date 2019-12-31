import got from "got";
import { USERS_SERVICE_URI } from "../helpers/constants";
import { USER_INSERT } from "../types";

export default class UsersService {
  static async createUser({ email, password }: USER_INSERT): Promise<[]> {
    const body: [] = await got
      .post(`${USERS_SERVICE_URI}/users`, {
        json: {
          email,
          password
        }
      })
      .json();

    return body;
  }

  static async createUserSession({
    email,
    password
  }: USER_INSERT): Promise<{ id: string }> {
    const body: { id: string } = await got
      .post(`${USERS_SERVICE_URI}/sessions`, {
        json: {
          email,
          password
        }
      })
      .json();

    return body;
  }
}
