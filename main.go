package main

import "fmt"

type course struct {
	name string
}

func (c course) Print() {
	fmt.Printf("curso %+v\n", c)
}

//declaracion de alias
// type myAlias = course

//definicion de tipo
//obtengo los campos de course , pero no los metodos
// type newCourse course


// puedo crear sobre tipos predeclarado
type newBool bool


func (b newBool) String() string {
  if b {
    return "VERDADERO"
  }

  return "FALSO"
}

func main() {
	// c := newCourse{name: "Go"}
  
	// c.Print()

  var b newBool = true
  
	// fmt.Printf("El tipo es %T\n y su valor es %v", c,c)
	fmt.Printf(b.String())
}
