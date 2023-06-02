// video: https://www.youtube.com/watch?v=YS4e4q9oBaU&t=12053s
// 3:51:48

package main

import (
	"fmt"
	// "io/ioutil"
	// "log"
	// "net/http"
)

func main() {
	/////////////////////////////////////////////
	// For loop

	// for i := 0; i < 5; i++ {
	// 	if i == 4 {
	// 		fmt.Println("Dont print 4")
	// 		break
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }

	// // Nested For loop
	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		fmt.Printf("i: %d", i)
	// 		fmt.Printf(" j: %d\n", j)
	// 	}
	// }

	// // MAP
	// statePopulations := map[string]int{}
	// statePopulations["key1"] = 5
	// statePopulations["key2"] = 19

	// // MAP ITERATION
	// for i := range statePopulations {
	// 	if i == "key2" {
	// 		fmt.Println("Found key!")
	// 	}
	// }
	// // or this way
	// for i, j := range statePopulations {
	// 	fmt.Println(i, j)
	// }

	/////////////////////////////////////////////
	// Defer, Panic, Recovery

	// Defer

	// defer causes it to print last
	// defer functions are LIFO order, so in reverse order from defer
	// calls: the first defer call will be last and latest will be
	// firtst.

	// fmt.Println("start")
	// defer fmt.Println("middle")
	// defer fmt.Println("2ndmiddle")
	// fmt.Println("end")
	// prints: start -> end -> 2ndmiddle -> middle

	// HTTP Request Example //////////////////////////////////////

	// Sends an http 'get' request to specified url, the results are
	// stored in "res" and any error is stored in "err"
	// res, err := http.Get("http://google.com/robots.txt")

	// // Checks to see if there was an error
	// if err != nil {
	// 	// Print error message and terminate the program
	// 	log.Fatal(err)
	// }

	// // 'ioutil.ReadAll' function: read the body of the http request
	// // it reads all the data from 'res.Body' and stores it in robots
	// // Any error encountered is stored in 'err'
	// robots, err := ioutil.ReadAll(res.Body)

	// // After reading the res.Body, its good to close it
	// defer res.Body.Close()

	// // Checks to see if there was an error
	// if err != nil {
	// 	// Print error message and terminate program
	// 	log.Fatal(err)
	// }

	// // Prints the contents of robots which from before will be the
	// // body of the http request
	// fmt.Printf("%s", robots)

	////////////////////////////////////////////////////////////////////

	// Panic

	// Web Application Example /////////////////////////////////////////

	// Handler function: a function that processes incoming HTTP requests
	// and generates corresponding HTTP responses. In this case
	// "func(w http.ResponseWriter, r *http.Request){...}" is the handler
	// function.

	// "Register a handler function": Associating a handler function
	// with a specific route or URL pattern. This allows the server
	// to invoke the appropriate handler when an incoming request
	// matches the route.

	// HandleFunc allows you to associate a given route or URL pattern
	// with a specific handler (handles http requests & responses).

	// HandleFunc takes two args:
	// Arg 1: a string representing the URL pattern or route.
	// Specifies the path or pattern that incoming requests should match.
	// Arg 2: the handler function that will be called to handle
	// requests.

	//	w': represents an 'http.ResponseWriter' interface. Used
	//	to write the HTTP response back to the client.
	//	'r': represents an '*http.Request' pointer. Holds the
	//	information about the incoming HTTP request.

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Go!"))
	// })

	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	fmt.Println("An error happened")
	// 	panic(err.Error())
	// }

	// // Should print 'Hello Go!' at localhost:8080

	// fmt.Println("Process successful")

	// fmt.Println("start")
	// defer fmt.Println("this was defered")
	// panic("something bad happened")
	// fmt.Println("end")
	////////////////////////////////////////////////////////////////////

	// Pointers
	// a := 42
	// var b *int = &a
	// c := &a
	// fmt.Println(a, *b, *c)
	// a = 27
	// fmt.Println(a, *b, *c)

	////////////////////////////////////////////////////////////////////

	//Functions

	// stringVar := "Hello Go: "
	// for i := 0; i < 5; i++ {
	//	// Function call passing it a (string, int)
	// 	sayMessage(stringVar, i)
	// }

	// var1 := 3
	// var2 := 6

	// fmt.Printf("The sum of %d and %d is: ", var1, var2)
	// fmt.Println(sumVar(var1, var2))

	////////////////////////////////////////////////////////////////////

	// Slice

	// // Slide is like array list
	// c := []int{}

	// // Print slice, its length, and capacity
	// fmt.Println(c)
	// fmt.Println(len(c))
	// fmt.Println(cap(c))

	// // add on 1 to the end of c
	// c = append(c, 1)
	// fmt.Println(c)

	// // add on 5,6,7,8 to the end of c
	// c = append(c, 5, 6, 7, 8)
	// doubleValueInSlice(c)

	////////////////////////////////////////////////////////////////////

	// Linked List Example
	// completeList := LinkedList{}

	// displayList(completeList)

	// insertNode(10, &completeList)
	// insertNode(30, &completeList)
	// insertNode(20, &completeList)
	// insertNode(20, &completeList)
	// insertNode(70, &completeList)
	// insertNode(60, &completeList)

	// deleteNode(30, &completeList)

	// displayList(completeList)

	////////////////////////////////////////////////////////////////////

}

// Struct
type Node struct {
	number int
	next   *Node
}

type LinkedList struct {
	head *Node
}

// Functions

func deleteNode(val int, completeList *LinkedList) bool {
	if completeList.head.number == val {
		completeList.head = completeList.head.next
	}

	for temp := completeList.head; temp.next != nil; temp = temp.next {
		// Found the value to delete
		if temp.next.number == val {
			connectionSave := temp.next.next
			prevSave := temp
			temp = temp.next
			temp = nil
			prevSave.next = connectionSave

			return true
		} else if temp.next.number > val {
			// Value doesnt exist
			return false
		}
	}
	return false
}

// Insert into linked list
func insertNode(val int, completeList *LinkedList) {
	// if list is empty
	if completeList.head == nil {
		newNode := &Node{number: val, next: nil}
		completeList.head = newNode
		return
	}

	temp := completeList.head

	for temp.next != nil {
		if temp.next.number < val {
			temp = temp.next
		} else if temp.next.number > val {
			newNode := &Node{number: val, next: nil}
			// 1->5->10
			// Save the current next node
			connectionSaver := temp.next
			// Connect the list with new node
			temp.next = newNode
			// Add on second part of the list
			newNode.next = connectionSaver

			// Finsihed inserting, so return
			return
		} else {
			fmt.Println("Duplicate Value Not Allowed!")
			return
		}
	}
	newNode := &Node{number: val, next: nil}
	temp.next = newNode

}

// Display Linked List
func displayList(listVar LinkedList) {
	if listVar.head == nil {
		fmt.Println("Empty list")
		return
	}

	for temp := listVar.head; temp != nil; temp = temp.next {
		fmt.Println(temp.number)
	}
}

// Purpose: Display the two args
// Arg(s): Two different variables
func sayMessage(msg string, number int) {
	fmt.Println(msg, number)
}

// Purpose: Return sum of variables
// Arg(s): Two variables
func sumVar(var1 int, var2 int) int {
	return var1 + var2
}

// Purpose: Double the element values in array
// Arg(s): An array
func doubleValueInSlice(sliceArray []int) {
	fmt.Println("Original array: ", sliceArray)

	for i := 0; i < len(sliceArray); i++ {
		fmt.Println(sliceArray[i])
	}
}
