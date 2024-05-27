# Search problem #1 - Search a sorted list

principle: search

Assume we work at a prominent internet search engine provider. Implement part of a web crawler.
Our web crawler system consists of a set of worker instances. Each worker instance visits many URLs and put them into a sorted array. 
There is one list per worker instance. 
Find the set of URLs which appear in every list.

Write a function called **findCommonUrl()**, which take an 2D array:  
**logUrls** a 2D array of URLs harvested from DNS server logs. Each sub array is the sorted list of URL from one worker.

The function **findCommonUrl()** returns the list of all URL that appear in every list

For example if there were just two workers and the list were 
```
[ "apple.com", "cans.com", "cats.com", "yahoo.com" ]
[ "awesome.com", "cans.com", "cats.com" ] 
```
then the output would be 
```
[ "cans.com", "cats.com" ] 
```




Ask your clarifying Assumptions.

## Assumptions

What to return if string not found
Each array is a collection of strings.    
Each string is an well formed ascii URL eg "http://www.yahoo.com/main/weather.htm"  
**logUrls** does not contain duplicates.  
The order of output does not matter.  
Memory is important as the input size gets large.  

