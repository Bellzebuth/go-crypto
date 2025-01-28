import React, { useEffect, useState, useCallback } from "react";
import clsx from "clsx";
import api from "../services/api";
import { Trash2 } from "lucide-react";
import {
  formatToPercentage,
  formatToTwoDecimalsPrice,
} from "../services/format";

type Crypto = {
  id: number;
  keyName: string;
  crypto: {
    keyName: string;
    name: string;
  };
  amount: number;
  purchasedPrice: number;
  created_at: string;
  gain: number;
  percentageGain: number;
  actualPrice: number;
  actualValue: number;
};

interface ListDetailsProps {
  keyName: string;
}

const ListDetails: React.FC<ListDetailsProps> = ({ keyName }) => {
  const [list, setList] = useState<Crypto[]>([]);

  const refresh = useCallback(() => {
    api
      .get<Crypto[]>(`/portfolio/list?keyName=${encodeURIComponent(keyName)}`)
      .then((response) => setList(response.data))
      .catch((error) => console.error("Error fetching cryptos:", error));
  }, [keyName]);

  useEffect(() => {
    refresh();
  }, [refresh]);

  const deleteAsset = (id: number) => {
    api
      .delete(`portfolio/${id}`)
      .then(() => refresh())
      .catch((error) => console.error("Error :", error));
  };

  return (
    <div className="w-full">
      <div
        className="grid bg-gray-200 text-gray-600 border-l border-r border-t border-gray-300 w-full hover:bg-gray-100 px-1"
        style={{ gridTemplateColumns: "2fr 2fr 2fr 2fr 2fr 2rem" }}
      >
        <div>Invested</div>
        <div>Purchased price</div>
        <div>Gain</div>
        <div>Percentage</div>
        <div>Value</div>
      </div>
      {!list || list.length === 0 ? (
        <div />
      ) : (
        list.map((item, index) => (
          <div
            key={index}
            className={clsx(
              "grid text-gray-600 border-l border-r border-t border-gray-300 w-full hover:bg-gray-100 px-1",
              list.length - 1 === index && "border-b rounded-b-md"
            )}
            style={{ gridTemplateColumns: "2fr 2fr 2fr 2fr 2fr 2rem" }}
          >
            <div className="tabular-nums">
              {formatToTwoDecimalsPrice(item.amount)}
            </div>
            <div className="tabular-nums">
              {formatToTwoDecimalsPrice(item.purchasedPrice)}
            </div>
            <div className="tabular-nums">
              {formatToTwoDecimalsPrice(item.gain)}
            </div>
            {formatToPercentage(item.percentageGain)}
            <div className="tabular-nums">
              {formatToTwoDecimalsPrice(item.actualValue)}
            </div>
            <div className="" onClick={() => deleteAsset(item.id)}>
              <Trash2 className="text-sm" />
            </div>
          </div>
        ))
      )}
    </div>
  );
};

export default ListDetails;
