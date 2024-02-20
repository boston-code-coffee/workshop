# Map problem #1 - Return list of novel strings from two list of strings.

principle: maps

Assume we work at a prominent internet search engine provider. Implement part of a web crawler.
A web crawler visits all the URLs and indexes each URL. Indexing a given URL is relatively expensive, 
so fetch and index a page just one time. Write a function called **setSubtract()**, which take two arrays,
**logUrls** a list of URLs harvested from DNS server logs. We want to fetch the pages from these logs.   
**indexedUrl**  a list of URLs that have already been fetched.   

The function **setSubtract()** returns all the URL strings in **logUrls** that are not **indexedUrl**. 
Why is the function name the way it is?

Part 2 - consider working at "web scale". Assume the arrays are very big. How might this change your solution?

Part 3 - string comparisons are expensive at scale; how might that change your solution?

Ask your clarifying questions now.



## Assumptions

Each array is a collection of strings.  
Each string is an well formed ascii URL eg "http://www.yahoo.com/main/weather.htm"
**logUrls** may contain duplicates.
The order of output does not matter.
Memory is important as the input size gets large.

