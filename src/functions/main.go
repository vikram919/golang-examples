package main

import (
	"fmt"
	"../structs"
)

func main() {
	fmt.Print("First function in function folder!!\n");
	fmt.Printf("The employee name is: %v", testStringConcatenation("vikram", "gopu"));
	a,b := 5,6
	fmt.Printf("Expected sum of two numbers is: %v\n", getSum(&a, &b));
	fmt.Printf("Employee details are: %v\n", packageItemsInAnArray("vikram", "gopu", "1994", "1stMay"));
	fmt.Printf("Employee details are: %v\n", packageItemsInSlice("vikram", "gopu", "1994", "1stMay"));
	employeeMap := getItemsInAMap();
	for k,v := range employeeMap {
		fmt.Printf("Employee name: %v and Employee email: %v\n",k,v );
	}
	circle := structs.NewCircle(0,5);
	fmt.Printf("Area of the circle is: %v\n", structs.GetArea(circle));
}

func testStringConcatenation(name,familyName string ) string {
	return name + familyName;
}

func getSum(num1, num2 *int) int {
    return *num1+*num2;
}

func packageItemsInAnArray(name, familyName, dob, startDate string) [4]string {
return [4]string{name, familyName, dob, startDate};
}

func packageItemsInSlice(name, familyName, dob, startDate string) []string {
	return []string{name, familyName, dob, startDate};	
}

func getItemsInAMap() map[string]string {
	return map[string]string{"vikram":"v.gopu@gmail.com", "employee2":"employee@gmail.com"}
}