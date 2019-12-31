import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  OneToOne,
  JoinColumn
} from "typeorm";
import { User } from "./users";

@Entity()
export class UserSessions {
  @PrimaryGeneratedColumn("uuid")
  id: string;

  @OneToOne(_ => User)
  @JoinColumn()
  user: User;

  @Column({ type: "datetime" })
  expiresAt: string;

  @Column({ type: "datetime" })
  createdAt: string;
}
