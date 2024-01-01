interface Post {
  UserID: number;
  SubforumID: number;
  Title: string;
  Body: string;
  Comments: Comment[];
  Votes: Vote[];
  Tags: Tag[];
  CreatedAt: Date;
}
