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
    value: any[];
    setter: React.Dispatch<React.SetStateAction<any[]>>;
  };
  posts: {
    value: any[];
    setter: React.Dispatch<React.SetStateAction<any[]>>;
  };
  comments: {
    value: any[];
    setter: React.Dispatch<React.SetStateAction<any[]>>;
  };
  votes: {
    value: any[];
    setter: React.Dispatch<React.SetStateAction<any[]>>;
  };
  createdAt: {
    value: number;
    setter: React.Dispatch<React.SetStateAction<number>>;
  };
}