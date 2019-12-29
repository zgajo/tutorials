import {MigrationInterface, QueryRunner} from "typeorm";

export class userSessionsCreatedAt1577654320081 implements MigrationInterface {
    name = 'userSessionsCreatedAt1577654320081'

    public async up(queryRunner: QueryRunner): Promise<any> {
        await queryRunner.query("ALTER TABLE `user_sessions` CHANGE `updatedAt` `createdAt` datetime NOT NULL", undefined);
    }

    public async down(queryRunner: QueryRunner): Promise<any> {
        await queryRunner.query("ALTER TABLE `user_sessions` CHANGE `createdAt` `updatedAt` datetime NOT NULL", undefined);
    }

}
