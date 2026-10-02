package main
/*Given two English words of the same length, say, "HEAD" and "TAIL", the goal is to code for 
# coming up with  a sequence  of valid English words, starting with "HEAD", and ending with 
# "TAIL", such that each word is formed by changing a single letter of the previous word. 
# In this example, one solution is:  Head -> Heal ->Teal ->Tell->Tall->Tail
*/
//HEADTAIL 

/*
 word_ladder(words, start, end):
    """Return a word ladder (a list of words each of which differs from
    the last by one letter) linking start and end, using the given
    collection of words. Raise NotFound if there is no ladder.

    >>> words = 'card care cold cord core ward warm'.split()
    >>> ' '.join(word_ladder(words, 'cold', 'warm'))
	'cold cord card ward warm'
	
*/

func main(){
   wordlist := map[string]struct{}{
	"HEAD", "HEAT" , "TEAL","TELL","TALL","TAIL",
   }
}


//HEADTAIL

func findWords(str1, str2 string, wordlist map[string]struct{}){
    
    letterSource := str2
    ll := 
    startStr := str1 // HEAD ->HEAT -> TEAL ->TELL->TALL->TAIL
    
    // replace  one character in startstr
    // check if it is a valid word  (including rearranging characters)
    // track invalid words 
    // check that it is equal to the destination word 
    invalidWords := make(map[string]struct{}{})
    for i, c := range str2{
        str1[len(a)-1] = c 
        checkValidWord(str1, invalidWords)
    }
    
    
}

func checkValidWord(str1 string, invalidWords map[string]struct{}{}) bool{
    
    
    
    
}

