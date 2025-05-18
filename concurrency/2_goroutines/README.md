## Goroutines are independent.

(!!!) main() is itself a goroutine.

worker 4 ended even after main() returns.

```
---- Start Program at 2025-05-18 12:34:26.999327 -0600 CST m=+0.000086334
| CPUs: 8
| Goroutines: 1
----------------------------
> worker 1 started
> worker 0 started
> worker 2 started
> worker 3 started
> worker 4 started
_ worker 3 done
_ worker 2 done
---- End Program. Duration: 1.001126167s
| Goroutines: 5
_ worker 4 done

```

worker 999 even ran before previous workers finished. 

```
---- Start Program at 2025-05-18 12:34:30.93309 -0600 CST m=+0.000053460
> worker 0 started
> worker 1 started
> worker 2 started
> worker 3 started
> worker 4 started
_ worker 4 done
_ worker 2 done
_ worker 3 done
_ worker 1 done
_ worker 0 done
_ instantWorker 999 done
---- End Program. Duration: 1.001194167s
| Goroutines: 1
```

Each go worker(i) call immediately queues a new goroutine with the Go scheduler.

Those goroutines run “in the background” at times the scheduler picks, not necessarily in the order they were launched.

- !! Program flow is unpredictable without sync.
- The main goroutine does not wait for other goroutines to finish.

After time.Sleep(time.Second), main resumes and even launches worker(999) but continues immediately to instantWorker(999) and the final prints.

Race between main and background work.

You saw > instantWorker 999 started and the “End Program” before > worker 999 started because the scheduler hadn’t yet started that goroutine.

By the time main prints the goroutine count, only the main and the just‐started worker(999) remain alive.

Main exit kills all goroutines.

As soon as main() returns (after its final Println), the Go runtime shuts down the process and stops all goroutines immediately—so any unfinished worker logs never appear.

Key takeaway: launching a goroutine is fire‐and‐forget unless you explicitly synchronize. To make your programs deterministic and ensure no work is lost, use synchronization primitives (e.g., sync.WaitGroup, channels) so that main only exits once all necessary goroutines have completed.
