# GolangCRUDTemplatingWeb
#### There is an endpoint with an little explanation:
* "/", this endpoint will render the file in the views folder, "index.html", as the main page of a website.

_in the "index.html" file, you will see a QUERY (READ) records view from a database with the connection status of a database.
Underneath there is a form that will be filled in using the POST method which will be forwarded to the main.go server and the data will be brought to the database through CUD (Create/Update/Delete) operations.
You will be given information on the failure or success of the CUD (Create/Update/Delete) operation._

#### Here some little explanation for the buttons:

* The buttons on the form will refresh (in HTML, action="/") to the "index.html" file, so that the records from the database will always be updated. 
* The buttons are separated by being given a value, so Golang will know which button to press (Presence/Absence/Delete).
