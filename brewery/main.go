package main;

import "fmt"

func GetMap(cityNumber int,tab [6][2]int)map[int][]int{

	mainCity:=make(map[int][]int,0)
	
	for i:=0;i<cityNumber;i++{
		cityNum:=i+1;
	
		mainCity[cityNum]=make([]int,0)

		mainCity[cityNum]=append(mainCity[cityNum],tab[i][0])

		mainCity[cityNum]=append(mainCity[cityNum],tab[i][1])
	}
	return mainCity;
}
func CountCityValue(city int,mapCity map[int][]int)int{
	
	allCostForCity:=0;
	for k,_:=range(mapCity){
		
		if(k!=city){

			countRounde:=CountCityRoad(city,k,mapCity)

			allCostForCity=allCostForCity+countRounde;
			
		}
	}

	return allCostForCity;
}

func CountCityRoad(fromCity int,destCity int,mapCity map[int][]int) int{

	cost:=mapCity[destCity][0];

	roundOneSum:=0;
	
	roundTwoSum:=0;

	mapLenght:=len(mapCity);
	
	for i:=fromCity;i!=destCity;i++{
		if(i>mapLenght){
			i=1;
		}
		if(i==destCity){
			break;
		}
		roundOneSum=roundOneSum+mapCity[i][1];
	
	}
	
	for j:=fromCity;j!=destCity;{
		j--;
		if(j<1){
			j=mapLenght;
		}
		
		index:=j;
		if(index==mapLenght){
			index=j-1;
		}

		roundTwoSum=roundTwoSum+mapCity[j][1];
		if(j==destCity){
			break;
		}
		
	}
	roundOneSum=roundOneSum*cost;

	roundTwoSum=roundTwoSum*cost;

	if(roundOneSum<roundTwoSum){
		return roundOneSum;
	}


	return roundTwoSum;
}

func main(){
	fmt.Println("Start");
	
	var cityNumber int=6;
	
	var enterTab [6][2]int=[6][2]int{
		{1,2},{2,3},{1,2},{5,2},{1,10},{2,3},
	}
	mainCity:=GetMap(cityNumber,enterTab)

	sumValue:=0;
	selectCity:=0;

	
	for k,_ :=range mainCity{

		if (selectCity==0){
			selectCity=k;
		}
		
		sum:=CountCityValue(k,mainCity)
		
		fmt.Printf("City count %d %d \n",k,sum)
		if(sumValue==0 || sum<sumValue){

			sumValue=sum;
			selectCity=k;
		}
	}

	fmt.Printf("city %d have cost %d \n",selectCity,sumValue)
	return;
}