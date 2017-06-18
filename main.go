package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

var monkey = `                                                                   
 _|      _|    _|_|    _|      _|  _|    _|  _|_|_|_|  _|      _|  
 _|_|  _|_|  _|    _|  _|_|    _|  _|  _|    _|          _|  _|    
 _|  _|  _|  _|    _|  _|  _|  _|  _|_|      _|_|_|        _|      
 _|      _|  _|    _|  _|    _|_|  _|  _|    _|            _|      
 _|      _|    _|_|    _|      _|  _|    _|  _|_|_|_|      _|      
                                                                   
`

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf(monkey)
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
