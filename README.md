# E-Voting System
> A secure electronic voting system built with Go, featuring user authentication, digital signatures, and blockchain integration.  

## Features
- **User Authentication**: Secure login system with username, password, and citizen ID verification
- **Digital Signature Verification**: Ensures vote integrity through cryptographic verification
- **Admin Verification**: Admins can verify users through a secure signature process
- **Blockchain Integration**: Records votes on a blockchain for immutability and transparency
- **Voting Server Management**: Supports multiple voting servers and candidates

## Architecture
The system follows a clean architecture approach with:  
- Service layer for business logic
- Repository layer for data access
- JWT-based authentication
- Password encryption
- Logging system

## API Endpoints
The system exposes several endpoints for:  
- User registration and login
- User verification by admins
- Vote submission with cryptographic verification
- User information retrieval

## Security Features
- Password encryption
- Digital signature verification
- JWT token authentication
- Citizen ID verification
- Public/private key pairs for users

## Requirements
- Go 1.23.+
- PostgresSQL database
- Aptos blockchain connection (for vote recording)

## Setup and Installation
    ```bash
    go mod tidy
    go mod vendor

    go run cmd/main.go api -c configs/config.yaml
    ```