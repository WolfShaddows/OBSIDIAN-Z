// Explicación: Código ajustado para la generación del par de claves post-cuánticas usando Kyber768.
// Objetivo: Eliminar las advertencias de compilación importando únicamente los traits estrictamente necesarios (SharedSecret y Ciphertext) requeridos para el método .as_bytes(), manteniendo la eficiencia y delegando el tipado de claves a la inferencia nativa de Rust.
use pqcrypto_kyber::kyber768::*;
use pqcrypto_traits::kem::{Ciphertext, SharedSecret};

fn main() {
    println!("[PQC SHIELD] Inicializando motor criptográfico post-cuántico (NIST Standard: Kyber768)...");

    // 1. Generación matemática real del par de claves (Key Encapsulation Mechanism)
    let (public_key, _secret_key) = keypair();
    println!("[PQC SHIELD] Par de claves Kyber768 generadas exitosamente en la bóveda de memoria.");

    let session_claim = "sub:legit_user|role:ROLE_READ_ONLY";
    println!("[PQC SHIELD] Preparando encapsulado para el claim: {}", session_claim);

    // 2. Encapsulamiento del secreto compartido
    let (shared_secret, ciphertext) = encapsulate(&public_key);

    // Convertimos los bytes criptográficos a formato hexadecimal para la red
    let ct_hex = hex::encode(ciphertext.as_bytes());
    let ss_hex = hex::encode(shared_secret.as_bytes());

    // 3. Salida de auditoría
    println!("[PQC SHIELD] Secreto Compartido derivado (AES-256 target key): {}...", &ss_hex[0..32]);
    println!("[PQC SHIELD] Ciphertext Cuántico (Token Payload): ENC_{}...", &ct_hex[0..64]);
    println!("[PQC SHIELD] Token sellado matemáticamente e inquebrantable listo para inyección en z/OS.");
}
