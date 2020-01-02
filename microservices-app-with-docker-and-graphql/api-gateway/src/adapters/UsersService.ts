import got from "got";
import { USERS_SERVICE_URI } from "../helpers/constants";
import { USER_INSERT, USER_SESSION_FETCHED } from "../types";

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

  static async fetchUser({
    userId
  }: USER_SESSION_FETCHED): Promise<{ id: string }> {
    const body: { id: string } = await got
      .get(`${USERS_SERVICE_URI}/users/${userId}`)
      .json();

    return body;
  }

  static async fetchUserSession({
    sessionId
  }: {
    sessionId: string;
  }): Promise<{ id: string }> {
    const body: { id: string } = await got
      .get(`${USERS_SERVICE_URI}/sessions/${sessionId}`)
      .json();

    return body;
  }
}
