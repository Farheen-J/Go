# Go

1. Go
- Statically typed language - we have to specify data type while assigning variable. Although, it can be inferred from the right hand side value. 
- Dynamically typed (JavaScript, Python, etc)
- Declaration of varibles:
    1. var card string = "Ace of Spades"
	2. card := "Ace of Spades" - > Relying on Go compiler, we only use := syntax for a brand new variable, else card = "test" for reassignment
- Array: Fixed list of things
- Slice: An array that can grow or shrink

2. OO Approach vs Go Approach
- Go is not OO
- Separate files and custom datatypes
- Receiver function: sets up methods on variables.
    1. d = actual copy of the deck (object). By convention, this has to be one/two lettered abbreviation which matches the type but not mandatory
    2. deck = every variable of type deck (class)

3. Slice
- slice_value[startIndexIncluding: endIndexExcluding]
- slice_value[:2] or slice_value[3:]

4. Type Conversion in Go:
- String to byte size: []byte("Hi there") = convert "Hi there" to slice of byte

5. Pointers:
- Go is pass by value language - Go will copy the value of parameter variable and store it in a new variable
- jimPointer := &jim - & - give access to memory address of jim
- variableName *type - this is a type description: It means we are working with a pointer to a person
- *variableName - give value in this memory address. This is an operator - It means we want to manipulate the value pointer is referencing
