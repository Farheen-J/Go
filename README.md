# Go - The Complete Developer's Guide (Golang, udemy)

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
- A slice has 3 attributes internally: ptr to head(head of the array), capacity and length

4. Type Conversion in Go:
- String to byte size: []byte("Hi there") = convert "Hi there" to slice of byte

5. Pointers:
- Go is pass by value language - Go will copy the value of parameter variable and store it in a new variable
- jimPointer := &jim - & - give access to memory address of jim
- variableName *type - this is a type description: It means we are working with a pointer to a person
- *variableName - give value in this memory address. This is an operator - It means we want to manipulate the value pointer is referencing
- *** Whenever you pass an integer, float, string, or struct into a function, it is pass by value but not applcable for slice ***

6. Gotchas with pointers
- In struct you need pointers but in a slice, you do not because slice has 3 attributes internally: ptr to head(head of the array), capacity and length. So when value of the slice is copied(pass by value), the pointer is copied as well.
- DS that behave like slices: maps, channels, pointers, functions
- DS who don't: int, float, string, bool, structs

7. Map:
- All Keys and all Values should be of same type

8. Struct vs Map:
- ![Alt text](images/image.png)
