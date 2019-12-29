import {MigrationInterface, QueryRunner} from "typeorm";

export class createListings1577652314937 implements MigrationInterface {
    name = 'createListings1577652314937'

    public async up(queryRunner: QueryRunner): Promise<any> {
        await queryRunner.query("CREATE TABLE `listing` (`id` int NOT NULL AUTO_INCREMENT, `title` varchar(255) NOT NULL, `description` text NOT NULL, `createdAt` datetime NOT NULL, `updatedAt` datetime NOT NULL, `deletedAt` datetime NULL, PRIMARY KEY (`id`)) ENGINE=InnoDB", undefined);
    }

    public async down(queryRunner: QueryRunner): Promise<any> {
        await queryRunner.query("DROP TABLE `listing`", undefined);
    }

}
