CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "first_last_name" varchar(200),
  "email" varchar(100) UNIQUE NOT NULL,
  "avatar_user" text NOT NULL,
  "password" varchar(255) NOT NULL,
  "status" bool NOT NULL DEFAULT true,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "roles" (
  "id" bigserial PRIMARY KEY,
  "role_name" char(50) NOT NULL DEFAULT 'regular',
  "status" bool NOT NULL DEFAULT true,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "users_roles" (
  "id" bigserial PRIMARY KEY,
  "fk_user" bigint,
  "fk_role" bigint
);

CREATE TABLE "menu_options" (
  "id" bigserial PRIMARY KEY,
  "label_url" char(20) NOT NULL,
  "name_icon" char(20) NOT NULL,
  "component_or_page_url" varchar(200) NOT NULL,
  "address_url" varchar(200) NOT NULL,
  "order_url" bigint NOT NULL DEFAULT 0,
  "tipo_url" char(1) NOT NULL DEFAULT 'L',
  "status" bool NOT NULL DEFAULT true,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "users_roles_menu_option" (
  "id" bigserial PRIMARY KEY,
  "fk_user" bigint,
  "fk_menu_option" bigint,
  "nivel" bigint NOT NULL DEFAULT 0
);

CREATE TABLE "companies" (
  "id" bigserial PRIMARY KEY,
  "name_company" varchar(150) NOT NULL,
  "addres_company" text NOT NULL,
  "phone_company" varchar(100) NOT NULL,
  "liable_company" varchar(100) NOT NULL,
  "banner_company" text NOT NULL,
  "status" bool NOT NULL DEFAULT true,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  "deleted_at" timestamp(6)
);

CREATE TABLE "institutions" (
  "id" bigserial PRIMARY KEY,
  "fk_companny" bigint,
  "name_intitution" varchar(150) NOT NULL,
  "addres_intitution" text NOT NULL,
  "phone_intitution" varchar(100) NOT NULL,
  "liable_intitution" varchar(100) NOT NULL,
  "logo_intitution" text NOT NULL,
  "status" bool NOT NULL DEFAULT true,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  "deleted_at" timestamp(6)
);

CREATE TABLE "users_intitutions" (
  "id" bigserial PRIMARY KEY,
  "fk_user" bigint,
  "fk_intitution" bigint
);

CREATE TABLE "municipalities" (
  "id" bigserial PRIMARY KEY,
  "code_municipality" varchar(2) UNIQUE NOT NULL,
  "name_municipality" varchar(100) UNIQUE NOT NULL
);

CREATE TABLE "parishes" (
  "id" bigserial PRIMARY KEY,
  "fk_municipality" bigint,
  "code_parish" varchar(3) UNIQUE NOT NULL,
  "name_parish" varchar(100) UNIQUE NOT NULL
);

CREATE TABLE "sectors" (
  "id" bigserial PRIMARY KEY,
  "fk_parish" bigint,
  "name_sector" varchar(250) NOT NULL
);

CREATE TABLE "taxpayers" (
  "id" bigserial PRIMARY KEY,
  "name_taxpayer" varchar(250) UNIQUE NOT NULL,
  "dni_taxpayer" varchar(15) UNIQUE NOT NULL,
  "address_taxpayer" text NOT NULL,
  "email_taxpayer" varchar(100),
  "latitud" float8,
  "longitud" float8,
  "liable_taxpayer" varchar(150),
  "email_liable_taxpayer" varchar(100),
  "status" bool NOT NULL DEFAULT true,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  "deleted_at" timestamp(6),
  "fk_municipality" bigint,
  "fk_parish" bigint,
  "fk_sector" bigint,
  "fk_user" bigint
);

CREATE TABLE "institutions_taxpayers" (
  "id" bigserial PRIMARY KEY,
  "number_file" varchar(20) UNIQUE NOT NULL,
  "fk_intitution" bigint,
  "fk_taxpayer" bigint
);

CREATE TABLE "payment_frequencies" (
  "id" bigserial PRIMARY KEY,
  "name_frequency" varchar(150) UNIQUE NOT NULL,
  "value_frequency" int8 NOT NULL DEFAULT 1
);

CREATE TABLE "payment_concepts" (
  "id" bigserial PRIMARY KEY,
  "name_payment_concept" varchar(150) UNIQUE NOT NULL,
  "formula" text NOT NULL,
  "fk_payment_frequency" bigint
);

CREATE TABLE "institutions_taxpayers_conceptos" (
  "id" bigserial PRIMARY KEY,
  "fk_institutions_taxpayers" bigint,
  "fk_payment_concept" bigint
);

CREATE TABLE "payments" (
  "id" bigserial PRIMARY KEY,
  "payment_date" date NOT NULL,
  "amount_payment" float8 NOT NULL DEFAULT 0,
  "fk_institutions_taxpayers" bigint,
  "observation" text,
  "status" bool NOT NULL DEFAULT true,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  "deleted_at" timestamp(6)
);

CREATE TABLE "payments_details" (
  "id" bigserial PRIMARY KEY,
  "fk_payment" bigint,
  "fk_institutions_taxpayers_concept" bigint,
  "amount_details" float8 NOT NULL DEFAULT 0
);

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "roles" ("role_name");

CREATE INDEX ON "users_roles" ("fk_user");

CREATE INDEX ON "users_roles" ("fk_role");

CREATE INDEX ON "menu_options" ("component_or_page_url");

CREATE INDEX ON "menu_options" ("address_url");

CREATE INDEX ON "users_roles_menu_option" ("fk_user");

CREATE INDEX ON "users_roles_menu_option" ("fk_menu_option");

CREATE INDEX ON "companies" ("name_company");

CREATE INDEX ON "municipalities" ("code_municipality");

CREATE INDEX ON "municipalities" ("name_municipality");

CREATE INDEX ON "parishes" ("code_parish");

CREATE INDEX ON "parishes" ("name_parish");

CREATE INDEX ON "taxpayers" ("name_taxpayer");

CREATE INDEX ON "taxpayers" ("dni_taxpayer");

CREATE INDEX ON "institutions_taxpayers" ("number_file");

CREATE INDEX ON "payment_frequencies" ("name_frequency");

CREATE INDEX ON "payment_concepts" ("name_payment_concept");

COMMENT ON COLUMN "companies"."liable_company" IS 'responsable de la empresa';

COMMENT ON COLUMN "institutions"."liable_intitution" IS 'responsable de la empresa';

COMMENT ON COLUMN "taxpayers"."liable_taxpayer" IS 'Responsable';

COMMENT ON COLUMN "taxpayers"."email_liable_taxpayer" IS 'Email del responsable';

ALTER TABLE "users_roles" ADD FOREIGN KEY ("fk_user") REFERENCES "users" ("id");

ALTER TABLE "users_roles" ADD FOREIGN KEY ("fk_role") REFERENCES "roles" ("id");

ALTER TABLE "users_roles_menu_option" ADD FOREIGN KEY ("fk_user") REFERENCES "users" ("id");

ALTER TABLE "users_roles_menu_option" ADD FOREIGN KEY ("fk_menu_option") REFERENCES "menu_options" ("id");

ALTER TABLE "institutions" ADD FOREIGN KEY ("fk_companny") REFERENCES "companies" ("id");

ALTER TABLE "users_intitutions" ADD FOREIGN KEY ("fk_user") REFERENCES "users" ("id");

ALTER TABLE "users_intitutions" ADD FOREIGN KEY ("fk_intitution") REFERENCES "institutions" ("id");

ALTER TABLE "parishes" ADD FOREIGN KEY ("fk_municipality") REFERENCES "municipalities" ("id");

ALTER TABLE "sectors" ADD FOREIGN KEY ("fk_parish") REFERENCES "parishes" ("id");

ALTER TABLE "taxpayers" ADD FOREIGN KEY ("fk_municipality") REFERENCES "municipalities" ("id");

ALTER TABLE "taxpayers" ADD FOREIGN KEY ("fk_parish") REFERENCES "parishes" ("id");

ALTER TABLE "taxpayers" ADD FOREIGN KEY ("fk_sector") REFERENCES "sectors" ("id");

ALTER TABLE "taxpayers" ADD FOREIGN KEY ("fk_user") REFERENCES "users" ("id");

ALTER TABLE "institutions_taxpayers" ADD FOREIGN KEY ("fk_intitution") REFERENCES "institutions" ("id");

ALTER TABLE "institutions_taxpayers" ADD FOREIGN KEY ("fk_taxpayer") REFERENCES "taxpayers" ("id");

ALTER TABLE "payment_concepts" ADD FOREIGN KEY ("fk_payment_frequency") REFERENCES "payment_frequencies" ("id");

ALTER TABLE "institutions_taxpayers_conceptos" ADD FOREIGN KEY ("fk_institutions_taxpayers") REFERENCES "institutions_taxpayers" ("id");

ALTER TABLE "institutions_taxpayers_conceptos" ADD FOREIGN KEY ("fk_payment_concept") REFERENCES "payment_concepts" ("id");

ALTER TABLE "payments" ADD FOREIGN KEY ("fk_institutions_taxpayers") REFERENCES "institutions_taxpayers" ("id");

ALTER TABLE "payments_details" ADD FOREIGN KEY ("fk_payment") REFERENCES "payments" ("id");

ALTER TABLE "payments_details" ADD FOREIGN KEY ("fk_institutions_taxpayers_concept") REFERENCES "institutions_taxpayers_conceptos" ("id");
