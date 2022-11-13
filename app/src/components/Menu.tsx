import { Web3Modal, useAccount, useConnectModal } from "@web3modal/react";

import { providers } from "@web3modal/ethereum";

import type { ConfigOptions } from "@web3modal/core";
import Safe from "./Safe";

const config: ConfigOptions = {
  projectId: "2cdad8abe02ff3f7d6b8bf7a928001b3",
  theme: "light",
  accentColor: "blackWhite",
  ethereum: {
    appName: "Web3dApp",
    chains: [
      {
        id: 31337,
        name: "Hardhat 31337",
        network: "localhost",
        rpcUrls: { default: "http://127.0.0.1:8545" },
      },
    ],
    providers: [
      providers.jsonRpcProvider({
        rpc: () => ({ http: "http://127.0.0.1:8545" }),
      }),
    ],
  },
};

const MenuAppBar = () => {
  const { account } = useAccount();
  const { isOpen, open, close } = useConnectModal();

  return (
    <>
      <Web3Modal config={config} />

      {account.address && (
        <>
          {account.address.slice(0, 6)}..{account.address.slice(-4)}
        </>
      )}
      <input
        type="button"
        onClick={() => {
          if (isOpen) {
            return close();
          }
          open();
        }}
        value="click"
      />
      <Safe owner={account.address} />
    </>
  );
};

export default MenuAppBar;
