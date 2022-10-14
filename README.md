# restaurant-problem

## Question 1
The logs are stored as json in `.txt` file 
  
	
```
food_id := make(map[string](map[string]bool))
```
Stores key as foodmenu_id and key value pair of eater_id and boolean 

	
```
freq := make(map[string]int)
```
counts the frequency of food_id ordered

To make sure there is no repeated value we check if the key-value pair exists or not

```
if food_id[log.FoodmenuID][log.EaterID] == true {
				return []string{}, errors.New("similar entries found")
			}
```


## Question 2
The following piece of code is an implementation of worker pool concurrency pattern 

In this piece of code 
```
package main

import "fmt"

func main() {

  cnp := make(chan func(), 10)
  for i := 0; i < 4; i++ {
    go func() {
       for f := range cnp {
         f()
        }
      }()
    }

  cnp <- func() {
    fmt.Println("HERE1")
  }

  fmt.Println("Hello")
}
```
We are creating a channel with buffer size of 10 (`cnp := make(chan func(), 10)`) 

and creating 4 worker(go routines) to distribute the tasks

This code adds tasks to the cnp channel
```
  cnp <- func() {
    fmt.Println("HERE1")
  }
```
Since we created a buffered channel, it will ensure that only ten tasks are run in a batch after the buffer is filled.

In this case "Here1" wont be printed as there is just one task added.

### Use case
When we spawn too many goroutine, our machine will quickly run out of memory and the CPU will keep processing the task until it reach the limit. By using limited pool of workers and keep the task on the queue, we can reduce the burst of CPU and memory since the task will wait on the queue until the the worker pull the task.

This approach is useful to process large tasks into batches. Collecting individual work items and distribute them amongst workers for concurrent processing.
In this case we have 4 worker which are responsible for 10 tasks per batch.