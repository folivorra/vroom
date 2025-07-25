-- +goose Up
-- +goose StatementBegin

ALTER TABLE "reservations"
    ALTER COLUMN "end_at" DROP NOT NULL,
ALTER COLUMN "total_price" DROP NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "reservations"
    ALTER COLUMN "end_at" SET NOT NULL,
ALTER COLUMN "total_price" SET NOT NULL;

-- +goose StatementEnd