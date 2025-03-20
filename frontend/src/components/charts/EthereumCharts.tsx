import { useEffect, useRef, memo } from "react"

const TradingViewWidget: React.FC = () => {
  const container = useRef<HTMLDivElement | null>(null)

  useEffect(() => {
    if (!container.current) return

    container.current.innerHTML = ""

    const script = document.createElement("script")
    script.src =
      "https://s3.tradingview.com/external-embedding/embed-widget-advanced-chart.js"
    script.type = "text/javascript"
    script.async = true
    script.innerHTML = JSON.stringify({
      autosize: true,
      symbol: "COINBASE:ETHUSD",
      interval: "D",
      timezone: "Europe/Paris",
      theme: "light",
      style: "2",
      locale: "en",
      hide_legend: true,
      allow_symbol_change: true,
      support_host: "https://www.tradingview.com",
    })

    container.current.appendChild(script)
  }, [])

  return (
    <div
      className="tradingview-widget-container"
      ref={container}
      style={{ height: "100%", width: "100%" }}
    >
      <div
        className="tradingview-widget-container__widget"
        style={{ height: "100%", width: "100%" }}
      ></div>
    </div>
  )
}

export default memo(TradingViewWidget)
