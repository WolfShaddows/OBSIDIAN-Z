// Explicación: Cliente en Go ultra-rápido que se comunica con una instancia local de Ollama corriendo el modelo IBM Granite.
// Objetivo: Actuar como un auditor forense cognitivo. Cuando el escudo en Rust detecta una anomalía (ej. intento de RCE), este microservicio envía el payload a Granite para su análisis semántico y clasificación de la amenaza en tiempo real, operando bajo un estricto presupuesto energético.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

func main() {
	fmt.Println("[GRANITE AUDITOR] Inicializando motor de inferencia local...")
	
	// Simulamos la recepción de un payload malicioso bloqueado por el Blue Team (enviado por el inyector)
	suspectPayload := `{"@type": "com.sun.rowset.JdbcRowSetImpl", "dataSourceName": "ldap://..."}`
	
	prompt := fmt.Sprintf("Analiza este payload interceptado en un entorno z/OS e identifica si es un intento de ataque de deserializacion maliciosa: %s", suspectPayload)

	reqBody := OllamaRequest{
		Model:  "granite-code:3b", // Modelo ultra-ligero para mantener el consumo eléctrico por el piso
		Prompt: prompt,
		Stream: false,
	}

	jsonData, _ := json.Marshal(reqBody)

	// Asumiendo que el motor local de Granite corre en localhost:11434
	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("[GRANITE AUDITOR] Error de conexion con motor cognitivo:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("[GRANITE AUDITOR] Veredicto Forense de IA (Respuesta en crudo):")
	fmt.Println(string(body))
}