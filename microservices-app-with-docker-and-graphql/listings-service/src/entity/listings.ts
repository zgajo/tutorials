import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  BaseEntity,
  CreateDateColumn
} from "typeorm";

@Entity()
export class Listing extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ nullable: false })
  title: string;

  @Column({
    nullable: false,
    type: "text"
  })
  description: string;

  @CreateDateColumn({})
  createdAt: Date;

  @CreateDateColumn({})
  updatedAt: Date;

  @Column({ type: "datetime", nullable: true })
  deletedAt: string;
}
