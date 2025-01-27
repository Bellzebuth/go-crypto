import React from "react";
import ListCryptos from "../components/List";
import Total from "../components/Total";

const Homepage: React.FC = () => {
  return (
    <div className="text-[#0b004b] p-5 flex items-center justify-center flex-col">
      <div className="bg-white rounded-md text-5xl text-center m-1 font-bold flex-row">
        Porto-crypto
      </div>
      <div className="flex w-3/4">
        <ListCryptos />
        <div className="flex flex-col">
          <Total />
        </div>
      </div>
    </div>
  );
};

export default Homepage;
