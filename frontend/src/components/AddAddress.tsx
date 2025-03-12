import React, { useState } from "react"
import api from "../services/api"
import Dropdown from "../UI/DropDown"
import { Plus } from "lucide-react"

const AddAddress: React.FC = () => {
  const [address, setAddress] = useState("")
  const [name, setName] = useState("")
  const [blockchain, setBlockchain] = useState<Blockchain>({ id: 0, name: "" })

  const resetBlockchain = () => setBlockchain({ id: 0, name: "" })

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    try {
      await api
        .post("/address/add", {
          Address: address,
          Name: name,
          BlockchainId: blockchain.id,
        })
        .then(() => {
          setAddress("")
          setName("")
          resetBlockchain()
        })
    } catch (error) {
      console.error("can't add address :", error)
    }
  }

  return (
    <div className="border-zinc-800">
      <div className="flex font-bold pb-1">
        <Plus />
        New wallet address
      </div>
      <form onSubmit={handleSubmit} className="flex flex-col justify-between">
        <Dropdown
          url="/blockchain/list"
          value={blockchain}
          setValue={setBlockchain}
        />
        <input
          type="text"
          placeholder="Wallet address…"
          value={address}
          onChange={e => setAddress(e.target.value)}
          className="h-8 border border-zinc-300 rounded-md focus:border-[#0b004b] focus:outline-none px-2 py-1 my-2"
          required
        />
        <input
          type="text"
          placeholder="Name…"
          value={name}
          onChange={e => setName(e.target.value)}
          className="h-8 border border-zinc-300 rounded-md focus:border-[#0b004b] focus:outline-none px-2 py-1 my-2"
          required
        />
        <button
          className="bg-zinc-800 text-white rounded-md my-2 p-1 disabled:bg-zinc-300"
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
