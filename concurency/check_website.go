package concurency

type result struct {
	string
	bool
}

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)} // goroutine writes the struct to the channel
		}(url)

	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel // statement called: receive expression
		results[r.string] = r.bool
	}

	// time.Sleep(2 * time.Second)

	return results
}
