interface User {
  username: {
    value: string;
    setter: React.Dispatch<React.SetStateAction<string>>;
  };
  password: {
    value: string;
    setter: React.Dispatch<React.SetStateAction<string>>;
  };
  moderatedSubforums: {
    value: string;
    setter: React.Dispatch<React.SetStateAction<string>>;
  };
  posts: {
    value: string;
    setter: React.Dispatch<React.SetStateAction<string>>;
  };
  comments: {
    value: string;
    setter: React.Dispatch<React.SetStateAction<string>>;
  };
  votes: {
    value: string;
    setter: React.Dispatch<React.SetStateAction<string>>;
  };
  createdAt: {
    value: string;
    setter: React.Dispatch<React.SetStateAction<string>>;
  };
}