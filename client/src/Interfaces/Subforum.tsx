interface Subforum {
  ID: number;
  Name: string;
  Description: string;
  Moderators: User[];
  Posts: Post[];
  CreatedAt: Date;
}
