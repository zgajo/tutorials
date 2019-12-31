import { Router } from "express";
import { Response, Request, NextFunction } from "express";
import { createQueryBuilder, getRepository } from "typeorm";

const router = Router();

import { User } from "../entity/users";
import { hashPassword } from "../helpers/hashPassword";

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
        await getRepository(User).insert({
          email: req.body.email,
          passwordHash: hashPassword(req.body.password)
        })
      );
    } catch (error) {
      next(error);
    }
  }
);

export default router;
