# Designing GraphQL Schemas - Nik Graf

https://www.youtube.com/watch?v=fBkmlFfwRu0&list=WL&index=39&t=8s

https://facebook.github.io/relay/graphql/connections.htm

Well the point is, if you implement a ConnectionType for an endpoint you just need to

- accept connectionArgs parameters
- return a suitable response in the shape { edges: [{ cursor, node }], pageInfo: { hasNextPage, hasPreviousPage, startCursor, endCursor }

Cursors can be any string you want. Normally Relay defaults to arrayconnection:<id> converted to base64 to make it opaque (you shouldn't actually relay on the implementation details of the cursors).

I used <type_name>:<id> converted to base64 as cursors.

Then a request to my endpoint look like this

```
endpoint(after: "sdflkjsdlfkjslkdf", first: 10) {
   # stuff
}
```

Basically, any request described by the specification is supported.

Thus when I get the request I process it in the following way:

Start from the greedy query: SELECT \* FROM table

1. If the after argument is provided, add id > parsed_cursor to the WHERE clause

1. If the before argument is provided, add id < parsed_cursor to the WHERE clause

1. If the first argument is provided, add ORDER BY id DESC LIMIT first+1 to the query

1. If the last argument is provided, add ORDER BY id ASC LIMIT last+1 to the query

1. If the last argument is provided, I reverse the order of the results

1. If the first argument is provided then I set hasPreviousPage: false (see spec for a description of this behavior).

1. If no less than first+1 results are returned, I set hasNextPage: true, otherwise I set it to false.

1. If the last argument is provided then I set hasNextPage: false (see spec for a description of this behavior).

1. If no less last+1 results are returned, I set hasPreviousPage: true, otherwise I set it to false.

Using this "algorithm", only the needed data is fetched. While after and before can be both set, I make sure first and last args are treated as mutually exclusive to avoid making a mess. The spec itself discourage making requests with both the arguments set.

Notably, I return an object with the shape described above (and in the linked spec) and I don't use the connectionFromArray helper, which expects a raw collection to slice accordingly to the args.
