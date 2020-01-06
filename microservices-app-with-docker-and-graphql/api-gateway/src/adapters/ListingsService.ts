import got from "got";
import { LISTINGS_SERVICE_URI } from "../helpers/constants";

export default class ListingsService {
  static async fetchAllListings(): Promise<[]> {
    const body: [] = await got.get(`${LISTINGS_SERVICE_URI}/listings`).json();

    return body;
  }
  static async createListing({ title, description }: any): Promise<[]> {
    const body: any = await got
      .post(`${LISTINGS_SERVICE_URI}/listings`, {
        json: {
          description,
          title
        }
      })
      .json();

    return body;
  }
}
