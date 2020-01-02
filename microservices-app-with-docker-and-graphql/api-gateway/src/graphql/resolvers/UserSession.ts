import { USER_SESSION_FETCHED } from "src/types";

const UserSession = {
  user: async (userSession: USER_SESSION_FETCHED) => {
    return userSession.user;
  }
};

export default UserSession;
