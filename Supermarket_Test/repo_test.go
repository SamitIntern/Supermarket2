package Supermarket_Test

import (
	"testing"
)


func TestRepoCreateRepo(t *testing.T) {
	todos = RepoCreateRepo()
	if len(todos) != 4 {
		t.Errorf("RepoCreateRepo was incorrect, got: %d elements, wanted: %d.", len(todos), 4)
	}
}

func TestCheckForExistingProducePositive(t *testing.T) {
	var todo Todo
	//var todos Todos
	todos = RepoCreateRepo()
	todo = todos[0]
	todo = CheckForExistingProduce(todo)
	if todo.Id != 1{
		t.Errorf("CheckForExistingProduce was incorrect, got: %d as ID, wanted: %d.", todo.Id, 1)
	}else{
		t.Log("CheckForExistingProduce was correct, got: 1 as ID, wanted: 1.")
	}
}


func TestCheckForExistingProduceNegative(t *testing.T) {
	var todo Todo

	todos = RepoCreateRepo()
	//todo = todos[0]
	todo.Id = 7
	todo.ProduceCode = "1111"
	todo.Name = "Nothing"
	todo.UnitPrice = 100
	todo = CheckForExistingProduce(todo)

	if todo.Id != 0{
		t.Errorf("Negative tests for CheckForExistingProduce was incorrect, got: %d as ID, wanted: %d.", todo.Id, 0)
	}else{
		t.Log("Negative tests for CheckForExistingProduce was correct, got: 0 as ID, wanted: 0.")
	}
}


func TestCheckForNonExistingProduceNegative(t *testing.T) {
	var todo Todo

	todos = RepoCreateRepo()
	todo = CheckForNonExistingProduce(7)

	if todo.Id != 0{
		t.Errorf("Negative tests for CheckForNonExistingProduce was incorrect, got: %d as ID, wanted: %d.", todo.Id, 0)
	}else{
		t.Log("Negative tests for CheckForNonExistingProduce was correct, got: 0 as ID, wanted: 0.")
	}
}

func TestCheckForNonExistingProducePositive(t *testing.T) {
	var todo Todo

	todos = RepoCreateRepo()
	todo = CheckForNonExistingProduce(2)

	if todo.Id != 2{
		t.Errorf("Positive tests for CheckForNonExistingProduce was incorrect, got: %d as ID, wanted: %d.", todo.Id, 0)
	}else{
		t.Log("Positive tests for CheckForNonExistingProduce was correct, got: 2 as ID, wanted: 2.")
	}
}

func TestRepoFindItemPositive(t *testing.T) {
	var todo Todo

	todos = RepoCreateRepo()
	todo = RepoFindItem(2)

	if todo.Id != 2{
		t.Errorf("Positive tests for RepoFindItem was incorrect, got: %d as ID, wanted: %d.", todo.Id, 0)
	}else{
		t.Log("Positive tests for RepoFindItem was correct, got: 2 as ID, wanted: 2.")
	}
}

func TestRepoFindItemNegative(t *testing.T) {
	var todo Todo

	todos = RepoCreateRepo()
	todo = RepoFindItem(7)

	if todo.Id != 0{
		t.Errorf("Negative tests for RepoFindItem was incorrect, got: %d as ID, wanted: %d.", todo.Id, 0)
	}else{
		t.Log("Negative tests for RepoFindItem was correct, got: 0 as ID, wanted: 0.")
	}
}

func TestRepoCreateItemPositive(t *testing.T) {
	var todo Todo
	var tempId = currentId
	todos = RepoCreateRepo()
	todo.ProduceCode = "1111"
	todo.Name = "Nothing"
	todo.UnitPrice = 100
	todo = RepoCreateItem(todo)
	if todo.Id != (tempId+1){
		t.Errorf("Positive tests for RepoCreateItem was incorrect, got: %d as ID, wanted: %d.", todo.Id, tempId+1)
	}else{
		t.Log("Positive tests for RepoCreateItem was correct, got: %d as ID, wanted: %d.", tempId+1, tempId+1)
	}
}

func TestRepoCreateItemNegative(t *testing.T) {
	var todo Todo
	var tempId = currentId
	todos = RepoCreateRepo()
	todo.ProduceCode = "1111"
	todo.Name = "Nothing"
	todo.UnitPrice = 100
	todo = RepoCreateItem(todo)
	if todo.Id == 0{
		t.Errorf("Positive tests for RepoCreateItem was incorrect, got: %d as ID, wanted: %d.",0, tempId+1)
	}else{
		t.Log("Positive tests for RepoCreateItem was correct, got: %d as ID, wanted: %d.", tempId+1, tempId+1)
	}
}

func TestRepoDeleteItemNegative(t *testing.T) {

	var contains bool
	contains = false
	todos = RepoCreateRepo()
	todos = RepoDeleteItem(1)

	for _, t2 := range todos {
		if t2.Id == 1{
			contains = true
		}
	}
	if contains{
		t.Errorf("Negative tests for RepoDeleteItem was incorrect, got produce with %d as ID, wanted it to be deleted.", 1)
	}else{
		t.Log("Negative tests for RepoDeleteItem was correct, got produce with 1 as ID deleted.")
	}
}

func TestRepoDeleteItemPositive(t *testing.T) {

	var count int

	todos = RepoCreateRepo()
	count = len(todos)
	todos = RepoDeleteItem(1)

	if count == len(todos){
		t.Errorf("Positive tests for RepoDeleteItem was incorrect, the element was not deleted")
	}else{
		t.Log("Positive tests for RepoDeleteItem was correct, got produce deleted.")
	}
}

func TestGetAllItems(t *testing.T) {
	todos = RepoCreateRepo()
	if len(todos) != 4{
		t.Errorf("Test for GetAllItems was incorrect, did not get the correct number of elements")
	}else{
		t.Log("Test for GetAllItems was correct, got the correct number of elements.")
	}
}

func TestCheckValidAddParamNumberNamesNegative(t *testing.T) {
	var todo Todo
	var valid bool
	todo.ProduceCode = "A12T-4GH7-QPL9-3N4M"
	todo.Id = 1
	todo.UnitPrice = 2
	valid = CheckValidAddParamNumber(todo)
	if valid{
		t.Errorf("Test for CheckValidAddParamNumber was incorrect, expected it to fail but it passed")
	}else{
		t.Log("Test for CheckValidAddParamNumber was correct, expected it to fail and it failed")
	}
}

func TestCheckValidAddParamNumberNamesPositive(t *testing.T) {
	var todo Todo
	var valid bool
	todo.ProduceCode = "A12T-4GH7-QPL9-3N4M"
	todo.Id = 1
	todo.UnitPrice = 2
	todo.Name = "TestCase"
	valid = CheckValidAddParamNumber(todo)
	if !valid{
		t.Errorf("Test for CheckValidAddParamNumber was incorrect, expected it to pass but it failed")
	}else{
		t.Log("Test for CheckValidAddParamNumber was correct, expected it to pass and it passed")
	}
}

func TestCheckValidAddParamNumberProduceCodePositive(t *testing.T) {
	var todo Todo
	var valid bool
	todo.ProduceCode = "A12T-4GH7-QPL9-3N4M"
	todo.Id = 1
	todo.UnitPrice = 2
	todo.Name = "TestCase"
	valid = CheckValidAddParamNumber(todo)
	if !valid{
		t.Errorf("Test for CheckValidAddParamNumber was incorrect, expected it to pass but it failed")
	}else{
		t.Log("Test for CheckValidAddParamNumber was correct, expected it to pass and it passed")
	}
}

func TestCheckValidAddParamNumberProduceCodeNegative(t *testing.T) {
	var todo Todo
	var valid bool
	todo.Id = 1
	todo.UnitPrice = 2
	todo.Name = "TestCase"
	valid = CheckValidAddParamNumber(todo)
	if valid{
		t.Errorf("Test for CheckValidAddParamNumber was incorrect, expected it to fail but it passed")
	}else{
		t.Log("Test for CheckValidAddParamNumber was correct, expected it to fail and it failed")
	}
}

func TestCheckValidAddParamTypesCodeLengthPositive(t *testing.T) {
	var todo Todo
	var valid bool
	todo.ProduceCode = "A12T-4GH7-QPL9-3N4M"
	todo.Id = 1
	todo.UnitPrice = 2
	todo.Name = "TestCase"
	valid = CheckValidAddParamNumber(todo)
	if !valid{
		t.Errorf("Test for CheckValidAddParamTypes was incorrect, expected it to pass but it failed")
	}else{
		t.Log("Test for CheckValidAddParamTypes was correct, expected it to pass and it passed")
	}
}

func TestCheckValidAddParamTypesCodeLengthNegative(t *testing.T) {
	var todo Todo
	var valid bool
	todo.ProduceCode = "A12T-4GH7-QPL9"
	todo.Name = "TestCase"
	todo.Id = 1
	todo.UnitPrice = 2
	valid = CheckValidAddParamNumber(todo)
	if valid{
		t.Errorf("Test for CheckValidAddParamTypes was incorrect, expected it to fail but it passed")
	}else{
		t.Log("Test for CheckValidAddParamTypes was correct, expected it to fail and it failed")
	}
}

func TestCheckValidAddParamTypesCodeStructurePositive(t *testing.T) {
	var todo Todo
	var valid bool
	todo.ProduceCode = "A12T-4GH7-QPL9-3N4M"
	todo.Name = "TestCase"
	todo.Id = 1
	todo.UnitPrice = 2
	valid = CheckValidAddParamNumber(todo)
	if !valid{
		t.Errorf("Test for CheckValidAddParamTypes was incorrect, expected it to pass but it failed")
	}else{
		t.Log("Test for CheckValidAddParamTypes was correct, expected it to pass and it passed")
	}
}

func TestCheckValidAddParamTypesCodeStructureNegative(t *testing.T) {
	var todo Todo
	var valid bool
	todo.ProduceCode = "A12T-4GH7-Q-L9"
	todo.Name = "TestCase"
	todo.Id = 1
	todo.UnitPrice = 2
	valid = CheckValidAddParamNumber(todo)
	if valid{
		t.Errorf("Test for CheckValidAddParamTypes was incorrect, expected it to fail but it passed")
	}else{
		t.Log("Test for CheckValidAddParamTypes was correct, expected it to fail and it failed")
	}
}