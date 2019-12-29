import { Entity, PrimaryGeneratedColumn, Column, ManyToOne } from "typeorm";
import { User } from "./users";

@Entity()
export class UserSessions {
  @PrimaryGeneratedColumn("uuid")
  id: string;

  @ManyToOne(
    _ => User,
    user => user.id,
    {
      nullable: false
    }
  )
  user: User;

  @Column({ type: "datetime" })
  expiresAt: string;

  @Column({ type: "datetime" })
  createdAt: string;
}
