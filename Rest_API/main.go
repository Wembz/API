package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"

)

//Define struct of program

type todo struct {
	ID        string    `json: "id"`
	Item      string	`json: "item"`
	Completed bool		`json: "completed"`
}

// define variable using a slice
var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

//How to GET variables from slice 
func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}

// how to ADD data from the client
func addTodo(context *gin.Context){
	
	//new variable to "ADD"
	var newTodo todo

	// how to BIND new variable to JSON
	if err := context.BindJSON(&newTodo); err !=nil{
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}


//A function that the handler will use 
func getTodo(context *gin.Context){
	
	//extract ID from path Param 
	id := context.Param("id")
	todo, err := getTodoById(id)

	//error alert if Todo doesnt exist
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
// todo alert message
	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context){
	
	//extract ID from path Param 
	id := context.Param("id")
	todo, err := getTodoById(id)

	//error alert if Todo doesnt exist
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	// flip boolean value if true flip to false and vice versa
	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

// How to get a SPECIFIC ID in a function
func getTodoById(id string)(*todo, error){
	for i, t := range todos{
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo Not found")
}

//Create the main function to run the app
func main (){
	
	//build our server
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)

	//local server host
	router.Run("localhost:9090")

}