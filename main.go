package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/alexandreLITHAUD/go-calculator/calculator"
)

func main() {
	var (
		port       = flag.Int("port", 8080, "Server port")
		serverMode = flag.Bool("server", false, "Run as HTTP server")
		a          = flag.Float64("a", 0, "First number")
		b          = flag.Float64("b", 0, "Second number")
		op         = flag.String("op", "add", "Operation: add, sub, mul, div")
	)
	flag.Parse()

	fmt.Println("🧮 Go Calculator with DevBox")
	fmt.Println("============================")

	if *serverMode {
		runServer(*port)
	} else {
		runCalculator(*a, *b, *op)
	}
}

func runCalculator(a, b float64, operation string) {
	calc := calculator.New()

	var result float64
	var err error

	switch operation {
	case "add":
		result = calc.Add(a, b)
		fmt.Printf("%.2f + %.2f = %.2f\n", a, b, result)
	case "sub":
		result = calc.Subtract(a, b)
		fmt.Printf("%.2f - %.2f = %.2f\n", a, b, result)
	case "mul":
		result = calc.Multiply(a, b)
		fmt.Printf("%.2f * %.2f = %.2f\n", a, b, result)
	case "div":
		result, err = calc.Divide(a, b)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%.2f / %.2f = %.2f\n", a, b, result)
	default:
		fmt.Printf("❌ Unknown operation: %s\n", operation)
		fmt.Println("Available operations: add, sub, mul, div")
		os.Exit(1)
	}
}

func runServer(port int) {
	fmt.Printf("🌐 Starting HTTP server on port %d...\n", port)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/calc", calcHandler)
	http.HandleFunc("/health", healthHandler)

	log.Printf("Server running at http://localhost:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>Go Calculator</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 600px; margin: 50px auto; padding: 20px; }
        .calc { background: #f5f5f5; padding: 20px; border-radius: 8px; }
        input, select { padding: 8px; margin: 5px; }
        button { padding: 10px 20px; background: #007bff; color: white; border: none; border-radius: 4px; cursor: pointer; }
        button:hover { background: #0056b3; }
    </style>
</head>
<body>
    <h1>🧮 Go Calculator with DevBox</h1>
    <div class="calc">
        <h2>Calculator</h2>
        <form action="/calc" method="get">
            <input type="number" step="any" name="a" placeholder="First number" required>
            <select name="op">
                <option value="add">+</option>
                <option value="sub">-</option>
                <option value="mul">×</option>
                <option value="div">÷</option>
            </select>
            <input type="number" step="any" name="b" placeholder="Second number" required>
            <button type="submit">Calculate</button>
        </form>
    </div>
    
    <h2>API Endpoints</h2>
    <ul>
        <li><code>GET /calc?a=10&b=5&op=add</code> - Calculate</li>
        <li><code>GET /health</code> - Health check</li>
    </ul>
</body>
</html>`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")
	op := r.URL.Query().Get("op")

	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		http.Error(w, "Invalid parameter 'a'", http.StatusBadRequest)
		return
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		http.Error(w, "Invalid parameter 'b'", http.StatusBadRequest)
		return
	}

	calc := calculator.New()
	var result float64

	switch op {
	case "add":
		result = calc.Add(a, b)
	case "sub":
		result = calc.Subtract(a, b)
	case "mul":
		result = calc.Multiply(a, b)
	case "div":
		result, err = calc.Divide(a, b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"a": %g, "b": %g, "operation": "%s", "result": %g}`, a, b, op, result)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status": "ok", "service": "go-calculator"}`)
}
