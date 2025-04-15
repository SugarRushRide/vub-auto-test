//  -
// Created: 2025/4/15

package selectors

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	selectors, err := Load("selectors.yaml")
	if err != nil {
		t.Fatal(err)
	}
	btn, err := Get(selectors, "overall_position.data_export")
	fmt.Printf("Loaded selectors: %#v\n", selectors)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(btn)
}
