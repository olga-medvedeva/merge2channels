# Solution of the problem from the GO training course

### The task was to write a function:

```
func merge2Channels(fn func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int)
```

Description of how the function should work:
- Read one number from each of the two channels in1 and in2, let's call them x1 and x2;
- Сalculate the value of f(x1) + f(x2);
- Pass the received value to the out channel in the same sequence in which the values from in1 and in2 were received.

Also merge2Channels function must be non-blocking, returning immediately and it has a strict time limit of 6 ms.
