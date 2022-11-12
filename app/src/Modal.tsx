import { Web3Modal, Web3Button } from "@web3modal/react";
import { providers } from "@web3modal/ethereum";

import type { ConfigOptions } from "@web3modal/core";

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

const Modal = () => {
  return (
    <>
      <Web3Modal config={config} />
      <Web3Button />
    </>
  );
};

export default Modal;
