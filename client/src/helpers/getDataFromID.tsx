import { useEffect, useState } from "react";

function getDataFromID<T>(id: number, table: string): T | undefined {
  const [data, setData] = useState();

  useEffect(() => {
    (async () => {
      const response = await fetch(
        "http://localhost:8080/api/" + table + "/" + id,
        {
          headers: { "Content-Type": "application/json" },
        }
      );

      const content = await response.json();

      setData(content.data);
    })();
  }, [id, table]);

  return data;
}

export default getDataFromID;
