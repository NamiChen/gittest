package hello

import "rsc.io/quote/v3"

func hello() string {
	return quote.HelloV3()
}

func proverb() string {
	return quote.Concurrency()
}

// func hello() string {

// 	return "hello"
// }
