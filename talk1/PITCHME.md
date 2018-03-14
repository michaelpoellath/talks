## Memory Sharing


---
### About me
![profile](https://avatars2.githubusercontent.com/u/24293821?s=460&v=4)
Michael Pöllath

Software Developer @Newstore
https://github.com/michaelpoellath

---

### What inspired this?

Check out the Go-Talk from Björn Rabenstein:
https://www.youtube.com/watch?v=1V7eJ0jN8-E

---
### Problem Statement

1. We want to increase a gobal value by a positiv amount.
2. It needs to be goroutine safe.

---?code=talk1/sample/interface/interface.go&lang=golang&title=Interface
---?code=talk1/sample/naive/naive_approach.go&lang=golang&title=Naive implementation
---?code=talk1/sample/naive/bench.go&lang=golang&title=Benchmark

--- 
### How to Benchmark

go test -bench=Counter -cpu=1,4,16 #-race

---
### Results Naive
![Results Naive](https://i.imgur.com/hJjyp6r.png)

---
### What is bad about this ?

1. It is not go routine safe.


---
### Lets try mutex

---?code=talk1/sample/go_routine_mutex/go_routine_mutex.go&lang=golang&title=Mutex implemenation
---?code=talk1/sample/go_routine_mutex/bench.go&lang=golang&title=Mutex benchmark

---
### Results Mutex
![Results Mutex](https://i.imgur.com/U9Mt4D0.png)

---
### Why is it so much slower?

Lock contention!

---
### Maybe channels?

---?code=talk1/sample/go_channel/go_channel.go&lang=golang&title=Channel implemenation
---?code=talk1/sample/go_channel/bench.go&lang=golang&title=Channel benchmark

---
### Results Channel
![Results Channel](https://i.imgur.com/xoXB1ee.png)

---
### This is expensive!
1. We do not lock - good
2. CPU overhead - bad 

---
### Atomic
1. No locking through the Dev
3. Not many operations supported
https://golang.org/pkg/sync/atomic/


---?code=talk1/sample/atomic/atomic.go&lang=golang&title=Atomic implemenation
---?code=talk1/sample/atomic/bench.go&lang=golang&title=Atomic benchmark

---
### Results Atomic
![Results Atomic](https://i.imgur.com/yIOwlAM.jpg)

---
### This is good!
1. ~ 45 times faster than channels
2. ~ 9 times faster than locking
3. ~ 3 times slower than the naive approach :D

---
### Learnings
1. Solution is depending on your use-case
2. Benchmarking is key for achieve good performance
3. Go channels are not a Swiss army knife

---
### Thank you!
