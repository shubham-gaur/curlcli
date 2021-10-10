package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"syscall/js"
)

func main() {
	c := make(chan bool)

	fmt.Println("Hello from WebAssembly")
	js.Global().Set("increment", js.FuncOf(jsIncr))
	js.Global().Set("formatJSON", js.FuncOf(jsFormatJSON))
	<-c
}

func jsIncr(this js.Value, args []js.Value) interface{} {

	jsDoc := js.Global().Get("document")
	nameElement := jsDoc.Call("getElementById", "welcome")
	nameElement.Set("textContent", "Hi Shubham")
	outElement := jsDoc.Call("getElementById", "notification-count")
	input := outElement.Get("textContent")
	fmt.Println("Input: ", input.String())
	incr := increment(input.String())
	fmt.Println("Output: ", incr)
	outElement.Set("textContent", incr)
	return incr
}
func increment(input string) string {
	int_input, _ := strconv.Atoi(input)
	int_input++
	return strconv.Itoa(int_input)
}

func jsFormatJSON(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return "Invalid no of arguments passed"
	}
	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		return "Unable to get document object"
	}
	inputJSON := args[0].String()
	fmt.Printf("non prettified json %s\n", inputJSON)
	prettyJSON, err := prettyJson(inputJSON)
	if err != nil {
		errStr := fmt.Sprintf("unable to parse JSON. Error %s occurred\n", err)
		result := map[string]interface{}{
			"error": errStr,
		}
		return result
	}
	fmt.Printf("prettified json %s\n", prettyJSON)
	jsonOuputTextArea := jsDoc.Call("getElementById", "jsonoutput")
	if !jsonOuputTextArea.Truthy() {
		return "Unable to get output text area"
	}
	fmt.Printf("prettified json %s\n", prettyJSON)
	jsonOuputTextArea.Set("textContent", prettyJSON)
	return nil
}

func prettyJson(input string) (string, error) {
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}
