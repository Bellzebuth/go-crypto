import React, { useEffect, useState } from "react";
import api from "../services/api";
import { formatToTwoDecimalsPrice } from "../lib/format";

interface Totals {
  totalInvested: number;
  totalValue: number;
}

const Totals: React.FC = () => {
  const [totals, setTotals] = useState<Totals>({
    totalInvested: 0,
    totalValue: 0,
  });

  useEffect(() => {
    api
      .get("/portfolio/total")
      .then((response) => setTotals(response.data))
      .catch((error) => console.error("Error fetching portfolio:", error));
  }, []);

  return (
    <div className="bg-white w-full flex justify-between rounded-md mt-4 p-1">
      <div className="">
        <div className="text-lg font-bold">Invested</div>
        <div className="text-3xl text-end">
          {formatToTwoDecimalsPrice(totals.totalInvested)}
        </div>
      </div>
      <div>
        <div className="text-lg font-bold text-end">Balance</div>
        <div className="text-3xl text-end">
          {formatToTwoDecimalsPrice(totals.totalValue)}
        </div>
      </div>
    </div>
  );
};

export default Totals;
