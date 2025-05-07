
# Hyperledger Fabric Asset Management System

This project is part of an internship assignment to build a blockchain-based asset management system using Hyperledger Fabric. It includes:

- *Smart contract (chaincode)* in Golang to manage financial asset accounts
- *REST API* built with Node.js to interact with the blockchain network
- *Dockerfile* to containerize the API
- Structured for clarity and ease of deployment

The system supports creating, updating, and querying assets, ensuring security, transparency, and traceability of transactions.

This project is built for managing financial asset accounts using Hyperledger Fabric. It includes:

- Fabric test network setup (based on `test-network`)
- Chaincode written in Golang
- Node.js REST API to interact with the blockchain
- Dockerfile for API containerization

## Project Structure

- `chaincode/go/` — Smart contract logic (Golang)
- `rest-api/` — Express-based API server
- `Dockerfile` — To containerize the API

## Steps to Run

1. Setup the Fabric test network
2. Deploy the chaincode
3. Run the REST API (`node app.js`)
4. Use Postman or curl to hit the API endpoints

---
