import {MigrationInterface, QueryRunner} from "typeorm";

export class dateColumns1578340863332 implements MigrationInterface {
    name = 'dateColumns1578340863332'

    public async up(queryRunner: QueryRunner): Promise<any> {
        await queryRunner.query("ALTER TABLE `listing` CHANGE `createdAt` `createdAt` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6)", undefined);
        await queryRunner.query("ALTER TABLE `listing` CHANGE `updatedAt` `updatedAt` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6)", undefined);
    }

    public async down(queryRunner: QueryRunner): Promise<any> {
        await queryRunner.query("ALTER TABLE `listing` CHANGE `updatedAt` `updatedAt` datetime(0) NOT NULL", undefined);
        await queryRunner.query("ALTER TABLE `listing` CHANGE `createdAt` `createdAt` datetime(0) NOT NULL", undefined);
    }

}
