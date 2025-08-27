# Go-Quiz-Portal-Backend

*Coded by:*
**Prajwal Krishna**

This **README** file contains :
 1. Information About the Quiz-App
 2. How to run the backend server
 3. Controls for game play
 4. File structure
 5. List of APIs

----------


About The Quiz Portal
-------------

This a backend written in GoLang using **Gin** as microframework **Gorm** 
as ORM and **modernc.org/sqlite** as database.

This server is caters to restful apis of a quiz app whose frontend is tota
lly decoupled with it.

List of apis and a breif description about those are written in API sectio
n of README.md.


----------

## Running the program

- Install GoLang in your machine.
- Inside GoLang src folder clone this repo
- Go inside the folder
- Running the program is easy
         go run
- This starts a go-server at **localhost:8080**
- The database is intialized with some values to make it empty run followi
ng commands before go run
         rm gorm.db
         touch gorm.db

   This starts a backend server for quiz app on localhost:8080 which can c
ater to any service making a api request to it.
_______________

#### Prajwal Krishna Maitin
