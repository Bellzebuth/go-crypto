import React from "react";
import ListSumCryptos from "../components/ListSum";

const Home: React.FC = () => {
  return (
    <div className="bg-white w-3/4 h-full text-zinc-800 p-5">
      <div className="text-zync-800 rounded-md text-5xl text-center m-1 font-bold flex-row">
        CryptoFolio
      </div>
      <ListSumCryptos />
    </div>
  );
};

export default Home;
