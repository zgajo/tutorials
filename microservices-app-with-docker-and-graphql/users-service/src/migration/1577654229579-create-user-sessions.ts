import {MigrationInterface, QueryRunner} from "typeorm";

export class createUserSessions1577654229579 implements MigrationInterface {
    name = 'createUserSessions1577654229579'

    public async up(queryRunner: QueryRunner): Promise<any> {
        await queryRunner.query("CREATE TABLE `user_sessions` (`id` varchar(36) NOT NULL, `expiresAt` datetime NOT NULL, `updatedAt` datetime NOT NULL, `userId` varchar(36) NOT NULL, PRIMARY KEY (`id`)) ENGINE=InnoDB", undefined);
        await queryRunner.query("ALTER TABLE `user_sessions` ADD CONSTRAINT `FK_55fa4db8406ed66bc7044328427` FOREIGN KEY (`userId`) REFERENCES `users`(`id`) ON DELETE NO ACTION ON UPDATE NO ACTION", undefined);
    }

    public async down(queryRunner: QueryRunner): Promise<any> {
        await queryRunner.query("ALTER TABLE `user_sessions` DROP FOREIGN KEY `FK_55fa4db8406ed66bc7044328427`", undefined);
        await queryRunner.query("DROP TABLE `user_sessions`", undefined);
    }

}
