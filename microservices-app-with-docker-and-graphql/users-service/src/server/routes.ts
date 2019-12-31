import { NextFunction, Request, Response, Router } from "express";
import { createQueryBuilder } from "typeorm";
import { addHours } from "date-fns";

import { User } from "../entity/users";
import { hashPassword } from "../helpers/hashPassword";
import { compareSync } from "bcrypt";
import { USER_SESSION_EXPIRY_HOURS } from "../helpers/constants";
import { UserSessions } from "../entity/userSessions";

const router = Router();

router.post("/sessions", async (req, res, next) => {
  if (!req.body.email || !req.body.password) {
    return next(new Error("Invalid body"));
  }

  try {
    const user = await User.findOne({ where: { email: req.body.email } });

    if (!user) {
      return next(new Error("Invalid email"));
    }

    if (!compareSync(req.body.password, user.passwordHash)) {
      return next(new Error("Incorecct password"));
    }

    const expiresAt = addHours(new Date(), USER_SESSION_EXPIRY_HOURS);

    const userSession = await UserSessions.create({
      user: user,
      expiresAt: expiresAt,
      createdAt: new Date()
    }).save();

    return res.json(userSession);
  } catch (error) {
    return next(error);
  }
});

router.get("/users", async (_, res: Response, __) => {
  const users = await createQueryBuilder(User, "user")
    .select("user")
    .getMany();

  return res.json(users);
});

router.post(
  "/users",
  async (req: Request, res: Response, next: NextFunction) => {
    if (!req.body.email || !req.body.password) {
      return next(new Error("Invalid body"));
    }

    try {
      return res.json(
        await User.create({
          email: req.body.email,
          passwordHash: hashPassword(req.body.password)
        }).save()
      );
    } catch (error) {
      return next(error);
    }
  }
);

export default router;
