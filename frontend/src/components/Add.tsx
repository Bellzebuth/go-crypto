import React, { useState } from "react";
import api from "../services/api";
import Autocomplete from "./autocomplete";

export type Crypto = {
  id: number;
  keyName: string;
  name: string;
};

interface AddCryptoProps {
  updateList: () => void;
}

const AddCrypto: React.FC<AddCryptoProps> = ({ updateList }) => {
  const [crypto, setCrypto] = useState<Crypto>({
    id: 0,
    keyName: "",
    name: "",
  });
  const [amount, setAmount] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await api
        .post("/portfolio/add", {
          cryptoId: crypto.id,
          amount: parseFloat(amount),
        })
        .then(() => {
          setAmount("");
          setCrypto({
            id: 0,
            keyName: "",
            name: "",
          });
          updateList();
        });
    } catch (error) {
      console.error("Error adding crypto:", error);
    }
  };

  return (
    <div className="bg-white rounded-md border border-gray-300 mb-2">
      <div className="bg-[#0b004b] text-white font-bold rounded-t-md px-1">
        Add crypto to your wallet
      </div>
      <form onSubmit={handleSubmit} className="flex justify-between p-1">
        <Autocomplete setCrypto={setCrypto} />
        <input
          placeholder="Amount (€)"
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
          className="border border-gray-300 rounded-md focus:border-[#0b004b] focus:outline-none p-1 my-2"
          required
        />
        <button
          className="bg-[#0b004b] text-white rounded-md my-2 p-1"
          type="submit"
        >
          Add Crypto
        </button>
      </form>
    </div>
  );
};

export default AddCrypto;
