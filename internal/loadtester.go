package internal

import (
	"net/http"
	"sync"
)

// Report armazena os resultados do teste de carga.
type Report struct {
	TotalRequests    int
	Status200        int
	OtherStatusCodes map[int]int
}

// RunStressTest executa o teste de carga conforme os parâmetros.
func RunStressTest(url string, totalRequests, concurrency int) Report {
	var wg sync.WaitGroup
	statusCh := make(chan int, totalRequests)

	// Worker function
	worker := func(requests int) {
		defer wg.Done()
		client := &http.Client{}
		for i := 0; i < requests; i++ {
			resp, err := client.Get(url)
			if err != nil {
				statusCh <- 0 // 0 para erro de requisição
				continue
			}
			statusCh <- resp.StatusCode
			resp.Body.Close()
		}
	}

	// Distribui requests entre workers
	requestsPerWorker := totalRequests / concurrency
	extra := totalRequests % concurrency

	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		reqs := requestsPerWorker
		if i < extra {
			reqs++
		}
		go worker(reqs)
	}

	go func() {
		wg.Wait()
		close(statusCh)
	}()

	report := Report{
		TotalRequests:    totalRequests,
		Status200:        0,
		OtherStatusCodes: make(map[int]int),
	}

	for code := range statusCh {
		if code == 200 {
			report.Status200++
		} else {
			report.OtherStatusCodes[code]++
		}
	}

	return report
}
