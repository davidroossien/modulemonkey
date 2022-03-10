package main

import (
	"fmt"

	ss "github.com/davidroossien/modulemonkey/semanticsort"
	ll "github.com/davidroossien/modulemonkey/simplelog"
)

/*
Used for testing various prototype modules.
*/
func main() {
	fmt.Println("Start...")

	// use the log package
	returnValue, err := ll.Config("log.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Log returnValue = " + fmt.Sprint(returnValue))

	// use the semanticVersion package

	inputSemanticVersions := [5]string{"10.0.1", "5.2.0", "1.1.1", "1.3.0", "1.2.11"}

	initalizedSemanticVersions, err := ss.Initialize(inputSemanticVersions[:])
	if err != nil {
		fmt.Println(err.Error())
	}

	sortedSemanticVersions := ss.SemanticSort(initalizedSemanticVersions)

	fmt.Println("Sorted semantic versions:")

	// print the result
	for i := 0; i < len(sortedSemanticVersions); i++ {
		fmt.Println(sortedSemanticVersions[i].Version)
	}

	fmt.Println("Done...")
}
