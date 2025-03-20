import React, { useEffect, useState } from "react"
import api from "../services/api"
import Transactions from "../components/Transactions"
import TradingViewWidget from "../components/charts/EthereumCharts"

interface WalletDashboardProps {
  addressId: number
}

const Dashboard: React.FC<WalletDashboardProps> = ({ addressId }) => {
  const [list, setList] = useState<TransactionSum[]>([])

  useEffect(() => {
    api
      .get(`/transactions/list?addressId=${addressId}`)
      .then(response => setList(response.data))
      .catch(error => console.error("Error fetching portfolio:", error))
  }, [addressId])

  return (
    <div className="w-full p-2">
      <div className="flex">
        <TradingViewWidget />
      </div>
      <Transactions list={list} />
    </div>
  )
}

export default Dashboard
