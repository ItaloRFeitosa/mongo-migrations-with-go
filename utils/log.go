package utils 

import (
	"fmt"
 	"time"
)

func Log(data interface{}){
	currentTime := time.Now()
	stringCurrentTime := "["+currentTime.Format(time.RFC3339)+"] "
	fmt.Println(stringCurrentTime + fmt.Sprint(data))
}