/*
Sample enters
6
9
1 2
2 3
2 5
1 4
1 6
3 4
4 5
5 6
3 6
*/

package main;

import ("fmt")

func FindBlackTriangle(allGraph map[int][]int,maxPoint int) map[int][]int{

	blackGraph:=make(map[int][]int);

	for i:=1;i<=maxPoint;i++{

		blackGraph[i]=make([]int,0)

		for j:=1;j<maxPoint;j++{
			
			if i==j{
				continue;
			}
			if ExistElement(allGraph[i],j)==false{
				blackGraph[i]=append(blackGraph[i],j)
			}
		}
	}

	return blackGraph;
}
func ExistElement(tab []int,el int) bool {
	for _,val :=range tab{
		if val==el{
			return true;
		}
	}
	return false;
}
func CanAddUniqueArray(arraySource [][]int,arrayToAdd []int)bool{
	if(len(arraySource)==0){
		return true;
	}
	for _,tab:=range arraySource{
		
		counter:=0;
		for _,el :=range arrayToAdd{
			
			if(ExistElement(tab,el)==true){

				counter++;
			}
		}
		if(len(arrayToAdd)==counter){
			return false;
		}
		
	}

	return true;
}


func main(){
	
	fmt.Println("Start")
	var maxPoints int;
	var maxLines int;
	//var grayTrangle [][]int;

	fmt.Scan(&maxPoints, &maxLines)
	
	var allGraph map[int][]int

	allGraph=make(map[int][]int, 0)
	
	for i:=1;i<=maxPoints;i++{
		
		allGraph[i]=make([]int,0);
	}
	//fill graph
	for i:=0;i<maxLines;i++{
		var k int;
		var v int;

		fmt.Scan(&k,&v)

		allGraph[k]=append(allGraph[k],v)
		
		allGraph[v]=append(allGraph[v],k)
	}

	fmt.Println("black graph")

	grayGraph:= (FindBlackTriangle(allGraph,maxPoints))

	response:=make([][]int,0);



	for k,v :=range(grayGraph){
		if len(v)<2{
			continue;
		}
		collect:=make([]int,0)

		collect=append(collect,k);

		for _,value:=range v{
			collect=append(collect,value)
		}	
		if (CanAddUniqueArray(response,collect)==true){
		response=append(response,collect)
		}
	}
	fmt.Println(len(response));

	fmt.Println("End")	
	return;
}