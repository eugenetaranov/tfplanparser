package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hashicorp/terraform/terraform"
)

func main() {
	filePath := flag.String("path", "", "Path to the terraform plan")
	flag.Parse()

	if _, err := os.Stat(*filePath); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist\n", *filePath)
		return
	}

	f, _ := os.Open(*filePath)
	defer f.Close()

	plan, err := terraform.ReadPlan(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(plan)

	resources := plan.Diff.Modules[0].Resources

	for k, v := range resources {
		fmt.Println(k, v.Attributes)
	}

}
