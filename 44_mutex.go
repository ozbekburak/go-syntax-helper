package main

// sync.mutex
// goroutineler arasında haberleşmek için kanalların işlevini gördük
// Ancak, haberleşmeye ihtiyacımız olmadığı durumda ne olacak?
// İstediğimiz, bir değişkene sadece bir goroutinein erişmesi ise ne yapacağız? (çakışmalardan kaçınmak için)
// bu konseptin ismine mutual exclusion deniyor
// Go'nun standart kütüphanesi sync.Mutex ile sağlıyor ve bunun "Lock(), Unlock()" olmak üzere iki methodu var

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc methodu argüman olarak aldığı key'i bir arttırıyor.
func (c *SafeCounter) Inc(key string) {
	// Lock() methodunu kullanıyor, çünkü çalıştığında ona sadece bir goroutinein erişmesini istiyor
	c.mux.Lock()

	c.v[key]++
	c.mux.Unlock() // başka bir goroutine eriştiğinde çalışması için unlock() yapıyoruz
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
