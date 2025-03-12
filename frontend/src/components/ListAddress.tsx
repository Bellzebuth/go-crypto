import React, { useEffect, useState, useCallback } from "react"
import clsx from "clsx"
import api from "../services/api"
import { ChevronRight, Wallet } from "lucide-react"
import BlockchainNames from "../types/blockchainNames"

const BlockchainIcons = ({ name, size }: { name: string; size: number }) => {
  if (name === BlockchainNames.Bitcoin) {
    return <img src="/bitcoin.png" alt="Bitcoin" width={size} height={size} />
  } else if (name === BlockchainNames.Ethereum) {
    return <img src="/ethereum.png" alt="Bitcoin" width={size} height={size} />
  } else {
    return <img src="/bnb.webp" alt="Bitcoin" width={size} height={size} />
  }
}

interface ListAddressProps {
  setAddressId: (id: number) => void
}

const ListAddress: React.FC<ListAddressProps> = ({
  setAddressId: setWalletId,
}) => {
  const [list, setList] = useState<Address[]>([])

  const refresh = useCallback(() => {
    api
      .get<Address[]>(`/address/list`)
      .then(response => setList(response.data))
      .catch(error => console.error("Error fetching cryptos:", error))
  }, [])

  useEffect(() => {
    refresh()
  }, [refresh])

  //   const deleteAddress = (id: number) => {
  //     api
  //       .delete(`/address/${id}`)
  //       .then(() => refresh())
  //       .catch(error => console.error("Error :", error))
  //   }

  return (
    <div className="w-full">
      <div className="flex font-bold gap-x-2 mb-2">
        <Wallet />
        My wallets
      </div>
      {!list || list.length === 0 ? (
        <div />
      ) : (
        list.map((item, index) => (
          <div
            key={index}
            className={clsx(
              "relative text-gray-600 border-b border-zinc-800 w-full hover:bg-gray-100 px-1 py-2"
            )}
            onClick={() => setWalletId(item.id)}
          >
            <div className="flex gap-x-2 w-full truncate p-1 pr-4">
              <BlockchainIcons name={item.blockchain.name} size={24} />
              {item.name ? item.name : item.address}
            </div>
            <ChevronRight className="absolute right-2 top-3" />
          </div>
        ))
      )}
    </div>
  )
}

export default ListAddress
