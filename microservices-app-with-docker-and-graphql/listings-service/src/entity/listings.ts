import { Entity, PrimaryGeneratedColumn, Column } from "typeorm";

@Entity()
export class Listing {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ nullable: false })
  title: string;

  @Column({
    nullable: false,
    type: "text"
  })
  description: string;

  @Column({ type: "datetime" })
  createdAt: string;

  @Column({ type: "datetime" })
  updatedAt: string;

  @Column({ type: "datetime", nullable: true })
  deletedAt: string;
}
