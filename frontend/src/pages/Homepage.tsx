import React, { useState } from "react"
import { LogOut } from "lucide-react"
import { useAuth } from "../context/AuthContext"
import AddAddress from "../components/AddAddress"
import ListAddress from "../components/ListAddress"
import Dashboard from "./Dashboard"

const Home: React.FC = () => {
  const { logout } = useAuth()
  const [addressId, setAddressId] = useState<number>(0)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    logout()
  }

  return (
    <div className="bg-white w-full h-full text-zinc-800 flex">
      <div className="relative flex flex-col h-full w-64 p-2 gap-y-8">
        <div className="text-zync-800 rounded-md text-4xl text-center m-1 font-bold flex-row">
          CryptoFolio
        </div>
        <AddAddress />
        <ListAddress setAddressId={setAddressId} />
        <button
          onClick={handleSubmit}
          className="absolute bottom-2 w-full font-bold bg-zinc-800 text-white rounded-md my-2 p-1"
        >
          Log out
          <LogOut className="absolute right-2 top-1" />
        </button>
      </div>
      {addressId != 0 && <Dashboard addressId={addressId} />}
    </div>
  )
}

export default Home
