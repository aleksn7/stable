# Stable coin project

This project demonstrates a stable coin;

### Run docker
```
docker-compose up -d
```

### Attach to container
```
docker exec -it $(docker-compose ps -q) bash 
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
