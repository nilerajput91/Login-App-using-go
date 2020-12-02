# Login-App-using-go


Why Go?
Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.

Prerequisites.
To follow along you will need to have these installed.

Go version 1.13, It is the version used to build this API.

PostgreSQL installed.

An IDE of choice, some of the recommended IDEs are listed here.

An API Request builder. e.g Postman.


Api Specifications.
Endpoints we will have at the end of this article.

Register a user after a valid POST request at /register.
Login user after a valid POST request at /login.


Getting Dependencies.
These are the package dependencies we will need.

badoux/checkmail - for validating user emails.
dgrijalva/jwt-go - to sign and verify jwt tokens.
gorilla/mux - it is a router and dispatcher, for matching URLs to their handlers.
jinzhu/gorm - its an ORM(Object Relational Mapper) will enable us interact with the database.
joho/godotenv - to load .env file, which holds our secrets. e.g database passwords.
crypto - to hash and verify user passwords.
To install these dependencies, open the terminal and type go get github.com/{package-name} e.g go get github.com/badoux/checkmail .

For crypto installation, type go get golang.org/x/crypto .

