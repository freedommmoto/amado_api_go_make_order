
CREATE TABLE "order" (
                         "id_order" serial PRIMARY KEY,
                         "first_name" varchar NOT NULL,
                         "last_name" varchar NOT NULL,
                         "email" varchar NOT NULL,
                         "phone" varchar NOT NULL,
                         "address" text,
                         "zip_code" text,
                         "created_at" timestamp NOT NULL DEFAULT 'now()',
                         "modified" timestamp NOT NULL DEFAULT 'now()',
                         "deleted" timestamp
);

CREATE TABLE "order_product" (
                                 "id_order_product" bigserial PRIMARY KEY,
                                 "id_order" int,
                                 "id_product" int,
                                 "qty" int NOT NULL,
                                 "price_per_qty" int NOT NULL,
                                 "created_at" timestamp NOT NULL DEFAULT 'now()',
                                 "modified" timestamp NOT NULL DEFAULT 'now()',
                                 "deleted" timestamp
);

CREATE TABLE "product" (
                           "id_product" serial PRIMARY KEY,
                           "name" text,
                           "stock" int NOT NULL,
                           "price" int NOT NULL,
                           "active" boolean DEFAULT true,
                           "created_at" timestamp NOT NULL DEFAULT 'now()',
                           "modified" timestamp NOT NULL DEFAULT 'now()',
                           "deleted" timestamp
);

CREATE TABLE "users" (
                         "id_user" serial PRIMARY KEY,
                         "username" varchar,
                         "password" varchar NOT NULL,
                         "active" boolean DEFAULT true,
                         "created_at" timestamp NOT NULL DEFAULT 'now()',
                         "modified" timestamp NOT NULL DEFAULT 'now()',
                         "deleted" timestamp
);

CREATE INDEX ON "order" ("first_name", "last_name");

CREATE INDEX ON "order" ("email");

CREATE INDEX ON "order" ("phone");

CREATE INDEX ON "order_product" ("id_order");

CREATE INDEX ON "order_product" ("id_product");

CREATE INDEX ON "product" ("name");

CREATE INDEX ON "users" ("username");

ALTER TABLE "order_product" ADD FOREIGN KEY ("id_order") REFERENCES "order" ("id_order");

ALTER TABLE "order_product" ADD FOREIGN KEY ("id_product") REFERENCES "product" ("id_product");
