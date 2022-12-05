My selected capstone project is to create a CRUD using SQL for a user generated database for collections and item managment. The project will either create a new database or utilize an existing one to log and store data of various items. 

I will demonstrate the following in this project:
1. Building a CRUD application
2. Integration of sqlite3 into golang
3. Error handling 
4. Building a network database to house multiple type of items (i.e. books, film, vinyl, etc.)

-----User Stories-----
EXAMPLE #1
As a collector, I would like to keep and record of my vinyl collection so I can privately store what I have, making it easier to track of everything in my collection when I am shopping at  second-hand stores.

After running the program, the user will be prompted to select from a small series of options to select the desired database. It will list all avaialbe databases currently attached to the program, as well as an option to create a new database. Databases will be listed as below:

Please Select a Databse:
1. Books
2. Film
3. Vinyl
4. CDs
5. 'Create New Database'

The selection will be made by simply entering the associated integer and pressing enter. The program will then give a second menu option that function along the same lines as the previous. In this example, the user had chosen 3:

Vinyl Collection
Please Select an action
1. Add new item 
2. Modify existing item
3. Delete item
4. Browse database 
5. Return to database selection

After selecting 1, Add new item, and new prompt will appear, one category at a time. It will first ask for a Title to be entered. This will add the new item to the databse with a auto generated id number. Each table would have its own unique categories based on what the database is for. For the Vinyl database, it will prompt for these categories:
1. Title*
2. Artist
3. Release Year
4. Genre
5. Special Edition
6. Description

For simplicity, the Title is the only field required to add a new item. The other can be left blank by simply hitting enter when prompted, and this will add a empty string in it's place. 

--------------
EXAMPLE #2
A small business owner wants to keep track of their inventory for their hobby shop. Using the program, they can can log what kind of items that they keep in stock, and update their quantity.

As a shop owner, using only one databse to keep track of everything, they would run the program to update their existing inventory. It would work very similar to the Vinyl Collection database, but the Inventory Database would house everything they carry. To add a quantity to an item, they would run the program and select the desired option:
1. Modify Quantity
2. Search Inventory 
3. Add New Item
4. Remove Item

It would be a very basic and simple interface that willbe easy to use once setup. The shopkeeper just received a small box of some yarn and they need to add it to there inventory. It contains 12 skeins, half of which is a new product.
