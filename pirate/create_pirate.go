package pirate

import (
	"errors"
	
)

type Pirate struct {
	Name  string
	Prime string
	Img   string
}

func New(name, prime, img string) (Pirate, error) {
	var err error
	// fmt.Println(name,prime,img)
	list_str := []byte(name)
	list_int := []byte(prime)
	// fmt.Println(string(list_str))
	// CHECK MAJ
	for _, v := range list_str {
		if v != 32 && v != 46 &&    v < 'A' || v > 'Z' {
			err = errors.New("error uppercase")
			
		}
	}
	// CHECK NUMBER
	for _, v := range list_int {
		if v < '0' || v > '9' {

			err = errors.New("error int")
		}
	}
	// CREATE PIRATE
	new_pirate := Pirate{Name: name, Prime: prime, Img: img}
	return new_pirate, err
}
