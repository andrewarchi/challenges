package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func fib(num chan *big.Int) {
	a := new(big.Int).SetUint64(0)
	b := new(big.Int).SetUint64(1)
	for {
		num <- a
		a, b = b, new(big.Int).Add(a, b)
	}
}

func fibFiller() func(int64) []*big.Int {
	var nums []*big.Int
	cur := make(chan *big.Int)
	go fib(cur)
	return func(n int64) []*big.Int {
		for i := int64(len(nums)); i < n; i++ {
			nums = append(nums, <-cur)
		}
		return nums[:n]
	}
}

type fibResponse struct {
	Nums []*big.Int `json:"nums"`
	Err  error      `json:"error"`
}

func serveFib() func(http.ResponseWriter, *http.Request, httprouter.Params) {
	fillFib := fibFiller()
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		nParam := p.ByName("n")
		n, err := strconv.ParseInt(nParam, 10, 64)
		if err != nil {
			writeResponse(w, nil, fmt.Errorf("n must be integer, got %s", nParam))
			return
		}
		if n < 0 {
			writeResponse(w, nil, fmt.Errorf("n must be positive, got %d", n))
			return
		}
		writeResponse(w, fillFib(n), nil)
	}
}

func writeResponse(w http.ResponseWriter, nums []*big.Int, err error) {
	b, err := json.Marshal(&fibResponse{nums, err})
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(b)
}

func main() {
	router := httprouter.New()
	router.GET("/api/fibonacci/:n", serveFib())
	log.Fatal(http.ListenAndServe(":8080", router))
}
