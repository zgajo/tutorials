const userSessionResolver = async (_: any, args: any, context: any) => {
  if (args.me !== true) {
    throw new Error("Unsuported argument value");
  }

  return context.res.locals.userSession;
};

export default userSessionResolver;
