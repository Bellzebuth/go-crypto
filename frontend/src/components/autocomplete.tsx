import React, { useState } from "react";
import axios from "axios";
import { Crypto } from "./Add";

type AutocompleteProps = {
  setCrypto: (crypto: Crypto) => void;
};

const Autocomplete: React.FC<AutocompleteProps> = ({ setCrypto }) => {
  const [query, setQuery] = useState("");
  const [results, setResults] = useState<Crypto[]>([]);
  const [isLoading, setIsLoading] = useState(false);

  const handleSearch = async (searchQuery: string) => {
    setQuery(searchQuery);

    if (searchQuery.trim() === "") {
      setResults([]);
      return;
    }

    setIsLoading(true);
    try {
      const response = await axios.get<Crypto[]>(
        `http://localhost:8080/cryptos/list?query=${encodeURIComponent(
          searchQuery
        )}`
      );
      setResults(response.data);
    } catch (error) {
      console.error("Error fetching cryptos:", error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="relative">
      <input
        type="text"
        placeholder="cryptocurrency..."
        value={query}
        onChange={(e) => handleSearch(e.target.value)}
        className="border border-gray-300 rounded-md focus:border-[#0b004b] focus:outline-none p-1 my-2"
        required
      />
      {results && results.length > 0 && (
        <div className="absolute w-full bg-white border rounded-md border-gray-400 w-full">
          <ul>
            {results.map((crypto) => (
              <li
                key={crypto.id}
                className="p-2 cursor-pointer hover:bg-gray-100"
                onClick={() => {
                  setCrypto(crypto);
                  setQuery(crypto.name);
                  setResults([]);
                }}
              >
                {crypto.name}
              </li>
            ))}
          </ul>
        </div>
      )}
      {isLoading && <p className="absolute text-sm bg-gray-050">Loading...</p>}
    </div>
  );
};

export default Autocomplete;
