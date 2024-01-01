interface User {
  Username: string;
  Password: string;
  ModeratedSubforums: Subforum[];
  Posts: Post[];
  Comments: Comment[];
  Votes: Vote[];
  CreatedAt: Date;
}