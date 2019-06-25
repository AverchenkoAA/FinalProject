# FinalProject
You should have launched mongo data base server on localhost:27017.
Data base and collections will create automatically.
Main.go for web-server you can find in folder "project".
Server and client run via "go run main.go".

Console-client should be used after web-server app start only.
Default user "ADMIN" has password "12345678".
If you want to find or change some fields you have to use special syntax of the fields.
For change some computer's fields you can use follow commands:

computer.inventorynumber --- integer - Inventory numder of the computer

computer.hddvolume --- integer - HDD volume of the computer

computer.ramvolume --- integer - RAM volume of the computer

computer.vendor --- string - Vendor of the computer

computer.core.frequency --- float64 - Frequency of core in the computer

computer.core.corevendor --- string - Vendor of core in the computer

computer.core.model --- string - Model of core in the computer

computer.owner.firstname --- string - First name of a person who use the computer

computer.owner.lastname --- string - Lastname name of a person who use the computer

computer.owner.roomnumber --- integer - Room number of a person who use the computer
