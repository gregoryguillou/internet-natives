import { useEffect, useState } from "react";

//
type SafeOptions = {
  owner: string;
};

type SafesResponse = { safes: [string] };

const Safe = ({ owner }: SafeOptions) => {
  const [safes, setSafes] = useState([] as [string] | []);
  const [loading, setLoading] = useState(false);

  const get = async (owner: string) => {
    if (!owner) {
      setSafes([]);
      setLoading(false);
      return;
    }
    const response = await fetch(
      `https://safe-client.safe.global/v1/chains/5/owners/${owner}/safes`,
      {
        method: "GET",
        mode: "cors",
        cache: "no-cache",
        headers: {
          "Content-Type": "application/json",
        },
        redirect: "follow",
        referrerPolicy: "no-referrer",
      }
    );
    const data: SafesResponse = await response.json();
    setSafes(data.safes);
    setLoading(false);
  };

  useEffect(() => {
    setLoading(true);
    get(owner);
  }, [owner]);

  return (
    <div>
      {loading ? (
        "loading"
      ) : safes && safes.length > 0 ? (
        <ol>
          {safes.map((value: string) => (
            <li key={value}>{value}</li>
          ))}
        </ol>
      ) : (
        "no safe"
      )}
    </div>
  );
};

export default Safe;
