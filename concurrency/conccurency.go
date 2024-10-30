package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

/*
	goroutines, the baisc unit of concurrency in Go, which let us manage more than one website check request.

	anonymous functions, which we used to start each of the concurrent processes that check websites.

	channels, to help organize and control the communication between the different process,  allowing us to avoid a race condition.

	the race decector which helped us debug problems with concurrent code.
*/

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChan := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultsChan <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultsChan

		results[r.string] = r.bool
	}
	return results
}
