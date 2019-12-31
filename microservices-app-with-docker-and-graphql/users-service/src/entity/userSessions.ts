import {
  Entity,
  PrimaryGeneratedColumn,
  OneToOne,
  JoinColumn,
  BaseEntity,
  CreateDateColumn
} from "typeorm";
import { User } from "./users";

@Entity()
export class UserSessions extends BaseEntity {
  @PrimaryGeneratedColumn("uuid")
  id: string;

  @OneToOne(_ => User)
  @JoinColumn()
  user: User;

  @CreateDateColumn({})
  expiresAt: Date;

  @CreateDateColumn({})
  createdAt: Date;
}
