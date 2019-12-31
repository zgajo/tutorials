import ListingsService from "../../adapters/ListingsService";

interface Listing {
  id?: number;
  title?: string;
  description?: string;
  createdAt?: string;
  updatedAt?: string;
  deletedAt?: string;
}

const listingsResolver = async (): Promise<[Listing?]> => {
  return await ListingsService.fetchAllListings();
};

export default listingsResolver;
