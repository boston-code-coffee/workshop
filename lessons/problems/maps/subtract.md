# Map problem #1 - Give list of novel url.

principle: maps

Assume we work at big internet search engine provider. Implement part of a web crawler.  
A web crawler visits all the url and indexes each url.  Indexing a given url is relatively expensive,
so fetch and index a page just one time. 
Write a function called setSubtract() which take two arrays,   
logUrls - a list of url harvested from DNS server logs. We want to fetch the pages from these logs.
indexedUrl  - a list of url that have already been fetched
The function setSubtract returns all the url strings in logUrls that are not indexedUrl.
Why is the function names the way it is?

Part 2 - consider working at "web scale". Assume the arrays are very big. How might this change your solution?

Part 3 - string comparison are expensive at scale, how might that change your solution?

ask your clarifying questions now.




## Assumptions

Each array is a collection of strings.  
Each string is an well formed ascii URL eg "http://www.yahoo.com/main/weather.htm"
logUrls may contain duplicates.
The order of output does not matter.
Memory is important as the input size gets large.

