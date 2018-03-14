## Memory Sharing

by Michael Pöllath

---
### About me

---

### What inspired this?

Check out the Go-Talk from Bjorn Rabenstein:
https://www.youtube.com/watch?v=1V7eJ0jN8-E

---
### Problem Statement

We want to increase a gobal value by a positiv amount.
It needs to be goroutine safe.
---?code=talk1/sample/interface/interface.go&lang=golang&title=Interface
---?code=talk1/sample/naive_approach/naive_approach.go&lang=golang&title=Naive implementation
---?code=talk1/sample/naive_approach/naive_approach_test.go&lang=golang&title=Benchmark

---
### What is bad about this ?
1. It is not goroutine safe. 

---
### Stuff

- Dispatcher: Manages Data Flow
- Stores: Handle State & Logic
- Views: Render Data via React

---

![Flux Explained](https://facebook.github.io/flux/img/flux-simple-f8-diagram-explained-1300w.png)
