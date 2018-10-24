package main

import (
    "fmt"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
)

//互斥锁
var lock sync.Mutex
var rwlock sync.RWMutex

// 读写锁

func testMap() {

    var a map[int]int = map[int]int{
        1: 2,
        2: 3,
        3: 4,
        4: 5,
        5: 6,
        6: 7,
    }

    for i := 0; i < 10; i++ {
        go func(b map[int]int) {

            lock.Lock()
            b[3] = rand.Intn(100)
            lock.Unlock()
        }(a)
    }

    lock.Lock()
    fmt.Println(a)
    lock.Unlock()
}

func RWlock() {

    var a map[int]int = map[int]int{
        1: 2,
        2: 3,
        3: 4,
        4: 5,
        5: 6,
        6: 7,
    }
    var count int32
    // var x int
    //原子操作
    for i := 0; i < 2; i++ {
        go func(b map[int]int) {
            rwlock.Lock()
            // lock.Lock()
            b[3] = rand.Intn(100)
            time.Sleep(time.Millisecond * 10)
            rwlock.Unlock()
            // lock.Unlock()

        }(a)
    }

    for i := 0; i < 100; i++ {
        go func(b map[int]int) {
            for {
                rwlock.RLock()
                // lock.Lock()
                time.Sleep(time.Millisecond)
                // x := b[3]
                // fmt.Println(b[3])
                // lock.Unlock()
                rwlock.RUnlock()
                atomic.AddInt32(&count, 1)
            }
        }(a)
    }
    time.Sleep(time.Second * 3)
    fmt.Println(atomic.LoadInt32(&count))
    time.Sleep(time.Second)
    // fmt.Println(x)
}

func main() {

    // testMap()
    RWlock()
}
