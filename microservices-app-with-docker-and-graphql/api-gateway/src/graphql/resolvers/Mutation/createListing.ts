import ListingsService from "../../../adapters/ListingsService";

const createListingResolver = async (_: any, { title, description }: any) => {
  return await ListingsService.createListing({ title, description });
};

export default createListingResolver;
