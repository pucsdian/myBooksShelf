# My Books Shelf

This is kind of Books E-Commerce web application where you can buy any books. <br />
This project is built for Programming Paradigm Course 2021.

## About
The myBooksShelf is a web application built using [React.js](https://reactjs.org/) as frontend and [Golang](https://golang.org/) as backend. <br />
The web application is made taking consideration of the following problem.<br />
As there are many E-Commerce platforms which are selling everythings, sometimes user may gets distract by seeing other things as well. They spend time by seeing other thing and later they reliease that they come to buy a book and they have wasted time on some other things.<br />
Well people who are interested to read books are precisely concern about their time so why not to save the minutes to buy a books in easiest way such that our product also get sold out and user also gets what books he/she need quickly.

### Prerequisites
```
Docker
Docker-Compose
```

### Installation

1. [Docker](https://www.docker.com/):
If you don't have docker install, please install it from [Here](https://docs.docker.com/engine/install/).
1. [Docker-Compose](https://docs.docker.com/compose/):
If you don't have docker-compose install, please install it from [Here](https://docs.docker.com/compose/install/).


## Getting Started

1. Clone the project.
2. cd myBooksShelf
3. Execute `make run` command (please wait this might take 3-4 minutes). If you don't have `make` installed, please execute `docker-compose build && docker-compose up -d` commands to start the containers.
4. Be sure that your machine is available for ports 8080, 8081 and 3306
5. Please wait while all containers get started
6. Use `docker ps` command to see whether all containers are started or not
7. Make sure that these three `users`, `books`, `mysql_db` containers are running
8. If they are working please open `http://localhost:3000` in any browser of your machine
9. Use `make stop` command to stop and delete all the running containers. If you don't have `make` installed, please execute `docker-compose down` command to stop the container.



## Current Services
&nbsp; &nbsp; &nbsp; &nbsp; Currently,  myBooksShelf has the following active services
```
----------------------------------------------------------------------------------------------
|                     |                                                                      |
|     Service Name    |                              Description                             |
|                     |                                                                      |
---------------------------------------------------------------------------------------------|
|  User Login         |        helps to login user                                           |
----------------------------------------------------------------------------------------------
|  User SignUp        |        helps to signup user                                          |
----------------------------------------------------------------------------------------------
|  Search books       |        helps to search the books by it's name or by author name      |
----------------------------------------------------------------------------------------------
|  Add books to       |        where user can add books to the cart                          |
|    the cart         |                                                                      |
----------------------------------------------------------------------------------------------
|  Remove books from  |        where user can remove books from the cart                     |
|     the cart        |                                                                      |
----------------------------------------------------------------------------------------------       
|      Logout         |        helps to logout the logged in user                            |
----------------------------------------------------------------------------------------------
```

## API
&nbsp; &nbsp; &nbsp; &nbsp; Currently, myBooksShelf built with following API which are written in [golang](https://golang.org/) version 1.13.<br />
&nbsp; &nbsp; &nbsp; &nbsp; All those services are running on the host `127.0.0.1` which is `localhost`

```
--------------------------------------------------------------------------------------------------------------------------
|                |          |                     |                                                                      |
|     Method     |   PORT   |      Endpoints      |                              Description                             |
|                |          |                     |                                                                      |
-------------------------------------------------------------------------------------------------------------------------|
|      POST      |   8080   |     /api/login      |  If user is valid, it will responses with PASETO token               |
--------------------------------------------------------------------------------------------------------------------------
|      POST      |   8080   |     /api/signup     |  If username or email is not already present, it creates account     |
--------------------------------------------------------------------------------------------------------------------------
|      GET       |   8081   |     /api/books      |  It search the books basis on what user typed in                     |
--------------------------------------------------------------------------------------------------------------------------
|      GET       |   8081   |  /api/randombooks   |  It gives as random 52 books details from database                   |
--------------------------------------------------------------------------------------------------------------------------
|      POST      |   8081   |   /api/addtocart    |  It adds the books to the users carts by specifying bookid's         |
--------------------------------------------------------------------------------------------------------------------------       
|      GET       |   8081   |   /api/addtocart    |  It gives us the user's cart items list                              |
--------------------------------------------------------------------------------------------------------------------------
```




## Built With Stack

* [Golang](https://golang.org/) - The language is used for implementation(golang version 1.13)
* [Gin-Gonic](https://gin-gonic.com/) - The web framework used to write API at backend
* [Go Mod](https://golang.org/ref/mod#introduction) - The package dependency management in golang
* [GORM](https://gorm.io) - The fantastic ORM library for Golang
* [PASETO](https://paseto.io/) - Used to generate token for user authentication(more secure than JWT)
* [Viper](https://github.com/spf13/viper) - Used to read configuration variables from .env files
* [Golang-Testing](https://golang.org/pkg/testing/) - The unit testing framework is used to test API
* [React.js](https://reactjs.org/) - A JavaScript library for building user interfaces, version 17.0.2 is used
* [React Hooks](https://reactjs.org/docs/hooks-reference.html#gatsby-focus-wrapper) - They let you use state and other React features without writing a class
* [React-Router](https://reactjs.org/docs/code-splitting.html#route-based-code-splitting) - Route-based code splitting
* [HTML](https://html.com/) - Used to create elements of the web pages
* [CSS]() - Used to design HTML elements
* [Dockerfile]() - Used to build custom images, each service has it's own dockerfile
* [Docker-Compose]() - To build project as dockerized application
* [MySQL](https://www.mysql.com/) - Database, MySQL version 8.0.23 is used

## Versioning
Version 1.0

