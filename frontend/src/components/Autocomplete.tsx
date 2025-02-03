import React, { useState, useEffect } from "react";
import api from "../services/api";
import { Crypto } from "./Add";
import clsx from "clsx";
import { X } from "lucide-react";

type AutocompleteProps = {
  query: string;
  setQuery: (query: string) => void;
  setCrypto: (crypto: Crypto) => void;
};

const Autocomplete: React.FC<AutocompleteProps> = ({
  query,
  setQuery,
  setCrypto,
}) => {
  const [results, setResults] = useState<Crypto[]>([]);
  const [indexResults, setIndexResults] = useState(0);
  const [showResults, setShowResults] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.key === "ArrowDown") {
        setIndexResults((prev) => Math.min(prev + 1, results.length - 1));
      } else if (event.key === "ArrowUp") {
        setIndexResults((prev) => Math.max(prev - 1, 0));
      } else if (event.key === "Enter") {
        setCrypto(results[indexResults]);
        setQuery(results[indexResults].name);
        setResults([]);
      }
    };

    window.addEventListener("keydown", handleKeyDown);
    return () => {
      window.removeEventListener("keydown", handleKeyDown);
    };
  }, [indexResults, results, setCrypto, setQuery]);

  const handleSearch = async (searchQuery: string) => {
    setQuery(searchQuery);

    if (searchQuery.trim() === "") {
      setResults([]);
      return;
    }

    setIsLoading(true);

    api
      .get<Crypto[]>(`/cryptos/list?query=${encodeURIComponent(searchQuery)}`)
      .then((response) => setResults(response.data))
      .catch((error) => console.error("Error fetching cryptos:", error))
      .finally(() => setIsLoading(false));
  };

  return (
    <div
      className="relative"
      onFocus={() => setShowResults(true)}
      onBlur={() => setShowResults(false)}
    >
      <input
        type="text"
        placeholder="cryptocurrency…"
        value={query}
        onChange={(e) => handleSearch(e.target.value)}
        className="border border-zinc-300 rounded-md focus:border-[#0b004b] focus:outline-none p-1 my-2"
        required
      />
      <X
        className="absolute text-zinc-500 right-0 top-0 mt-3 mr-1"
        onClick={() => {
          setCrypto({ keyName: "", name: "" });
          setQuery("");
          setResults([]);
        }}
      />
      {showResults && results && results.length > 0 && (
        <div className="absolute w-full bg-white border rounded-md border-zinc-400 w-full">
          <ul>
            {results.map((crypto, index) => (
              <li
                key={crypto.keyName}
                className={clsx(
                  "p-2 cursor-pointer",
                  indexResults === index && "bg-blue-600 text-white"
                )}
                onClick={() => {
                  setCrypto(crypto);
                  setQuery(crypto.name);
                  setResults([]);
                }}
                onMouseOver={() => setIndexResults(index)}
              >
                {crypto.name}
              </li>
            ))}
          </ul>
        </div>
      )}
      {isLoading && <p className="absolute text-sm bg-zinc-050">Loading…</p>}
    </div>
  );
};

export default Autocomplete;
