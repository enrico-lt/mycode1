/* Alta3 Research | RZFeeser
   Template - HTML template  */

package main

import (
	"html/template"
	"os"
)

type Todo1 struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo1
}

func main() {

	// check our template for potential errors with Must
	tmpl := template.Must(template.ParseFiles("C:\\data\\repo\\mycode1\\day04\\index.html"))

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo1{
			{Title: "Laundry", Done: false},
			{Title: "Pull weeds in the garden", Done: true},
			{Title: "Sweep the stairs", Done: true},
		},
	}
	tmpl.Execute(os.Stdout, data)

}
