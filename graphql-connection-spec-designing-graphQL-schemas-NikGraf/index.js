const { ApolloServer, gql } = require("apollo-server");

// A schema is a collection of type definitions (hence "typeDefs")
// that together define the "shape" of queries that are executed against
// your data.
const typeDefs = gql`
  # Comments in GraphQL strings (such as this one) start with the hash (#) symbol.

  # This "Book" type defines the queryable fields for every book in our data source.
  type Product {
    id: ID!
    title: String
    author: String
    image: Image
    recommendedProducts(limit: Int, after: Int): RecommendedConnection
  }

  type RecommendedConnection {
    edges: [RelatedProduct]!
    pageInfo: PageInfo
  }

  type PageInfo {
    startCursor: Int
    endCursor: Int
    hasMoreItems: Boolean
  }

  type RelatedProduct {
    product: Product
    boughtTogetherPercentage: Float
  }

  type Image {
    id: ID!
    url: String
    title: String
    thumbnailUrl(width: Int, height: Int): String
  }

  # The "Query" type is special: it lists all of the available queries that
  # clients can execute, along with the return type for each. In this
  # case, the "books" query returns an array of zero or more Books (defined above).
  type Query {
    products: [Product]!
    product(id: Int!): Product
  }
`;

const products = [
  {
    id: 1,
    title: "Harry Potter and the Chamber of Secrets",
    author: "J.K. Rowling",
    imageId: 1,
    recommendedProductIds: [3, 4, 5, 6, 7]
  },
  {
    id: 2,
    title: "Jurassic Park",
    author: "Michael Crichton",
    imageId: 2,
    recommendedProductIds: [8]
  },
  {
    id: 3,
    title: "Recc prod 3",
    author: "Michael Crichton",
    imageId: 2
  },
  {
    id: 4,
    title: "Recc prod 4",
    author: "Michael Crichton",
    imageId: 2
  },
  {
    id: 5,
    title: "Recc prod 5",
    author: "Michael Crichton",
    imageId: 2
  },
  {
    id: 6,
    title: "Recc prod 6",
    author: "Michael Crichton",
    imageId: 2
  },
  {
    id: 7,
    title: "Recc prod 7",
    author: "Michael Crichton",
    imageId: 2
  },
  {
    id: 8,
    title: "Recc prod 8",
    author: "Michael Crichton",
    imageId: 2
  }
];

const images = [
  {
    id: 1,
    url: "url 1",
    title: "J.K. Rowling image"
  },
  {
    id: 2,
    url: "url 2",
    title: "Michael Crichton image"
  }
];

const resolvers = {
  Query: {
    products: () => products,
    product: (_, { id }) => {
      return products.find(prod => prod.id === id);
    }
  },
  Product: {
    image: ({ imageId, recommendedProducts }) =>
      images.find(img => img.id === imageId),
    recommendedProducts: ({ recommendedProductIds }, args) => {
      const recProds = products.filter(({ id }) =>
        recommendedProductIds.includes(id)
      );

      return {
        recommendedProducts: recProds,
        arguments: args
      };
    }
  },
  Image: {
    thumbnailUrl: (parentImage, { width, height }) => ""
  },
  RecommendedConnection: {
    edges: ({ recommendedProducts, arguments }) => {
      const { limit, after } = arguments;

      if (after && limit) {
        const afterId =
          recommendedProducts.findIndex(rp => rp.id === after) + 1;

        const slicedRecProd = recommendedProducts.slice(
          afterId,
          afterId + limit
        );

        return slicedRecProd.map(prod => ({
          product: prod
        }));
      }

      return recommendedProducts.map(prod => ({
        product: prod
      }));
    },
    pageInfo: ({ recommendedProducts, arguments }) => {
      const { limit, after } = arguments;

      let startCursor = recommendedProducts.length
        ? recommendedProducts[0].id
        : null;

      let endCursor = recommendedProducts.length
        ? recommendedProducts[recommendedProducts.length - 1].id
        : null;

      let hasMoreItems = false;

      if (limit && after) {
        const afterId =
          recommendedProducts.findIndex(rp => rp.id === after) + 1;

        const slicedRecProd = recommendedProducts.slice(
          afterId,
          afterId + limit
        );

        startCursor =
          slicedRecProd && slicedRecProd.length ? slicedRecProd[0].id : null;

        endCursor =
          slicedRecProd && slicedRecProd.length
            ? slicedRecProd[slicedRecProd.length - 1].id
            : null;

        hasMoreItems = !!recommendedProducts.slice(afterId + limit).length;
      }

      return {
        startCursor,
        endCursor,
        hasMoreItems
      };
    }
  },
  RelatedProduct: {
    boughtTogetherPercentage: p => {
      console.log(p);
    }
  }
};

// The ApolloServer constructor requires two parameters: your schema
// definition and your set of resolvers.
const server = new ApolloServer({ typeDefs, resolvers });

// The `listen` method launches a web server.
server.listen().then(({ url }) => {
  console.log(`🚀  Server ready at ${url}`);
});
