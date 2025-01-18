# Consensus-Breaking Change
A consensus-breaking change is a change made such that blockchain nodes running different versions of the software no longer agree on the state of the blockchain, leading to potential chain split. 

## Minimum Value Requirement
The consensus-breaking change I've implemented is that assets must have a minimum value of 100 when created / updated. This is a xonsensus-breaking change because nodes running the old version will not enforce this rule and allow assets valued under 100, while updated nodes will reject them as invalid.

---
## Running the application
From the assetregistry folder:
```bash
ignite chain build; ignite chain serve
```

---
## CRUD Operations
```bash
/Users/abi/go/bin/assetregistryd tx assetregistry create-asset "Valid Asset" "Valued over 100" alice 150 --from alice --chain-id assetregistry

/Users/abi/go/bin/assetregistryd tx assetregistry create-asset "Invalid Asset" "Valued under 100" alice 75 --from alice --chain-id assetregistry

/Users/abi/go/bin/assetregistryd q assetregistry list-asset

/Users/abi/go/bin/assetregistryd q assetregistry show-asset 0

/Users/abi/go/bin/assetregistryd tx assetregistry update-asset 0 "Seiko 5" "Seiko 5 SNZG15 field watch, purchased in Singapore" alice 75 --from alice --chain-id assetregistry

/Users/abi/go/bin/assetregistryd tx assetregistry delete-asset 0 --from alice --chain-id asset
```
