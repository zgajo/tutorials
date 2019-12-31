import {MigrationInterface, QueryRunner} from "typeorm";

export class changeUserSession1577804396602 implements MigrationInterface {
    name = 'changeUserSession1577804396602'

    public async up(queryRunner: QueryRunner): Promise<any> {
        await queryRunner.query("DROP INDEX `IDX_55fa4db8406ed66bc704432842` ON `user_sessions`", undefined);
        await queryRunner.query("ALTER TABLE `user_sessions` CHANGE `expiresAt` `expiresAt` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6)", undefined);
        await queryRunner.query("ALTER TABLE `user_sessions` CHANGE `createdAt` `createdAt` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6)", undefined);
    }

    public async down(queryRunner: QueryRunner): Promise<any> {
        await queryRunner.query("ALTER TABLE `user_sessions` CHANGE `createdAt` `createdAt` datetime(0) NOT NULL", undefined);
        await queryRunner.query("ALTER TABLE `user_sessions` CHANGE `expiresAt` `expiresAt` datetime(0) NOT NULL", undefined);
        await queryRunner.query("CREATE UNIQUE INDEX `IDX_55fa4db8406ed66bc704432842` ON `user_sessions` (`userId`)", undefined);
    }

}
