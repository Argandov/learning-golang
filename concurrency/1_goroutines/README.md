## Goroutines are independent.

!!! main() is a goroutine

worker _goroutine did not finish after main() goroutine exited

```
> worker no_goroutine started
> worker _goroutine started
_ worker no_goroutine done
---- End Program. Duration: 1.001199708s
| Goroutines: 2
```
