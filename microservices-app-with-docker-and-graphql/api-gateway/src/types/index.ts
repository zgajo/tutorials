export interface USER_INSERT {
  email: string;
  password: string;
}
export interface USER {
  id: string;
  email: string;
}

export interface USER_SESSION_FETCHED {
  userId: string;
  user: USER;
}
