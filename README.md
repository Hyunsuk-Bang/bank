# Simple Bank

## Description
Simple Bank Project to practice skills

## Getting Started

### Dependencies
* migrate-golang - tools for DB migration
* sqlc - CRUD code generator for golang 
* viper - Environment Variable management

### Executing program
* spawn postgres
    make postgres

* create DB
    make create_db

* drop DB
    make drop_db

* migrate up DB
    make migrate_up
    
* migrate down DB
    make migrate_down

### Topic Covered
1. PostgreSQL
2. CRUD with GO
3. Testing for CRUD
4. ACID Properties
    - Atomicity
    - Consistency
    - Isolation
    - Durability
