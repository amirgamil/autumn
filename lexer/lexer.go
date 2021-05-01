package lexer

type Lexer struct {
	input        string
	position     int  //represents the position of the current character in the input we've processed
	readPosition int  //represents the positions of the character we are READING in the input - need this to be able to peek further down
	char         byte //represents the current character we're processing
}

func New(input string) {
	l := Lexer{input: input}
	return &l
}

//go pointer reciever - a method that operations on the pointer and modify it. Similar to a method in a class
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}
	//move processed position to old readPosition
	l.position = l.readPosition
	//update read position
	l.readPosition += 1

}
