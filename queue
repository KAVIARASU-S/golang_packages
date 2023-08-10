package main

import "fmt"

type queue struct {
 prev *queue
 num int
 next *queue
}

var front *queue
var back *queue
var count=0

func addb(num int) {
 count++
 insert := &queue{nil, num, nil}
 if back == nil {
  front = insert
  back = insert
 }
 insert.prev = back
 back.next = insert
 back = insert
}
func addf(num int) {
 count++
 insert := &queue{nil, num, nil}
 if back == nil {
  front = insert
  back = insert
 }
 insert.next = front
 front.prev = insert
 front = insert
}
func addm(num int,pos int){
 var variable*queue=front
 count++
 insert := &queue{nil, num, nil}
 if back == nil {
  front = insert
  back = insert
 }else{
    for i:=0;i<pos-2;i++{
             variable=variable.next
     }
  fmt.Println("Well")
  insert.next=variable.next
  variable.next.prev=insert
  variable.next=insert
  insert.prev=variable

}
}
func deletef() {
 count--
 front = front.next
 if front == nil {
  back = nil
 }
}
func backdele() {
 count--
 back = back.prev

}
func deletem(pos int){
	
     
  if(pos==1){
   deletef()
  }else if(pos==count){
               backdele()
  }else if(pos<=0){
   fmt.Println("Enter correct numue")
  }else{
  pos=pos-1
       var variable*queue=front
    var supp*queue
    for true{
     pos--
    if(pos==0){
   break
    }
    variable=variable.next
    
    }
    supp=variable
    variable=variable.next
    supp.next=variable.next
    variable=variable.next
    variable.prev=supp
 }
 count--

}
func display() {
 var variable*queue=front
 for i:=0;i<count;i++{
  fmt.Println(variable.num)
        variable=variable.next
 }
}
func main() {
 addf(23)
            addb(45)
          display()

}
