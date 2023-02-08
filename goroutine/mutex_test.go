package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// untuk menghandle kasus race condition maka kita bisa menggunakan Mutex (Mutual Exclution)
// dengan membuat seperti antrian saat sebelum kita memutasi shared variable di goroutine
func TestMutex(t *testing.T) {
	number := 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				// hanya satu goroutine yang dapat melakukan lock & unlock-
				// sehingga goroutine yang lain harus menunggu-
				// dengan begitu goroutine secara bergantian memutasi variable "number"
				mutex.Lock()
				number += 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Number: ", number)

}

type BankAccount struct {
	// Read Write Mutex digunakan untuk melocking proses read/write
	// sehingga kita tidak perlu membuat 2 Mutex
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()

	}

	time.Sleep(3 * time.Second)
	fmt.Println("Final Balance: ", account.Balance)
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(userA *UserBalance, userB *UserBalance, amount int) {
	userA.Lock()
	fmt.Println("Lock UserA", userA.Name)
	userA.Change(-amount)

	time.Sleep(1 * time.Second)

	userB.Lock()
	fmt.Println("Lock UserB", userB.Name)
	userB.Change(amount)

	time.Sleep(1 * time.Second)

	userA.Unlock()
	userB.Unlock()
}

func TestDeadLock(t *testing.T) {
	userNanda := UserBalance{Name: "Nanda", Balance: 1000000}
	userAfif := UserBalance{Name: "Afif", Balance: 1000000}

	// ketika "userAfif" ingin dilock oleh goroutine ke-1 maka akan terjadi deadlock-
	// dikarenakan "userAfif" sudah dilock duluan oleh goroutine yang ke-2 -
	// (karena dimasukan sebagai arguemen pertama, sehingga diesekusi terlebih dahulu).
	// maka si goroutine ke-1 akan menunggu terus menerus hingga "userAfif" diunlock,
	// yg nyatanya malah terjadi deadlock
	// karna menunggu terus menerus dan tak pernah selesai.
	// sebaliknya ini berlaku juga untuk "userNanda" di goroutine ke-1 sudah dilock namun goroutine ke-2 mau me-lock juga-
	// sehingga terjadi saling tunggu
	go Transfer(&userNanda, &userAfif, 100000)
	go Transfer(&userAfif, &userNanda, 200000)

	time.Sleep(5 * time.Second)

	fmt.Println("User A: ", userNanda.Name, " Balance: ", userNanda.Balance)
	fmt.Println("User B: ", userAfif.Name, " Balance: ", userAfif.Balance)
}
