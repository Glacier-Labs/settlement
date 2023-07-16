go run .
solc --evm-version london  --bin --abi -o ./abi ./contract/zkglacier_g16.sol
abigen --bin ./abi/Verifier.bin --abi abi/Verifier.abi --pkg contract --out ./contract/contract.go --type GlacierVerifier
cp ./contract/zkglacier_g16.sol ../contracts/