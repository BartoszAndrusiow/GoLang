package main;
import (
	"fmt" 
	"bufio"
	"os"
	"strings"
	"strconv"
)
//7 3 -> numer roller, how many will be will
//5 6 4 3 6 2 3 -> roller space
//3 2 5 -> roller in
func ReadConsoleToArray(arrayLenght int)[]int{
	
	reader := bufio.NewScanner(os.Stdin)

	reader.Scan()
	text:=reader.Text();

	tab:=make([]int,arrayLenght);

	for index,value:=range strings.Split(text, " "){

		tab[index],_=strconv.Atoi(value)
	}

	return tab;
	
}

func GetLastRollerPos(roller []int, rollerThrow []int)int{
	var actualPoz int=len(roller);

	if(roller[0]<rollerThrow[0]){
		return 0;
	}

	for _,v :=range rollerThrow{

		for i:=0;i<len(roller);i++{

			if(v>roller[i]){
				actualPoz=i;
				break;
			}
		}
		actualPoz=actualPoz-1;
	}
	return actualPoz+1;
}
func main(){

	fmt.Println("Start")
	var rollerNumber int;
	var rollerInNumber int;

	tabIn:=ReadConsoleToArray(2)

	rollerNumber,rollerInNumber=tabIn[0],tabIn[1];

	roller:=ReadConsoleToArray(rollerNumber)
	
	rollerThrow:=ReadConsoleToArray(rollerInNumber)

//	roller:=[]int{5, 6 ,4 ,3 ,6 ,2, 3};
//	rollerThrow:=[]int{3 ,2 ,5};

	
	fmt.Println(GetLastRollerPos(roller,rollerThrow));

	fmt.Println("End")
	return;
}