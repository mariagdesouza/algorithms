package main

/*
Find missing int

Given an input file of 4 Billion non-negative integers

Write an algorithm to generate an integer not in the file


2^32 = 4B  integers
That means we could have only 2 ^31 non negative integers  i.e. about 2B


there are some duplicates

1GB memory = 8Billion bits

1. Create a Bit Vector with 4B bits Array of ints where each int  int represents 32 boolean values

2. Initialize BV to all 0's

3. Scan numbers and ser BV

4. Return the first inde x that has value 0 - tht is the missing number

*/

func main() {

}
