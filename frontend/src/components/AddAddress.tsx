import React, { useState } from "react"
import api from "../services/api"
import Dropdown from "../UI/DropDown"

const AddAddress: React.FC = () => {
  const [address, setAddress] = useState("")
  const [blockchain, setBlockchain] = useState<Blockchain>({ id: 0, name: "" })

  const resetBlockchain = () => setBlockchain({ id: 0, name: "" })

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    try {
      await api
        .post("/address/add", {
          Address: address,
          BlockchainId: blockchain.id,
        })
        .then(() => {
          setAddress("")
          resetBlockchain()
        })
    } catch (error) {
      console.error("can't add address :", error)
    }
  }

  return (
    <div className="mb-2 border-b border-zinc-800">
      <div className="text-lg font-bold">New wallet address</div>
      <form
        onSubmit={handleSubmit}
        className="flex flex-col justify-between p-1"
      >
        <Dropdown
          url="/blockchain/list"
          value={blockchain}
          setValue={setBlockchain}
        />
        <input
          type="text"
          placeholder="wallet address..."
          value={address}
          onChange={e => setAddress(e.target.value)}
          className="border border-gray-300 rounded-md focus:border-[#0b004b] focus:outline-none p-1 my-2"
          required
        />
        <button
          className="bg-zinc-800 text-white rounded-md my-2 p-1 disabled:bg-gray-300"
          type="submit"
          disabled={blockchain.id === 0 || address == ""}
        >
          Add
        </button>
      </form>
    </div>
  )
}

export default AddAddress
