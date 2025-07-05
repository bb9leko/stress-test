package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/bb9leko/stress-test/internal"
)

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 1, "Número total de requests")
	concurrency := flag.Int("concurrency", 1, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" || *requests <= 0 || *concurrency <= 0 {
		fmt.Println("Uso: --url=<url> --requests=<n> --concurrency=<n>")
		os.Exit(1)
	}

	start := time.Now()
	report := internal.RunStressTest(*url, *requests, *concurrency)
	elapsed := time.Since(start)

	fmt.Println("===== Relatório de Teste de Carga =====")
	fmt.Printf("Tempo total: %v\n", elapsed)
	fmt.Printf("Total de requests: %d\n", report.TotalRequests)
	fmt.Printf("HTTP 200: %d\n", report.Status200)
	fmt.Println("Distribuição de outros códigos de status:")
	for code, count := range report.OtherStatusCodes {
		fmt.Printf("  %d: %d\n", code, count)
	}
}
