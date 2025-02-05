declare global {
  interface User {
    id: number;
    username: string;
    password: string;
  }

  interface Address {
    id: number;
    address: string;
    userId: number;
    user: User;
    blockchainId: number;
    blockchain: Blockchain;
  }

  interface Blockchain {
    id: number;
    name: string;
  }

  interface Asset {
    id: number;
    name: string;
  }

  interface Address {
    id: number;
    assetId: string;
    asset: Asset;
    price: number;
    lastUpdate: string;
  }

  interface Address {
    addressId: number;
    address: Address;
    priceId: number;
    price: Price;
    value: number;
    avgPurchasedPrice: number;
    gain: number;
    percentageGain: number;
    actualValue: number;
  }

  interface AuthContextType {
    user: User | null;
    login: (username: string, password: string) => Promise<void>;
    register: (username: string, password: string) => Promise<void>;
    logout: () => void;
  }
}

export {};
