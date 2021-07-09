# Tokopedia TCO Feature

Repository for Tokopedia DevCamp 2021 Hackathon project - Group 02

## Folder Structure
```
.
|
├── delivery    # HTTP delivery (API)
├── fixtures    # For seed and migrate
├── repository  # all database logic
├── usecase     # all business logic
│   |
.   .
```

## Getting Started
### Install dependencies
#### Go
https://golang.org/doc/install/source

### Setup Database
```./docker-compose up```

### Build
`go build`

### Migrate and Seed Database
- ```./tokopedia-hackathon migrate```
- ```./tokopedia-hackathon seed```

### Run
```./tokopedia-hackathon```
