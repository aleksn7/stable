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
      ✔ Should be reverted (83ms)
    Defund
      ✔ Should be defunded (81ms)
      ✔ Should be reverted (118ms)
    Liquidate
      ✔ Should be liquidated (133ms)
      ✔ Should be reverted (69ms)


  6 passing (1s)

-----------------|----------|----------|----------|----------|----------------|
File             |  % Stmts | % Branch |  % Funcs |  % Lines |Uncovered Lines |
-----------------|----------|----------|----------|----------|----------------|
 contracts/      |    86.11 |    44.74 |    72.22 |     81.9 |                |
  Coin.sol       |      100 |       50 |      100 |      100 |                |
  Controller.sol |      100 |       50 |      100 |      100 |                |
  Core.sol       |    89.47 |    54.76 |    78.57 |    86.89 |... 170,171,173 |
  Math.sol       |    33.33 |       20 |    41.67 |    31.25 |... 51,52,54,55 |
-----------------|----------|----------|----------|----------|----------------|
All files        |    86.11 |    44.74 |    72.22 |     81.9 |                |
-----------------|----------|----------|----------|----------|----------------|

> Istanbul reports written to ./coverage/ and ./coverage.json
```

## Run

### Run docker
```
docker-compose up -d
```

### Attach to container
```
docker exec -it dai-hardhat-1 bash 
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
