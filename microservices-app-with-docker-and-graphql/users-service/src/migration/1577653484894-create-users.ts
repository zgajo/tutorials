import {MigrationInterface, QueryRunner} from "typeorm";

export class createUsers1577653484894 implements MigrationInterface {
    name = 'createUsers1577653484894'

    public async up(queryRunner: QueryRunner): Promise<any> {
        await queryRunner.query("CREATE TABLE `users` (`id` varchar(36) NOT NULL, `email` varchar(255) NOT NULL, `passwordHash` char(64) NOT NULL, UNIQUE INDEX `IDX_97672ac88f789774dd47f7c8be` (`email`), PRIMARY KEY (`id`)) ENGINE=InnoDB", undefined);
    }

    public async down(queryRunner: QueryRunner): Promise<any> {
        await queryRunner.query("DROP INDEX `IDX_97672ac88f789774dd47f7c8be` ON `users`", undefined);
        await queryRunner.query("DROP TABLE `users`", undefined);
    }

}
