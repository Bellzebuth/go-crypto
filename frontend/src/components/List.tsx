import React, { useEffect, useState } from "react";
import clsx from "clsx";
import api from "../services/api";
import AddCrypto from "./Add";
import { Trash2 } from "lucide-react";

type Crypto = {
  id: number;
  cryptoId: number;
  crypto: {
    id: number;
    keyName: string;
    name: string;
  };
  amount: number;
  created_at: string;
};

const ListCryptos: React.FC = () => {
  const [list, setList] = useState<Crypto[]>([]);
  const [needUpdate, setNeedUpdate] = useState<number>(0);

  const incrementNeedUpdate = () => {
    setNeedUpdate((prev) => prev + 1);
  };

  useEffect(() => {
    api
      .get("/portfolio/list")
      .then((response) => setList(response.data))
      .catch((error) => console.error("Error fetching portfolio:", error));
  }, [needUpdate]);

  const deleteAsset = (id: number) => {
    api
      .delete(`portfolio/${id}`)
      .then(() => incrementNeedUpdate())
      .catch((error) => console.error("Error :", error));
  };

  return (
    <div className="bg-white rounded-md w-full m-1 p-2">
      <AddCrypto updateList={incrementNeedUpdate} />
      <div
        className="grid border-t border-r border-l border-gray-300 bg-gray-100 rounded-t-md text-gray-600 p-1"
        style={{ gridTemplateColumns: "4fr 2fr 2fr 2rem" }}
      >
        <div className="">Asset</div>
        <div>Total investment (€)</div>
        <div>Last investment</div>
      </div>
      {!list || list.length === 0 ? (
        <div />
      ) : (
        list.map((item, index) => (
          <div
            key={index}
            className={clsx(
              "grid border-l border-r border-t border-gray-300 w-full hover:bg-gray-100 px-1",
              list.length - 1 === index && "border-b rounded-b-md"
            )}
            style={{ gridTemplateColumns: "4fr 2fr 2fr 2rem" }}
          >
            <div>{item.crypto.name}</div>
            <div>{item.amount}</div>
            <div>{new Date(item.created_at).toLocaleDateString()}</div>
            <div className="" onClick={() => deleteAsset(item.id)}>
              <Trash2 className="text-sm" />
            </div>
          </div>
        ))
      )}
    </div>
  );
};

export default ListCryptos;
