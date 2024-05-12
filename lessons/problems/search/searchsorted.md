# Search problem #1 - Search a sorted list

principle: search

Assume we work at a prominent internet search engine provider. Implement part of a web crawler.
A web crawler visits has visited many URLs and put them into a sorted array. 

Write a function called **findUrl()**, which take an array:  
**logUrls** a sorted array of URLs harvested from DNS server logs. 
**target** a URL to find.

The function **findUrl()** returns the string **target** if its present in the sorted list. 

Part 2 - Implement **findUrls()** which takes two arrays
**logUrls** a sorted array of URLs harvested from DNS server logs. 
**targets** an array of URL to find.


How will you implement the more general findUrls()_


Ask your clarifying Assumptions.

## Assumptions

What to return if string not found
Each array is a collection of strings.    
Each string is an well formed ascii URL eg "http://www.yahoo.com/main/weather.htm"  
**logUrls** does not contain duplicates.  
**targets** does not contain duplicates.  
The order of output does not matter.  
Memory is important as the input size gets large.  

