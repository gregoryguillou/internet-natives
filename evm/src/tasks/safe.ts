import { task } from "hardhat/config";
import EthersAdapter from "@gnosis.pm/safe-ethers-lib";
import Safe, { SafeFactory, SafeAccountConfig } from "@gnosis.pm/safe-core-sdk";

task("safe", "Create a simple set of MNEMONIC").setAction(
  async (_, { ethers }) => {
    const [user1, user2, user3] = await ethers.getSigners();
    const ethAdapter = new EthersAdapter({
      ethers: ethers,
      signer: user1,
    });
    const safeFactory = await SafeFactory.create({ ethAdapter });
    const owners = [user1.address, user2.address, user3.address];
    const threshold = 3;
    const safeAccountConfig: SafeAccountConfig = {
      owners,
      threshold,
    };
    const safeSdk: Safe = await safeFactory.deploySafe({ safeAccountConfig });
    console.log("address", await safeSdk.getAddress());
    console.log("version", await safeSdk.getContractVersion());
    console.log("owners", await safeSdk.getOwners());
  }
);

export {};
