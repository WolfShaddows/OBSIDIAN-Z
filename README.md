# 🏴‍☠️ Project OBSIDIAN-Z: The USS Java Siege

**Zero-Trust, Zero-Carbon & Quantum-Secure Architecture for IBM z/OS.**

OBSIDIAN-Z is a polyglot, ultra-low footprint engineering showcase designed to expose the vulnerabilities of legacy Java microservices running on z/OS UNIX System Services (USS), while proposing a mathematically secure, quantum-resistant, and thermodynamically governed alternative.

## 📊 Live Hardware Telemetry (Verified)
The absolute core of OBSIDIAN-Z is sustainability and resilience. Running the full Blue Team defensive enclave (UBI 9 Minimal + Quarkus Native + Rust proxy) yields a verified hardware footprint that obliterates traditional JVM architectures:

| System State | CPU Usage | RAM Usage | Workload Description |
| :--- | :--- | :--- | :--- |
| **Idle (Standby)** | 0.01% | ~4.01 MB | Waiting for network requests. |
| **Cryptographic Stress** | 69.99% | **10.44 MB** | Hashing 1.0 GB of pure data (/dev/urandom) via SHA-256 @ 187 MB/s. |

*Note: A traditional legacy Java application would typically require hundreds of megabytes of heap space to buffer and process equivalent data streams, risking severe Garbage Collection pauses and OutOfMemory vulnerabilities on the mainframe.*

## 🧠 Architectural Thesis
Traditional JVM-based microservices (even JDK 25) on mainframes expand the attack surface (e.g., deserialization RCEs) and consume excessive energy. OBSIDIAN-Z replaces this bloated paradigm with a polyglot architecture utilizing **Go** for high-speed networking, **Rust** for memory-safe RACF/SAF proxying and lattice-based Post-Quantum Cryptography (PQC), and **IBM Granite 4.2** for forensic payload auditing.

## 📂 Ecosystem Modules

### 🔴 1. Red Team Assault (Go)
High-performance compiled Go binaries designed to stress-test the mainframe's perimeter:
- **JWT Forger:** Exploits weak symmetric HMAC signatures to simulate privilege escalation.
- **Payload Injector:** Simulates malicious JSON deserialization attacks aimed at triggering RCE in z/OS USS.

### 🏢 2. Vulnerable Core Legacy (Java 25/Docker)
The 'Mastodon'. A heavy, unoptimized Java container built on eclipse-temurin:25-jdk, demonstrating the high CPU/RAM footprint and security pitfalls of legacy enterprise deployments.

### 🔵 3. Blue Team Obsidian (Rust/Quarkus)
The titanium shield. Features a minimal footprint using Red Hat UBI 9 Minimal and GraalVM native compilation, sitting behind a strictly typed **Rust-based RACF/SAF proxy** that drops unauthorized access before it hits the application logic.

### ⚛️ 4. Cognitive & PQC Auditor (Rust/Go)
The non-negotiable sustainability and security core:
- **PQC Token Shield (Rust):** Seals session tokens using Lattice-based cryptography (CRYSTALS-Kyber model) making them mathematically unbreakable.
- **Granite Forensic Auditor (Go):** A lightweight client interacting with local sovereign AI to classify dropped payloads.
- **Thermal Governor (Go):** A thermodynamic enforcer ensuring the entire ecosystem operates under a strict Net-Zero energy budget.

---
*Developed by Paul Musso (WolfShaddows) - Darkshadows Group | Córdoba, Argentina*
