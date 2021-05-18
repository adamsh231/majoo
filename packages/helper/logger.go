package helper

import "fmt"

func Log(err error, logDescription string) {
	if err != nil {
		fmt.Println("Error:", err.Error(), "Description:", logDescription)
	}
}

func LogOnly(err string, logDescription string){
	fmt.Println("Error:", err, "Description:", logDescription)
}
