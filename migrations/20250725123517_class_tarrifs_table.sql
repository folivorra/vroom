-- +goose Up
-- +goose StatementBegin

CREATE TABLE "class_tariffs" (
     "id" SERIAL PRIMARY KEY NOT NULL,
     "class_type" class_type UNIQUE NOT NULL,
     "price_per_minute" decimal(10, 2) NOT NULL,
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "class_tariffs";

-- +goose StatementEnd
