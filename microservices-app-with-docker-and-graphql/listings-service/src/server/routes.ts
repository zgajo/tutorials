import { Router } from "express";
import { Response } from "express";
import { createQueryBuilder } from "typeorm";

const router = Router();

import { Listing } from "../entity/listings";

router.get("/listings", async (_, res: Response, __) => {
  const listings = await createQueryBuilder(Listing, "listing")
    .select("listing")
    .getMany();

  return res.json(listings);
});

export default router;
