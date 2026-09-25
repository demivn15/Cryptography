# 🔑 Cryptography

## Lab 1 - DES Algorithm Implementation

Lab 1 contains the technical implementation, mode analysis, and parallel cryptanalysis of the Data Encryption Standard (DES) cipher written in Go

### Environment setup

- Operating System: Linux Debian 13
- Programming Language: Go 1.24.4
- Execution Environment: 4 Physical CPU Cores

### Implementation details

The repository consists of modular packages handling core cryptographic primitives, operational modes, and parallel search routines:

- **`utils`**: Auxiliary functions handling bitwise manipulation, ASCII/binary conversions, bit rotations, and user I/O.
- **`myDES`**: The full Feistel cipher pipeline, including key generation, S-Box substitution, initial/final permutations, and the core $F$-function
- **`modes`**: Implementations of Electronic Codebook (ECB) and Cipher Block Chaining (CBC)] modes along with PKCS#7 padding.
- **`bruteforce`**: A concurrent key-recovery engine utilizing goroutines and worker pools.

#### How to run:

- Clone this repo: `git clone https://github.com/demivn15/Cryptography`.
- Install Go on your system: `sudo apt install golang-go`.
- Run the interactive CLI: `go run .`.
- Execute unit and verification tests against standard vectors: `go test -v ./...`.
- Run mode experiments and parallel benchmarks: `go run main.go --benchmark`.

---

### DES Algorithm Overview

The Data Encryption Standard Algorithm is a symmetric key algorithm, that is, an algorithm that uses the same cryptographic key for both encryption and decryption. This enctyption algorithm takes plaintext of 64-bit in size as input, and returns the corresponding ciphertext. The encryption is done in a process of 16 rounds, each of which uses a different subkey of 48 bits long. The whole process relies on predefined permutation tables that are applied to the plaintext and the key.

#### Step-by-step:

- The algorithm takes the plaintext (or a block of it) of 64 bits long as input and performs an **initial permutation**;
- it does the same thing for the cryptographic key (which initially has a length of 64 bits) and performs a **permuted choice one** that shortens the length of the key down to 56 bits;
- the algorithm continues by performing **16 rounds** of processing where the permuted plaintext and subkeys of the shortened key are used;
- at the end, the algorithm performs a **final permutation** which outputs the ciphertext.

##### Rounds breakdown:

- The first round takes the permuted text as input and splits it into two halves (32 bits each);
- the left-side half (L_n) is xor'd with the output of a function called the **Fesitel function**, while the right-side half (R_n) is the input of said function together with a **transformed subkey**;
- the ouput of the xor operation becomes the right half of the processed text, while R_n becomes the left half of it;
- this new formed string becomes the input for the next round.

##### Feistel function breakdown:

- The Feistel function takes R_n and performs an **expansion** over it;
- then, the output of the expansion is xor'd with the generated subkey for the current round;
- the xor output is processed to accomplish *nonlinearity* by shortening the input from 48 bits down to 32 bits through a **mapping using eight tables called S-Boxes**;
- then, the result goes through another **permutation** and becomes the output of the function.

> The purpose of the **expansion** is to expand R_n from 32 to 48 bits so the xor operation can be performed with the 48-bit round key.

> The purpose of performing a final **permutation** is contributing to diffusion.

##### S-Boxes:

- This process receives an input of 48 bits and divides it into eight chunks, each of 6 bits long. There is an S-Box for each chunk that allows to go from 6 bits long down to 4 bits;
- the mapping takes both the first and last bit of each chunk and the resulting number from concatenating them becomes the row of the box;
- the 4 bits left in the middle become the column of the box;
- then, the value at that position becomes the output of processing the chunk. This value is represented with 4 bits;
- all the results from processing each chunk are concatenated together to form a 32-bit long word.

##### Transformation of subkeys breakdown:

- In order to generate the subkeys for each round the resulting key from performing **permuted choice one** is split into halves (28 bits each);
- a left rotation is performed. The number of positions shifted depends on the current round (one position is shifted for rounds 1, 2, 9, 16, and two are shifted for any other round);
- the resulting 56-bits key is performed a **round key permutation** on. This derives a 48-bit long subkey, the one used as input for the Feistel function.

#### Decryption

The algorithm used is the same. The only difference is the order of the round keys.

#### ECB vs. CBC Operational Modes
- **Electronic Codebook (ECB)**: Encrypts 64-bit blocks independently ($C_i = E_K(P_i)$). Deterministic mapping causes identical input blocks to produce identical output blocks ($P_i = P_j \implies C_i = C_j$), leaking statistical and visual patterns in structured plaintexts.
- **Cipher Block Chaining (CBC)**: XORs each plaintext block with the preceding ciphertext block ($C_i = E_K(P_i \oplus C_{i-1})$) using a 64-bit Initialization Vector ($IV$) for the initial block. This feedback loop breaks input repetition, mapping identical plaintext blocks to pseudorandom, distinct ciphertexts.

#### Error Propagation Mechanics
Introducing a single-bit flip into Bit 5 of Block 2 ($C_2$) yields distinct recovery profiles upon decryption:
- **ECB Mode**: Corrupts Block 2 completely ($\approx 50\%$ bit error rate due to the Feistel avalanche effect), while surrounding blocks ($P_1, P_3$) recover with 0% error.
- **CBC Mode**: Corrupts Block 2 completely ($57.81\%$ bit error rate) and propagates exactly 1 bit flip into $P_3$ ($1.56\%$ error rate) because $P_3 = D_K(C_3) \oplus C_2$.

#### Parallel Brute-Force Key Recovery & Extrapolation
To evaluate 56-bit key entropy:
- **Worker Pool Scaling**: Evaluated over a $2^{20}$ search space across concurrent workers on 4 CPU cores. Throughput scales from $20,719.54\text{ keys/s}$ (1 worker) to $53,145.75\text{ keys/s}$ (4 workers, $3.85\times$ speedup).
- **Early Termination Dynamics**: Larger worker pools split search chunks, encountering target keys early in chunk intervals and yielding super-linear speedups ($11.20\times$ speedup with 8 workers).
- **Full $2^{56}$ Key Space Extrapolation**: At a peak throughput rate of $R = 500,000\text{ keys/s}$ on a single workstation, average recovery ($2^{55}$ keys) takes $\approx 2,283.37\text{ years}$. However, distributing the search across 10,000 GPU nodes operating at $10^{11}\text{ keys/s}$ reduces search time to $\approx 4.17\text{ days}$, rendering 56-bit keys cryptographically obsolete.

---

### DES Algorithm Architecture
The Data Encryption Standard is a symmetric Feistel block cipher that maps 64-bit plaintext blocks to 64-bit ciphertext blocks across 16 processing rounds.

## Lab 3:  AES  Implementation

This project provides a modular, pure-software implementation of the **Advanced Encryption Standard (AES)** written in Go. It supports all three standard key lengths (**AES-128**, **AES-192**, and **AES-256**) and includes comprehensive validation test suites and performance benchmarking tools.

---

### What is AES?

The **Advanced Encryption Standard (AES)** is a symmetric block cipher algorithm established by the U.S. National Institute of Standards and Technology (NIST). It operates on a fixed 128-bit (16-byte) block size and supports three key lengths:
* **AES-128**: Uses a 128-bit key across 10 rounds of transformation.
* **AES-192**: Uses a 192-bit key across 12 rounds of transformation.
* **AES-256**: Uses a 256-bit key across 14 rounds of transformation.

The cipher relies on a substitution-permutation network (SPN) architecture involving core operations such as byte substitution (`SubBytes`), row shifting (`ShiftRows`), column mixing (`MixColumns`), and key addition (`AddRoundKey`).

---

### Project Structure

The project is structured as a modular package (`aesCore`) containing the core cryptographic logic, automated tests, and performance benchmarks:

```text
AES-Implementation/
└── aesCore/
    ├── addRoundKey.go       # Implements the AddRoundKey state operation
    ├── aes.go               # Main encryption, decryption, and key expansion coordination
    ├── aes_test.go          # Functional, round-trip, and NIST test vector validation
    ├── aes_bench_test.go    # Performance benchmarking suite (1 MB and 10 MB workloads)
    ├── keySchedule.go       # Key expansion schedule for 128, 192, and 256-bit keys
    ├── mixColumns.go        # Galois Field GF(2^8) arithmetic and MixColumns / InvMixColumns
    ├── shiftRows.go         # Row shifting and inverse row shifting operations
    ├── subBytes.go          # Forward and inverse S-box byte substitution tables/mappings
    └── utils.go             # Helper utilities (e.g., byte-level XOR operations)
```

### Go Installation

To build and test this project, you need the Go runtime installed on your machine.
Download Go: Visit the official Go Downloads page and download the installer for your operating system (Linux, macOS, or Windows).
Verify Installation: Open your terminal or command prompt and verify the installation by running:

```bash
    go version
```

### How to Test and Execute the Code
Navigate to the project root directory where the aesCore package is located to run the tests and benchmarks.

1. Run Functional and NIST Test Suites
To verify that all encryption, decryption, inverse transformations, and official NIST test vectors pass correctly:

```bash
    go test -v ./aesCore
```

2. Run Performance Benchmarks
To evaluate the encryption and decryption performance (measuring execution time and throughput for 1 MB and 10 MB workloads across 3 repeated runs):
```bash
    go test -v -run TestAESBenchmark ./aesCore
```

---

# C++ Cryptography Toolkit & Integer Factorization Algorithms

A robust C++ implementation of core cryptographic utilities, integer factorization algorithms, automated unit tests, and experimental performance benchmarks.

---

## Project Structure

```text
├── include/
│   ├── fermat_method.h       # Fermat's factorization header
│   ├── pollards_rho.h        # Pollard's Rho factorization header
│   ├── rsa.h                 # RSA encryption/decryption header
│   ├── rsa_utils.h           # Extended GCD & modular inverse utilities
│   └── trial_division.h      # Trial Division factorization header
├── src/
│   ├── benchmark.cc          # Experimental evaluation & Fermat analysis
│   ├── fermat_method.cc      # Precision-safe square root & Fermat implementation
│   ├── main.cc               # Interactive toolkit CLI
│   ├── pollards_rho.cc       # Pollard's Rho cycle-finding implementation
│   ├── rsa.cc                # RSA cryptographic operations
│   ├── rsa_utils.cc          # Extended Euclidean algorithm implementation
│   └── trial_division.cc     # Trial Division implementation
├── tests/
│   └── test_implementation.cc # Catch2 unit test suite
└── CMakeLists.txt            # CMake build configuration

# Features & Getting Started

## Features

* **RSA Cryptographic Suite**: Key generation, modular exponentiation, and full encryption/decryption round-trip validation using custom Extended GCD modular inverses.
* **Integer Factorization Toolkit**:
  * **Trial Division**: Baseline sequential factor search.
  * **Fermat's Method**: Difference-of-squares factorization optimized with precision-safe 128-bit binary search square roots.
  * **Pollard's Rho**: Probabilistic cycle-finding factorization algorithm.
* **Automated Unit Testing**: Comprehensive test suite powered by Catch2 validating GCD bounds, modular inverses, square roots, factor verification ($p \times q = n$), and RSA round-trip consistency.
* **Experimental Benchmarking**: Performance evaluation program measuring execution speeds across varying bit sizes and analyzing the impact of factor distance ($\vert{}p - q\vert{}$) on Fermat's efficiency.

## Getting Started

### Prerequisites

* A modern C++ compiler supporting C++17 or higher (e.g., `g++` / `clang++`).
* CMake (version 3.14 or higher).

´´´bash
mkdir build
cd build
cmake ..
make
´´´

# Build Outputs

This will generate three primary executables inside your `build` directory:

* **`app`**: The interactive Cryptography Toolkit CLI.
* **`run_tests`**: The Catch2 automated test suite.
* **`run_benchmark`**: The experimental evaluation & Fermat analysis runner.
