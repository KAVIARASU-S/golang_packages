package main

import "fmt"

type link struct {
 prev *link
 val int
 next *link
}

var front *link
var back *link
var count=0

func addb(val int) {
 count++
 insert := &link{nil, val, nil}
 if back == nil {
  front = insert
  back = insert
 }
 insert.prev = back
 back.next = insert
 back = insert
}
func addf(val int) {
 count++
 insert := &link{nil, val, nil}
 if back == nil {
  front = insert
  back = insert
 }
 insert.next = front
 front.prev = insert
 front = insert
}
func addm(val int,pos int){
 var variable*link=front
 count++
 insert := &link{nil, val, nil}
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
func deleteb() {
 count--
 back = back.prev

}
func deletem(pos int){
	
     
  if(pos==1){
   deletef()
  }else if(pos==count){
               deleteb()
  }else if(pos<=0){
   fmt.Println("Enter correct value")
  }else{
  pos=pos-1
       var variable*link=front
    var supp*link
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
 var variable*link=front
 for i:=0;i<count;i++{
  fmt.Println(variable.val)
        variable=variable.next
 }
}
func main() {
 addf(23)
            addb(45)
          display()

}
