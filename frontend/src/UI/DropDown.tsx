import { useEffect, useState } from "react"
import api from "../services/api"
import { ChevronDown } from "lucide-react"

interface DropdownProps<T extends { name: string }> {
  url: string
  value: T
  setValue: (value: T) => void
}

const Dropdown = <T extends { name: string }>({
  url,
  value,
  setValue,
}: DropdownProps<T>) => {
  const [items, setItems] = useState<T[]>([])
  const [error, setError] = useState<string | null>(null)
  const [isOpen, setIsOpen] = useState<boolean>(false)

  useEffect(() => {
    const fetchData = async () => {
      try {
        await api.get(url).then(response => setItems(response.data))
      } catch (err) {
        setError((err as Error).message)
      }
    }
    fetchData()
  }, [url])

  const handleClick = (item: T) => {
    setValue(item)
    setIsOpen(false)
  }

  return (
    <div className="relative">
      <div
        className="relative bg-white rounded-md border border-zinc-300 h-8 px-2 py-1"
        onClick={() => setIsOpen(!isOpen)}
      >
        {value.name}
        <ChevronDown className="absolute right-2 top-1" />
      </div>
      {isOpen && (
        <div className="absolute left-0 mt-2 w-full bg-zinc-100 overflow-hidden border rounded-lg shadow-lg">
          {error && <p className="text-red-500 p-2">{error}</p>}
          {!error && (
            <ul>
              {items.map((item, index) => (
                <li
                  key={index}
                  className="px-4 py-1 hover:bg-blue-600 hover:text-white cursor-pointer"
                  onClick={() => handleClick(item)}
                >
                  {item.name}
                </li>
              ))}
            </ul>
          )}
        </div>
      )}
    </div>
  )
}

export default Dropdown
