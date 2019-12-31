import {MigrationInterface, QueryRunner} from "typeorm";

export class userSessionOneToOne1577784907212 implements MigrationInterface {
    name = 'userSessionOneToOne1577784907212'

    public async up(queryRunner: QueryRunner): Promise<any> {
        await queryRunner.query("ALTER TABLE `user_sessions` ADD `userId` varchar(36) NULL", undefined);
        await queryRunner.query("ALTER TABLE `user_sessions` ADD UNIQUE INDEX `IDX_55fa4db8406ed66bc704432842` (`userId`)", undefined);
        await queryRunner.query("CREATE UNIQUE INDEX `REL_55fa4db8406ed66bc704432842` ON `user_sessions` (`userId`)", undefined);
        await queryRunner.query("ALTER TABLE `user_sessions` ADD CONSTRAINT `FK_55fa4db8406ed66bc7044328427` FOREIGN KEY (`userId`) REFERENCES `users`(`id`) ON DELETE NO ACTION ON UPDATE NO ACTION", undefined);
    }

    public async down(queryRunner: QueryRunner): Promise<any> {
        await queryRunner.query("ALTER TABLE `user_sessions` DROP FOREIGN KEY `FK_55fa4db8406ed66bc7044328427`", undefined);
        await queryRunner.query("DROP INDEX `REL_55fa4db8406ed66bc704432842` ON `user_sessions`", undefined);
        await queryRunner.query("ALTER TABLE `user_sessions` DROP INDEX `IDX_55fa4db8406ed66bc704432842`", undefined);
        await queryRunner.query("ALTER TABLE `user_sessions` DROP COLUMN `userId`", undefined);
    }

}
