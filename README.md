# Stable coin project

This project demonstrates a stable coin.

### Code coverage 
```
Network Info
============
> HardhatEVM: v2.12.2
> network:    hardhat



  Stable
    Fund
      ✔ Should be funded (854ms)
      ✔ Should be reverted (80ms)
    Defund
      ✔ Should be defunded (76ms)
      ✔ Should be reverted (88ms)
    Liquidate
      ✔ Should be liquidated (98ms)
      ✔ Should be reverted (69ms)


  6 passing (1s)

-----------------|----------|----------|----------|----------|----------------|
File             |  % Stmts | % Branch |  % Funcs |  % Lines |Uncovered Lines |
-----------------|----------|----------|----------|----------|----------------|
 contracts/      |    93.94 |    53.13 |    83.87 |    89.58 |                |
  Coin.sol       |      100 |       50 |      100 |      100 |                |
  Controller.sol |      100 |       50 |      100 |      100 |                |
  Core.sol       |    89.47 |    54.76 |    78.57 |    86.89 |... 170,171,173 |
  Math.sol       |      100 |       50 |    71.43 |    71.43 |          24,32 |
-----------------|----------|----------|----------|----------|----------------|
All files        |    93.94 |    53.13 |    83.87 |    89.58 |                |
-----------------|----------|----------|----------|----------|----------------|
```

## Run

### Run docker
```
docker-compose up -d
```

### Attach to container
```
docker exec -it stable-hardhat-1 bash 
```

### CD to the project directory
```
cd home/devproj
```

### Try running some of the following tasks:
```
npx hardhat help
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat node
npx hardhat run scripts/deploy.ts
```
