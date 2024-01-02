import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Profile from "./Profile";
import { ENDPOINT } from "../App";

function UserPage() {
  const { ID } = useParams();
  const [user, setUser] = useState<User>({
    Username: "",
    Password: "",
    ModeratedSubforums: [],
    Posts: [],
    Comments: [],
    Votes: [],
    CreatedAt: new Date(),
  });

  useEffect(() => {
    (async () => {
      const response = await fetch(ENDPOINT + "/api/user/" + ID, {
        headers: { "Content-Type": "application/json" },
      });

      const content = await response.json();

      setUser({
        ...content.data,
        ModeratedSubforums: content.data.Moderated_Subforums,
      });
    })();
  }, [ID]);

  return <Profile user={user} />;
}

export default UserPage;
