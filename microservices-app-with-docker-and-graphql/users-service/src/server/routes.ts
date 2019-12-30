import { Router } from "express";
import { Response } from "express";
import { createQueryBuilder } from "typeorm";

const router = Router();

import { User } from "../entity/users";

router.get("/users", async (_, res: Response, __) => {
  const users = await createQueryBuilder(User, "user")
    .select("user")
    .getMany();

  return res.json(users);
});

export default router;
