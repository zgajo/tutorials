import got from "got";
import { LISTINGS_SERVICE_URI } from "../helpers/constants";

export default class ListingsService {
  static async fetchAllListings(): Promise<[]> {
    const body: [] = await got.get(`${LISTINGS_SERVICE_URI}/listings`).json();

    return body;
  }
}
