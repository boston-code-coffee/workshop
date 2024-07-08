Design the order system for a pizza delivery service.  Each order comes from a UI and must be saved for later billing. 
An order consists of:
 user address, 
pizza size (small, medium, large),  
optional toppings(extra cheese, pepperoni, mushroom) 
There are 100 users in the pizza order system. 



# Assumptions and Contraints
In the real world, finding all the assumptions in a problem description is essential in the design process.  A problem owner may only lay part of the problem description.  The engineer needs to uncover the assumptions and constraints. 

 

## Guide Points
The system should be robust to typical network, hardware, and process failure.

The proposed design may reference User management, Login, or billing but is considered "out of scope" for this challenge.  
The candidate does not need to specify these components in their design.

Users are concurrent and distributed.  i.e., two users can send requests to overlapping time periods on a slot, in a data race.

Users are assumed to have a single order.
orders can be any number of days in advance.

Users should be able to create, delete, update, and read their orders.

Responses from the system should be timely enough for a UI and good user experience.

order data should be persistent once a user has successfully made an order.

Post each list of bullet points during the interview, giving the candidate a prompt for each expected section of their design.


# Component Diagram
Draw a system diagram of the system components.

# Guide Points
* UI layer 
* Service layer
* Persistence storage layer
* What DB is chosen - many candidates propose a SQL or in-key-value memory store like Redis - it would be hard to do this with most document stores. 
* Scaling 
* Caching 
* Fail-over 
* Deployment model 
* Logging, monitoring, and alarms

# Estimate size
Estimate the size of the system

## Guide points
* Estimate the load and size of requests per second.
* Concurrent requests from users can happen.
* How will the system scale if the number of users or parking spaces increases?
* How will the system be designed for robustness?  How will it deal with downtime or outages?
* What are the bottlenecks?
* How will the system scale if the number of users or parking spaces increases


# Data model
Select a persistence layer
SQL, document store, column storage,  key-value 
caching
explain tradeoffs and reasons

What DB architecture - SQL, documents store, or even Redis- could work?  SQL is a good default choice.

## Guide points
* Tables
* Columns 
* Values
* backup
* replication
* geolocation
* consistency
* contention

The candidate handles the order as datetime to handle orders further in the future.  Orders are the primary table?
Detail tables could be users or slots with metadata. 
How is concurrency handled? 
How is order overlapping checked? 
Is there some form of DB or table locking if there are several service instances? 

# Design An API 
Select an API model
REST, RPC, message based
explain tradeoffs and reasons

## Guide points
* What are the operations on the data?
* Is there an API?  How will the UI access the data?
* How are errors and failures handled?
* What happens if two users try to reserve a spot at overlapping times?

# Testing
How do you test the system?
What automated tests do you write?  When do you run them?
What manual tests are executed?
What environments do you test with?


## Guide points
* Unit testing, component testing
* Separation of concerns between logic and UI
* End-to-end testing 

# Tooling and delivery

How would you build the app?
How would you deploy it to a production environment?


## Guide points
* CI/CD 
* Deployment pipeline 
* Docker/Kubernetes  

# Monitoring and observation
After deployment, how do you know it's working as expected?

## Guide points
* Logging and metrics 
* Errors 
* Performance
* Observability (e.g., Google Analytics) vs technical
* Page views, users, etc


# Scaling 
It is two years later.  The parking system is very successful.  Scale this design to handle a million users and 100,000 slots.  How would the existing design scale or would it need to change?

## Guide points
* horizontal scaling strategy
* dealing with high contention on a particular slot
* push data to the edge
* write optimizations - ie a persistent Redis layer vs SQL db
* tradeoffs discussion



