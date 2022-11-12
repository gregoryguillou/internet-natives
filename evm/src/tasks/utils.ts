import { Wallet } from "ethers";
import { task } from "hardhat/config";

task("mnemonic", "generate a random set of MNEMONIC").setAction(async () => {
  const mnemonic = Wallet.createRandom().mnemonic;
  console.log("MNEMONIC:", mnemonic);
});

task("refill", "Create a simple set of MNEMONIC").setAction(
  async (_, { ethers }) => {
    const [user1] = await ethers.getSigners();
    const tx = {
      to: "0xb93e876cDe2c07a446b3c5849d1dD4DCC4cB8F0f",
      value: ethers.utils.parseEther("2"),
    };
    const transaction = await user1.sendTransaction(tx);
    const receipt = await transaction.wait(1);
    console.log("receipt", receipt);
  }
);

export {};
