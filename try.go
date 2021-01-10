package main

import "fmt"


//Implement the structure of a node data type
type node struct {
	data string
	nextElement *node
}
//Implement the structure of the list
type linkedList struct {
	head *node
	length int
}
//Add items to the list
func (myList *linkedList) addNodeToList(n *node) {
	//Create pointer from the current new head node to the next item in list
	n.nextElement = myList.head
	//Setup the new head in our list
	myList.head = n
	//point new node at original head node that we moved forward 1 spot
	//update the length of this list
	myList.length++
}


//delete items from the list

func (myList *linkedList) deleteNodeFromList(node int) {
	currentHead := myList.head
	for i := 1;i < node - 1;i++ {
	currentHead = currentHead.nextElement
	}
	currentHead.nextElement = currentHead.nextElement.nextElement
	myList.length--
}
// see what's in the list
func (myList linkedList) listNodes() {
	printMe := myList.head
	for myList.length > 0 {
		fmt.Printf("%v\n", printMe.data) 
		printMe = printMe.nextElement
		myList.length--
	}
}
// see individual node
//func populateList(data []string) {
//	for i := 0;i < len(data);i++ {	
//	ourList.addNodeToList(&node{data: data[i]})	
//	}
//}


func main() {
	ourList := linkedList{}
	
	ourList.addNodeToList(&node{data: "one"})

	ourList.addNodeToList(&node{data: "two"})
	ourList.addNodeToList(&node{data: "three"})

	ourList.addNodeToList(&node{data: "four"})

	ourList.deleteNodeFromList(3)
	//dataList := []string{"one", "two", "three"}
	//populateList(dataList)
	fmt.Printf("%+v\n", ourList)
	ourList.listNodes()
}

