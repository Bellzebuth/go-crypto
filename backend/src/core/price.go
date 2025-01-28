package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Bellzebuth/go-crypto/src/db"
	"github.com/Bellzebuth/go-crypto/src/utils"
)

var priceURL = "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum,tether,binancecoin,usd-coin,ripple,cardano,dogecoin,solana,tron,litecoin,polkadot,polygon,shiba-inu,dai,wrapped-bitcoin,uniswap,avalanche,chainlink,leo-token,cosmos,the-open-network,monero,stellar,okb,bitcoin-cash,trueusd,filecoin,internet-computer,hedera,lido-dao,arbitrum,aptos,quant-network,cronos,vechain,algorand,maker,elrond,the-sandbox,eos,immutable-x,aave,decentraland,tezos,theta-token,axie-infinity,fantom,stacks,flow,neo,huobi-token,kucoin-shares,sui,curve-dao-token,gmx,bitdao,convex-finance,rocket-pool,frax,thorchain,paxos-standard,pancakeswap,usdd,zcash,iota,render-token,kaspa,loopring,1inch,enjincoin,mina-protocol,bittorrent,gala,floki,gemini-dollar,nexo,trust-wallet-token,stepn,mask-network,compound-governance-token,baby-doge-coin,holotoken,optimism,decred,ravencoin,theta-fuel,oasis-network,bitcoin-sv,tether-gold,just,dash,balancer,kusama&vs_currencies=eur"

func UpdateCryptoPrices() error {
	resp, err := http.Get(priceURL)
	if err != nil {
		return fmt.Errorf("failed to fetch price: %w", err)
	}

	defer resp.Body.Close()

	now := time.Now()
	if resp.StatusCode == 200 {
		var result map[string]map[string]float64
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return fmt.Errorf("failed to parse price response: %w", err)
		}

		for keyName, currencies := range result {
			for _, price := range currencies {
				// update cache
				query := `UPDATE cache_prices SET price = ?, last_update = ? WHERE key_name = ?`

				_, err = db.DB.Exec(query, utils.ConvertToMicroUnits(price), now, keyName)
				if err != nil {
					return fmt.Errorf("failed to insert price for %s: %w", keyName, err)
				}
			}
		}
	} else {
		return fmt.Errorf("failed request with status code %d", resp.StatusCode)
	}

	return nil
}
