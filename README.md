# Explanations about GO

 - When you wish export something in GO you should name your function with first letter in UpperCase and add a comment up the function
 - You can return two data in one return

# Concurrency in GO

 - Concurrency isn't the same than parallelism
 - To create concurrents funcions you should use the reserved word 'go' before function call, and it's called as goroutine
 - Channels are pipes that connect concurrent goroutine
 	- We use the method make() to create a channel
 	- We use the method close() to finish the channel
