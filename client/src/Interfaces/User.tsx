interface User {
  ID: number;
  Username: string;
  Password: string;
  ModeratedSubforums: Subforum[];
  Posts: Post[];
  Comments: Comment[];
  Votes: Vote[];
  CreatedAt: Date;
}