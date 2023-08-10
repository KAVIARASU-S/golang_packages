import "fmt"
type center struct{
 left*center
 val int
 right*center
}
var binary*center
var enter=1
var foun=1
var level=1
func insert(num int){
 insert:=&center{nil,num,nil}
        if(binary==nil){
 binary=insert
  }else{
   var temp*center=binary
   for true{
    if(temp.val>=num){
     if(temp.left==nil){
      temp.left=insert
      break
     }else{
      temp=temp.left
     }
    }else{
     if(temp.right==nil){
      temp.right=insert
      break
     }else{
      temp=temp.right
     }
    }
   }
  }
}
func search(temp*center,digit int){
 if(temp==nil){
  return 
 }
 if temp.val>binary.val && enter==1 {
  enter =0
  level=2
 }
 if(temp.val==digit) {
  foun=0
  fmt.Println("Element Found",level)
  return 
 }
	
     level++
        search(temp.left,digit)
  search(temp.right,digit)
}
func main(){
 insert(11)
             insert(45)
             insert(56)
             insert(98)
            search(11)
	
	
}
