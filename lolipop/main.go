/*
5 3 //numer of leght of lolipop, how many number
TWTWT //one lolipope to share T cost 2 and W 1
5 //customer money
1
7
//resut
1 3
2 1
NIE
*/

package main;

import ("fmt"
"bufio"
"strings"
"os"
"strconv")

func CheckLolip(lolipop string,customerMoney int) (bool,int,int){
	
	for index,_:=range lolipop{

		var actualPrice int;

		for innerIndex,s := range lolipop[index:]{

			actualPrice+=GetTasteValue(string(s))

			if actualPrice==customerMoney{

				return true,index+1,index+innerIndex+1;

			}else if actualPrice>customerMoney{

				break;
			}

		}
	
	}
	return false,0,0;
}
func GetTasteValue(character string)int{
	
	if character=="T"{
		return 2;
	}else{
		return 1;
	}
}

func GetFirstLine(text string)(int,int){

	firstLine:=( strings.Split( text," "))

	lenghtLolipop,_:=strconv.Atoi(firstLine[0]);

	customer,_:=strconv.Atoi(firstLine[1]);
	
	return lenghtLolipop,customer;
}


func main(){

	fmt.Println("Start")
	
	reader := bufio.NewScanner(os.Stdin)

	reader.Scan()
	lenghtLolipop,customer :=GetFirstLine(reader.Text())
	
	var lolipop string;
	fmt.Scanln(&lolipop)

	var tabCustomer []int=make([]int,0,0);

	if(len(lolipop)<lenghtLolipop){
		return;
	}

	for i:=0;i<(customer);i++{
		
		var valuCustomer int;

		fmt.Scanln(&valuCustomer)

		tabCustomer=append(tabCustomer,valuCustomer)
	}

	for _,price :=range tabCustomer{

		hasValue,start,lengh:=CheckLolip(lolipop,price);

		if hasValue==true{

			fmt.Printf("%d %d \n",start,lengh)
		}else{

			fmt.Println("NIE")

		}

	}
return;
}