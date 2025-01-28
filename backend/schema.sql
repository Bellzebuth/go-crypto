CREATE TABLE IF NOT EXISTS assets (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    key_name INTEGER NOT NULL,
    amount REAL NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    buying_price INTEGER NOT NULL,
    FOREIGN KEY (key_name) REFERENCES cryptos(key_name)
);

CREATE TABLE IF NOT EXISTS cache_prices (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    key_name INTEGER NOT NULL,
    price INTEGER,
    last_update DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (key_name) REFERENCES cryptos(key_name)
);

CREATE TABLE IF NOT EXISTS cryptos (
    key_name TEXT NOT NULL PRIMARY KEY,
    name TEXT NOT NULL
);

INSERT INTO cryptos (key_name, name) VALUES
('bitcoin', 'Bitcoin'),
('ethereum', 'Ethereum'),
('tether', 'Tether'),
('binancecoin', 'BNB'),
('usd-coin', 'USD Coin'),
('ripple', 'XRP'),
('cardano', 'Cardano'),
('dogecoin', 'Dogecoin'),
('solana', 'Solana'),
('tron', 'TRON'),
('litecoin', 'Litecoin'),
('polkadot', 'Polkadot'),
('polygon', 'Polygon'),
('shiba-inu', 'Shiba Inu'),
('dai', 'Dai'),
('wrapped-bitcoin', 'Wrapped Bitcoin'),
('uniswap', 'Uniswap'),
('avalanche', 'Avalanche'),
('chainlink', 'Chainlink'),
('leo-token', 'LEO Token'),
('cosmos', 'Cosmos'),
('the-open-network', 'Toncoin'),
('monero', 'Monero'),
('stellar', 'Stellar'),
('okb', 'OKB'),
('bitcoin-cash', 'Bitcoin Cash'),
('trueusd', 'TrueUSD'),
('filecoin', 'Filecoin'),
('internet-computer', 'Internet Computer'),
('hedera', 'Hedera'),
('lido-dao', 'Lido DAO'),
('arbitrum', 'Arbitrum'),
('aptos', 'Aptos'),
('quant-network', 'Quant'),
('cronos', 'Cronos'),
('vechain', 'VeChain'),
('algorand', 'Algorand'),
('maker', 'Maker'),
('elrond', 'MultiversX'),
('the-sandbox', 'The Sandbox'),
('eos', 'EOS'),
('immutable-x', 'Immutable X'),
('aave', 'Aave'),
('decentraland', 'Decentraland'),
('tezos', 'Tezos'),
('theta-token', 'Theta Network'),
('axie-infinity', 'Axie Infinity'),
('fantom', 'Fantom'),
('stacks', 'Stacks'),
('flow', 'Flow'),
('neo', 'NEO'),
('huobi-token', 'Huobi Token'),
('kucoin-shares', 'KuCoin Token'),
('sui', 'Sui'),
('curve-dao-token', 'Curve DAO'),
('gmx', 'GMX'),
('bitdao', 'BitDAO'),
('convex-finance', 'Convex Finance'),
('rocket-pool', 'Rocket Pool'),
('frax', 'Frax'),
('thorchain', 'THORChain'),
('paxos-standard', 'Pax Dollar'),
('pancakeswap', 'PancakeSwap'),
('usdd', 'USDD'),
('zcash', 'Zcash'),
('iota', 'IOTA'),
('render-token', 'Render'),
('kaspa', 'Kaspa'),
('loopring', 'Loopring'),
('1inch', '1inch'),
('enjincoin', 'Enjin Coin'),
('mina-protocol', 'Mina Protocol'),
('bittorrent', 'BitTorrent'),
('gala', 'Gala'),
('floki', 'Floki'),
('gemini-dollar', 'Gemini Dollar'),
('nexo', 'NEXO'),
('trust-wallet-token', 'Trust Wallet Token'),
('stepn', 'STEPN'),
('mask-network', 'Mask Network'),
('compound-governance-token', 'Compound'),
('baby-doge-coin', 'Baby Doge Coin'),
('holotoken', 'Holo'),
('optimism', 'Optimism'),
('decred', 'Decred'),
('ravencoin', 'Ravencoin'),
('theta-fuel', 'Theta Fuel'),
('oasis-network', 'Oasis Network'),
('bitcoin-sv', 'Bitcoin SV'),
('tether-gold', 'Tether Gold'),
('just', 'JUST'),
('dash', 'Dash'),
('balancer', 'Balancer'),
('kusama', 'Kusama');

INSERT INTO cache_prices (key_name, price) VALUES
('bitcoin', 75200000000),    
('ethereum', 2950000000),     
('tether', 0930000),     
('binancecoin', 205000000),      
('usd-coin', 0500000),     
('ripple', 0610000),     
('cardano', 5900000),     
('dogecoin', 0070000),     
('solana', 23500000),    
('tron', 164000000),  
('litecoin', 550000000),  
('polkadot', 0850000),    
('polygon', 88000000),   
('shiba-inu', 0450000),    
('dai', 6800000),    
('wrapped-bitcoin', 4200000),    
('uniswap', 0070000),    
('avalanche', 23500000),   
('chainlink', 164000000),  
('leo-token', 550000000),  
('cosmos', 0850000),    
('the-open-network', 88000000),   
('monero', 0450000),    
('stellar', 6800000),    
('okb', 4200000),    
('bitcoin-cash', 0070000),    
('trueusd', 23500000),   
('filecoin', 164000000),  
('internet-computer', 550000000),  
('hedera', 0850000),    
('lido-dao', 88000000),   
('arbitrum', 0450000),    
('aptos', 6800000),    
('quant-network', 4200000),    
('cronos', 0070000),    
('vechain', 23500000),   
('algorand', 164000000),  
('maker', 550000000),  
('elrond', 0850000),    
('the-sandbox', 88000000),   
('eos', 0450000),    
('immutable-x', 6800000),    
('aave', 4200000),    
('decentraland', 0070000),    
('tezos', 23500000),   
('theta-token', 164000000),  
('axie-infinity', 550000000),  
('fantom', 0850000),    
('stacks', 88000000),   
('flow', 0450000),    
('neo', 6800000),    
('huobi-token', 4200000),    
('kucoin-shares', 0070000),    
('sui', 23500000),   
('curve-dao-token', 164000000),  
('gmx', 550000000),  
('bitdao', 0850000),    
('convex-finance', 88000000),   
('rocket-pool', 0450000),    
('frax', 6800000),    
('thorchain', 4200000),    
('paxos-standard', 0070000),    
('pancakeswap', 23500000),   
('usdd', 164000000),  
('zcash', 550000000),  
('iota', 0850000),    
('render-token', 88000000),   
('kaspa', 0450000),    
('loopring', 6800000),    
('1inch', 4200000),    
('enjincoin', 0070000),    
('mina-protocol', 23500000),   
('bittorrent', 164000000),  
('gala', 550000000),  
('floki', 0850000),    
('gemini-dollar', 88000000),   
('nexo', 0450000),    
('trust-wallet-token', 6800000),    
('stepn', 4200000),    
('mask-network', 0070000),    
('compound-governance-token', 4200000),    
('baby-doge-coin', 0070000),    
('holotoken', 23500000),   
('optimism', 164000000),  
('decred', 550000000),  
('ravencoin', 0850000),    
('theta-fuel', 88000000),   
('oasis-network', 0450000),    
('bitcoin-sv', 6800000),    
('tether-gold', 4200000),    
('just', 0070000),    
('dash', 23500000),   
('balancer', 164000000),  
('kusama', 550000000);



