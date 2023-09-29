package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"
)

func extractLevel(s string) (string, error) {
	prefix := "Level:"
	if !strings.Contains(s, prefix) {
		return "", fmt.Errorf("level not found")
	}
	level := strings.TrimSpace(strings.TrimPrefix(s, prefix))
	return level, nil
}

func testPort(serverIP string, port int, wg *sync.WaitGroup) {
	var levelStr string // 1. Déclarez levelStr ici
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", serverIP, port)

	// Tentative de connexion au serveur
	conn, err := net.Dial("tcp", address)
	if err == nil {
		conn.Close()

		// Faire une requête HTTP GET pour /ping
		pingURL := fmt.Sprintf("http://%s:%d/ping", serverIP, port)
		respPing, err := http.Get(pingURL)
		if err == nil {
			defer respPing.Body.Close()
			bodyPing, _ := ioutil.ReadAll(respPing.Body)
			fmt.Printf("Port %d accessible - GET Response for /ping: %s\n", port, bodyPing)
		}

		// Effectuer une requête HTTP POST pour /signup
		hintURL := fmt.Sprintf("http://%s:%d/signup", serverIP, port)
		jsonStr := []byte(`{"User": "Yassine"}`)
		respHint, err := http.Post(hintURL, "application/json", bytes.NewBuffer(jsonStr))
		if err == nil {
			defer respHint.Body.Close()
			bodyHint, _ := ioutil.ReadAll(respHint.Body)
			fmt.Printf("Port %d accessible - POST Response for /signup: %s\n", port, bodyHint)
		}

		// Effectuer une requête HTTP POST pour /check
		hintURL = fmt.Sprintf("http://%s:%d/check", serverIP, port)
		jsonStr = []byte(`{"User": "Yassine"}`)
		respHint, err = http.Post(hintURL, "application/json", bytes.NewBuffer(jsonStr))
		if err == nil {
			defer respHint.Body.Close()
			bodyHint, _ := ioutil.ReadAll(respHint.Body)
			fmt.Printf("Port %d accessible - POST Response for /check: %s\n", port, bodyHint)
		}
		// Effectuer une requête HTTP POST pour /iNeedAHint
		hintURL = fmt.Sprintf("http://%s:%d/iNeedAHint", serverIP, port)
		jsonStr = []byte(`{"User": "Yassine","secret": "75a2a2e61700e659a095e005a5738d82dd9de2aab048faeb9d56efc69f074f9e"}`)
		respHint, err = http.Post(hintURL, "application/json", bytes.NewBuffer(jsonStr))
		if err == nil {
			defer respHint.Body.Close()
			bodyHint, _ := ioutil.ReadAll(respHint.Body)
			fmt.Printf("Port %d accessible - POST Response for /iNeedAHint %s\n", port, bodyHint)
		}

		// Effectuer une requête HTTP POST pour /getUserLevel
		hintURL = fmt.Sprintf("http://%s:%d/getUserLevel", serverIP, port)
		jsonStr = []byte(`{"User": "Yassine","secret": "75a2a2e61700e659a095e005a5738d82dd9de2aab048faeb9d56efc69f074f9e"}`)
		respHint, err = http.Post(hintURL, "application/json", bytes.NewBuffer(jsonStr))
		if err == nil {
			defer respHint.Body.Close()
			bodyHint, _ := ioutil.ReadAll(respHint.Body)
			fmt.Printf("Port %d accessible - POST Response for : oui%soui\n", port, bodyHint)
			levelStr, err := extractLevel(string(bodyHint))
			if err != nil {
				fmt.Printf("Error extracting level: %v\n", err)
			} else {
				fmt.Printf("Extracted Level: %s\n", levelStr)
			}
		}

		// Effectuer une requête HTTP POST pour /getUserPoints
		hintURL = fmt.Sprintf("http://%s:%d/getUserPoints", serverIP, port)
		jsonStr = []byte(`{"User": "Yassine","secret": "75a2a2e61700e659a095e005a5738d82dd9de2aab048faeb9d56efc69f074f9e"}`)
		respHint, err = http.Post(hintURL, "application/json", bytes.NewBuffer(jsonStr))
		if err == nil {
			defer respHint.Body.Close()
			bodyHint, _ := ioutil.ReadAll(respHint.Body)
			fmt.Printf("Port %d accessible - POST Response for /getUserPoints: %s\n", port, bodyHint)
		}

		// Effectuer une requête HTTP POST pour /getUserSecret
		hintURL = fmt.Sprintf("http://%s:%d/getUserSecret", serverIP, port)
		jsonStr = []byte(`{"User": "Yassine","secret": "75a2a2e61700e659a095e005a5738d82dd9de2aab048faeb9d56efc69f074f9e"}`)
		respHint, err = http.Post(hintURL, "application/json", bytes.NewBuffer(jsonStr))
		if err == nil {
			defer respHint.Body.Close()
			bodyHint, _ := ioutil.ReadAll(respHint.Body)
			fmt.Printf("Port %d accessible - POST Response for /getUserSecret %s\n", port, bodyHint)
		}
		// Effectuer une requête HTTP POST pour /enterChallenge
		hintURL = fmt.Sprintf("http://%s:%d/enterChallenge", serverIP, port)
		jsonStr = []byte(`{"User": "Yassine","secret": "75a2a2e61700e659a095e005a5738d82dd9de2aab048faeb9d56efc69f074f9e"}`)
		respHint, err = http.Post(hintURL, "application/json", bytes.NewBuffer(jsonStr))
		if err == nil {
			defer respHint.Body.Close()
			bodyHint, _ := ioutil.ReadAll(respHint.Body)
			fmt.Printf("Port %d accessible - POST Response for /enterChallenge %s\n", port, bodyHint)
		}

		// Effectuer une requête HTTP POST pour /enterChallenge
		hintURL = fmt.Sprintf("http://%s:%d/submitSolution", serverIP, port)
		jsonStr = []byte(fmt.Sprintf(`{"User": "Yassine", "secret": "75a2a2e61700e659a095e005a5738d82dd9de2aab048faeb9d56efc69f074f9e", "level": "0"}`, levelStr))
		respHint, err = http.Post(hintURL, "application/json", bytes.NewBuffer(jsonStr))
		if err == nil {
			defer respHint.Body.Close()
			bodyHint, _ := ioutil.ReadAll(respHint.Body)
			fmt.Printf("Port %d accessible - POST Response for /submitSolution %s\n", port, bodyHint)
		}

	}

}

func main() {
	serverIP := "10.49.122.144"
	minPort := 1000
	maxPort := 65535

	var wg sync.WaitGroup

	for port := minPort; port <= maxPort; port++ {
		wg.Add(1)
		go testPort(serverIP, port, &wg)
	}

	// Attendre que toutes les goroutines se terminent
	wg.Wait()
}
