CREATE SCHEMA IF NOT EXISTS `daily_work` DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
USE `daily_work`;

CREATE TABLE IF NOT EXISTS `kakao`
(
    `id`                       INT           NOT NULL AUTO_INCREMENT,
    `kakao_id`                 BIGINT        NOT NULL DEFAULT '0',
    `nickname`                 VARCHAR(100)  NOT NULL,
    `profile_image`            VARCHAR(1000) NOT NULL,
    `mobile`                   VARCHAR(100)  NOT NULL COMMENT '휴대전화',
    `email`                    VARCHAR(100)  NOT NULL,
    `gender`                   VARCHAR(10)   NULL COMMENT '성별',
    `age_range`                VARCHAR(30)   NULL COMMENT '연령대',
    `created_at`               DATETIME      NOT NULL,
    `updated_at`               DATETIME      NOT NULL,
    `withdraw_at`              DATETIME      NULL COMMENT '탈퇴날짜',
    PRIMARY KEY (`id`),
    INDEX `idx_members_mobile` (`mobile` ASC)
);