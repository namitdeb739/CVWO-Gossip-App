interface Comment {
  UserID: number;
  PostID: number;
  ParentCommentID: number | null;
  ChildrenComments: Comment[];
  CreatedAt: Date;
}
