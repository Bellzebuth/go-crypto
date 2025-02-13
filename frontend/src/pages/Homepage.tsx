import React from "react"
import ListSumCryptos from "../components/ListSum"
import { LogOut } from "lucide-react"
import { useAuth } from "../context/AuthContext"

const Home: React.FC = () => {
  const { logout } = useAuth()

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    logout()
  }

  return (
    <div className="bg-white w-3/4 h-full text-zinc-800 p-5">
      <div className="relative">
        <div className="text-zync-800 rounded-md text-5xl text-center m-1 font-bold flex-row">
          CryptoFolio
        </div>
        <button
          onClick={handleSubmit}
          className="absolute right-0 top-2 bg-zinc-800 text-white rounded-md my-2 p-1 disabled:bg-gray-300"
        >
          <LogOut />
        </button>
      </div>
      <ListSumCryptos />
    </div>
  )
}

export default Home
