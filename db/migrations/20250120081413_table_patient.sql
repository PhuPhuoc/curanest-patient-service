-- +goose Up
-- +goose StatementBegin
CREATE TABLE `patients` (
    `id` varchar(36) NOT NULL,
    `relatives_id` varchar(36) NOT NULL,
    `full_name` varchar(50) NOT NULL,
    `gender` bool NOT NULL,
    `dob` varchar(12) DEFAULT NULL,
    `phone_number` varchar(12) NOT NULL,
    `address` varchar(300) NOT NULL,
    `ward` varchar(70) NOT NULL,
    `district` varchar(70) NOT NULL,
    `city` varchar(70) NOT NULL,
    `desc_pathology` longtext DEFAULT NULL,
    `note_for_nurse` longtext DEFAULT NULL,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` datetime,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_phone` (`phone_number`),
    CONSTRAINT `fk_relatives_patient` FOREIGN KEY (`relatives_id`) REFERENCES `relatives` (`id`) ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `patients` DROP FOREIGN KEY `fk_relatives_patient`;
DROP INDEX `unique_phone` ON `patients`;
DELETE FROM `patients`;
DROP TABLE `patients`;
-- +goose StatementEnd
