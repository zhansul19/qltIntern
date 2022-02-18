CREATE TABLE IF NOT EXISTS  categories(
    id              serial       unique not null ,
    name            varchar(255) not null,
    price           int          not null    
);

CREATE TABLE IF NOT EXISTS  payments
(
    id          serial      not null unique,
    type        varchar(255)not null,
    date        timestamptz NOT NULL DEFAULT (now())
);
CREATE TABLE  IF NOT EXISTS  categorie_payments(
    id serial                                           unique not null ,
    category_id  int references categories (id)         on delete cascade not null,
    payment_id   int references payments(id)            on delete cascade not null
);