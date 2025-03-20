import React, { useState } from "react"
import clsx from "clsx"
import { ChevronDown, ChevronUp } from "lucide-react"
import {
  formatToPercentage,
  formatToTwoDecimalsPrice,
} from "../services/format"

const Row: React.FC<{ item: TransactionSum }> = ({ item }) => {
  const [showDetails, setShowDetails] = useState<boolean>(false)

  return (
    <div>
      <div
        className={clsx(
          "grid border border-gray-300 w-full hover:bg-gray-100 px-1"
        )}
        style={{ gridTemplateColumns: "2fr 2fr 2fr 2fr 2fr 2rem" }}
      >
        <div className="tabular-nums">
          {formatToTwoDecimalsPrice(item.value)}
        </div>
        <div className="tabular-nums">
          {formatToTwoDecimalsPrice(item.price.price)}
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
      {/* {showDetails && <ListDetails keyName={item.keyName} />} */}
    </div>
  )
}

interface TransactionsProps {
  list: TransactionSum[]
}

const Transactions: React.FC<TransactionsProps> = ({ list }) => {
  return (
    <div className="bg-white rounded-md w-full m-1 p-2">
      {!list || list.length === 0 ? (
        <div />
      ) : (
        list.map((item, index) => (
          <div className="mt-4">
            <div className="text-lg font-bold">
              {item.address.blockchain.name}
            </div>
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
    </div>
  )
}

export default Transactions
