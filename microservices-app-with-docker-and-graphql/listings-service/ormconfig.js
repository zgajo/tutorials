module.exports = {
  type: "mysql",
  host: "localhost",
  port: 3306,
  username: process.env.TYPEORM_USERNAME,
  password: process.env.TYPEORM_PASSWORD,
  database: "db",
  synchronize: false,
  logging: true,
  entities: ["src/entity/**/*.ts"],
  migrations: ["src/migration/**/*.ts"],
  subscribers: ["src/subscriber/**/*.ts"],
  seeds: ["src/seeds/**/*.seed.ts"],
  factories: ["src/factories/**/*.factory.ts"],
  charset: "utf8mb4_unicode_ci",
  cli: {
    entitiesDir: "src/entity",
    migrationsDir: "src/migration",
    subscribersDir: "src/subscriber"
  }
};
