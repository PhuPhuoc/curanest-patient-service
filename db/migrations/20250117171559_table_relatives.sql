-- +goose Up
-- +goose StatementBegin
CREATE TABLE `relatives` (
    `id` varchar(36) NOT NULL,
    `dob` varchar(12) DEFAULT NULL,
    `address` varchar(300) NOT NULL,
    `ward` varchar(70) DEFAULT NULL,
    `district` varchar(70) DEFAULT NULL,
    `city` varchar(70) DEFAULT NULL,
    PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM `relatives`;
DROP TABLE `relatives`;
-- +goose StatementEnd
