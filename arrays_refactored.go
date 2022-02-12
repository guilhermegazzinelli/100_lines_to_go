package main 

import "fmt"

func main(){
	DeploymentOptions := [...]string{"R-Pi", "AWS", "GCP", "Azure"}

	for index, option := range DeploymentOptions{
		fmt.Println(index, option)
	}
}
