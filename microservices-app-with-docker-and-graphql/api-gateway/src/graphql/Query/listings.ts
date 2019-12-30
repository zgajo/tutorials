interface Listing {
  id?: number;
  title?: string;
  description?: string;
  createdAt?: string;
  updatedAt?: string;
  deletedAt?: string;
}

const listingsResolver = async (): Promise<[Listing?]> => {
  try {
    return [
      {
        description: "TERTE"
      }
    ];
  } catch (error) {
    return [];
  }
};

export default listingsResolver;
