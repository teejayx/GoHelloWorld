package integers 

import "testing"

//%d prints interger 
//%q prints out string 

func TestAdder(t *testing.T){
    sum := Add(2, 2)
	expected := 4

	if(sum != expected){
  t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}