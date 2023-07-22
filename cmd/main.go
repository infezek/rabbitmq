package main

import (
	"os"
	"rabbitmqgo/pkg"
)

func Execute() {
	err := pkg.RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	pkg.RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func main() {
	Execute()
}
