package main

// import logger "E:/0. CrawlApps/Projects/GoLang/logger/logger.go"

import "GoLang/logger"

// import (
// 	"fmt"
// 	"log"
// )

// func main() {
	// fmt.Println("You are good to go with GO")
	// log.Println("You are good to go with GO")

	// log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	// log.Println("setted up the more details to the logs")

	// log.Panic("Your codehave some errors.")
	
// }

func main()  {
	
	logger.Info("This is Info logger's log")
	logger.Warn("This is Warn logger's log")
	logger.Error("This is Error logger's log")
	
	logger.SetLevel(logger.ErrorLevel)
	
	logger.Info("This is Info logger's log")
	logger.Warn("This is Warn logger's log")
	logger.Error("This is Error logger's log")
}
	