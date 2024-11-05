package main

import (
	"sort"
	"strings"
)

// clean the string for palindrome
// reverse the string and compare the last and first letter and continue ot move
//
//
//



func Panlindrome(text string) bool  {
	
	test := strings.ToLower(text)
	n := len(test)

	for	i := 0; i < n/2; i++{
		if test[i] != test[n-1-i]{
			return false
		}
		    
	}
   return true;
}


func AnagramSort(text1, text2 string) bool {
   // Check if the lengths of the strings are different
    if len(text1) != len(text2) {
        return false
    }
    
  s1 := []byte(text1)
  s2 := []byte(text2)

  sort.SliceStable(s1, func(i, j int) bool {
    return s1[i] < s1[j]
})
 sort.SliceStable(s2, func(i, j int) bool {
    return s2[i] < s2[j]
})

sortedText1 := string(s1)
sortedText2 := string(s2)

return sortedText1 == sortedText2


}

func AnagramCountFreq (text1, text2 string) bool {
	 // Check if the lengths of the strings are different
    if len(text1) != len(text2) {
        return false
    }
	map1 := map[rune] int {}
	map2 := map[rune] int {}

     for _, char := range text1 {
		map1[char]++
       
		}


	   for _, char := range text2 {
		map2[char]++
       
		}

		 // Compare the maps
    for char, count := range map1 {
        if map2[char] != count {
            return false
        }
    }


	 return true

}