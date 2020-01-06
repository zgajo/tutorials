import { Router } from "express";
import { Response, Request, NextFunction } from "express";
import { createQueryBuilder } from "typeorm";

const router = Router();

import { Listing } from "../entity/listings";

router.get("/listings", async (_, res: Response, __) => {
  const listings = await createQueryBuilder(Listing, "listing")
    .select("listing")
    .getMany();

  return res.json(listings);
});

router.post(
  "/listings",
  async (req: Request, res: Response, next: NextFunction) => {
    if (!req.body.description || !req.body.title) {
      return next(new Error("Invalid body"));
    }

    try {
      const listing = await Listing.create({
        description: req.body.description,
        title: req.body.title
      }).save();

      return res.json(listing);
    } catch (error) {
      next(error);
    }
  }
);

export default router;
