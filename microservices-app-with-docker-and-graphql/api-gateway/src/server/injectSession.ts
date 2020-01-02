import UsersService from "../adapters/UsersService";

const injectSession = async (req: any, res: any, next: any) => {
  if (req.cookies.userSessionId) {
    try {
      const userSession = await UsersService.fetchUserSession({
        sessionId: req.cookies.userSessionId
      });

      res.locals.userSession = userSession;
    } catch (error) {
      return next();
    }
  }

  return next();
};

export default injectSession;
