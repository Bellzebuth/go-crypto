import React, { useEffect, useState } from "react";
import clsx from "clsx";
import api from "../services/api";
import AddCrypto from "./Add";
import { ChevronDown, ChevronUp } from "lucide-react";
import ListDetails from "./ListDetails";
import {
  formatToPercentage,
  formatToTwoDecimalsPrice,
} from "../services/format";
import Totals from "./Totals";

type CryptoSum = {
  keyName: string;
  crypto: {
    keyName: string;
    name: string;
  };
  amount: number;
  gain: number;
  percentageGain: number;
  actualPrice: number;
  actualValue: number;
};

const Row: React.FC<{ item: CryptoSum }> = ({ item }) => {
  const [showDetails, setShowDetails] = useState<boolean>(false);

  return (
    <div>
      <div
        className={clsx(
          "grid border border-gray-300 w-full hover:bg-gray-100 px-1"
        )}
        style={{ gridTemplateColumns: "2fr 2fr 2fr 2fr 2fr 2rem" }}
      >
        <div className="tabular-nums">
          {formatToTwoDecimalsPrice(item.amount)}
        </div>
        <div className="tabular-nums">
          {formatToTwoDecimalsPrice(item.actualPrice)}
        </div>
        <div className="tabular-nums">
          {formatToTwoDecimalsPrice(item.gain)}
        </div>
        {formatToPercentage(item.percentageGain)}
        <div className="tabular-nums">
          {formatToTwoDecimalsPrice(item.actualValue)}
        </div>
        <div>
          {showDetails ? (
            <ChevronUp onClick={() => setShowDetails(!showDetails)} />
          ) : (
            <ChevronDown onClick={() => setShowDetails(!showDetails)} />
          )}
        </div>
      </div>
      {showDetails && <ListDetails keyName={item.keyName} />}
    </div>
  );
};

const ListSumCryptos: React.FC = () => {
  const [list, setList] = useState<CryptoSum[]>([]);
  const [needUpdate, setNeedUpdate] = useState<number>(0);

  const incrementNeedUpdate = () => {
    setNeedUpdate((prev) => prev + 1);
  };

  useEffect(() => {
    api
      .get("/portfolio/listsum")
      .then((response) => setList(response.data))
      .catch((error) => console.error("Error fetching portfolio:", error));
  }, [needUpdate]);

  return (
    <div className="bg-white rounded-md w-full m-1 p-2">
      <AddCrypto updateList={incrementNeedUpdate} />
      {!list || list.length === 0 ? (
        <div />
      ) : (
        list.map((item, index) => (
          <div className="mt-4">
            <div className="text-lg font-bold">{item.crypto.name}</div>
            <div
              className="grid border-t border-r border-l border-gray-300 bg-gray-100 rounded-t-md text-gray-600 p-1"
              style={{ gridTemplateColumns: "2fr 2fr 2fr 2fr 2fr 2rem" }}
            >
              <div>Invested</div>
              <div>Actual price</div>
              <div>Gain</div>
              <div>Percentage</div>
              <div>Value</div>
            </div>
            <Row key={index} item={item} />
          </div>
        ))
      )}
      <Totals />
    </div>
  );
};

export default ListSumCryptos;
