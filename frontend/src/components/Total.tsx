import React, { useEffect, useState } from "react";
import api from "../services/api";

const Total: React.FC = () => {
  const [total, setTotal] = useState<number>(0);

  useEffect(() => {
    api
      .get("/portfolio/total")
      .then((response) => setTotal(response.data))
      .catch((error) => console.error("Error fetching portfolio:", error));
  }, []);

  return (
    <div className="bg-white rounded-md m-1 p-1">
      <div className="text-lg text-center">Total investi</div>
      <div className="text-5xl text-center">{total}€</div>
    </div>
  );
};

export default Total;
