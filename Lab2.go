package main

import (
    "net/http"
	"github.com/gorilla/mux"
	"log"
	"fmt"
	"encoding/json"
	"io"
	"io/ioutil"
)

var currentId int

type Todo struct {

    Name      string    `json:"name"`
   
}

type Respo struct {

    Greeting      string    `json:"greeting"`
   
}
type Todos []Todo


type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}


func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Hello,", todoId)
}




type Routes []Route

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return router
}

var routes = Routes{
   
    Route{
        "TodoShow",
        "GET",
        "/hello/{todoId}",
        TodoShow,
    },
	Route{
		"TodoCreate",
		"POST",
		"/hello",
		TodoCreate,
	},
	

}

func main() {

    router := NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	var respo Respo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
		respo.Greeting = "Hello, " + todo.Name + "!"
	//t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(respo); err != nil {
		panic(err)
	}
}
