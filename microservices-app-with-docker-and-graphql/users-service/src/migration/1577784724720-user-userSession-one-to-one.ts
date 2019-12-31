import {MigrationInterface, QueryRunner} from "typeorm";

export class userUserSessionOneToOne1577784724720 implements MigrationInterface {
    name = 'userUserSessionOneToOne1577784724720'

    public async up(queryRunner: QueryRunner): Promise<any> {
        await queryRunner.query("ALTER TABLE `user_sessions` DROP FOREIGN KEY `FK_55fa4db8406ed66bc7044328427`", undefined);
        await queryRunner.query("ALTER TABLE `user_sessions` DROP COLUMN `userId`", undefined);
    }

    public async down(queryRunner: QueryRunner): Promise<any> {
        await queryRunner.query("ALTER TABLE `user_sessions` ADD `userId` varchar(36) NOT NULL", undefined);
        await queryRunner.query("ALTER TABLE `user_sessions` ADD CONSTRAINT `FK_55fa4db8406ed66bc7044328427` FOREIGN KEY (`userId`) REFERENCES `users`(`id`) ON DELETE NO ACTION ON UPDATE NO ACTION", undefined);
    }

}
