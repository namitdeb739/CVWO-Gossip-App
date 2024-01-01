interface Subforum {
  Name: string;
  Description: string;
  Moderators: User[];
  Posts: Post[];
  CreatedAt: Date;
}
