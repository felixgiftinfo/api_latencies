package main

import (
	"log"

	"github.com/felixgiftinfo/api_latencies/utils"
)

func main() {
	log.Println("Hello API Latency")
	filePath := utils.FileAccess{
		PathToDirectory: "./files",
	}
	filePath.Run([]int{50, 60, 90, 1000})
}
