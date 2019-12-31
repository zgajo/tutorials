import { NextFunction, Request, Response, Router } from "express";
import { createQueryBuilder } from "typeorm";
import { User } from "../entity/users";
import { hashPassword } from "../helpers/hashPassword";

const router = Router();

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
      next(error);
    }
  }
);

export default router;
