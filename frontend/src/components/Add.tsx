import React, { useState } from "react";
import api from "../services/api";
import Autocomplete from "./Autocomplete";

export type Crypto = {
  keyName: string;
  name: string;
};

interface AddCryptoProps {
  updateList: () => void;
}

const AddCrypto: React.FC<AddCryptoProps> = ({ updateList }) => {
  const [query, setQuery] = useState("");
  const [crypto, setCrypto] = useState<Crypto>({
    keyName: "",
    name: "",
  });
  const [amount, setAmount] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await api
        .post("/portfolio/add", {
          keyName: crypto.keyName,
          amount: parseFloat(amount),
        })
        .then(() => {
          setQuery("");
          setAmount("");
          setCrypto({
            keyName: "",
            name: "",
          });
          updateList();
        });
    } catch (error) {
      alert("Price is unavailable for this crypto…");
      console.error("Error adding crypto:", error);
    }
  };

  return (
    <div className="mb-2 border-b border-zinc-800">
      <div className="text-lg font-bold">New asset</div>
      <form onSubmit={handleSubmit} className="flex justify-between p-1">
        <Autocomplete query={query} setQuery={setQuery} setCrypto={setCrypto} />
        <input
          placeholder="Amount (€)"
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
          className="border border-gray-300 rounded-md focus:border-[#0b004b] focus:outline-none p-1 my-2"
          required
        />
        <button
          className="bg-zinc-800 text-white rounded-md my-2 p-1 disabled:bg-gray-300"
          type="submit"
          disabled={crypto.keyName === "" || amount == ""}
        >
          Add Crypto
        </button>
      </form>
    </div>
  );
};

export default AddCrypto;
