go run .
solc --evm-version london  --bin --abi -o ./abi ./contract/zkglacier_g16.sol --combined-json abi,opcodes
abigen --bin ./abi/Verifier.bin --abi abi/Verifier.abi --pkg contract --out ./contract/contract.go --type GlacierVerifier
cp ./contract/zkglacier_g16.sol ../contracts/
solc  --bin --abi -o ./abi ./contract/zkglacier_g16.sol --combined-json abi,opcodes
solc  --bin --abi -o ./abi ./contract/zkglacier_g16.sol --opcodes

npx hardhat test test/glacier_g16.js --network localhost
npx hardhat test test/zk_g16.js --network localhost

npx hardhat run scripts/deploy_zk.js --network opbnbtestnet
npx hardhat run scripts/deploy_gov.js --network opbnbtestnet
npx hardhat run scripts/deploy.js --network opbnbtestnet
npx hardhat run scripts/test_opbnb.js  --network opbnbtestnet


npx hardhat run scripts/deploy_gov.js --network NautilusTestnet
