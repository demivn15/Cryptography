
# 🔑 Cryptography

## Lab 1 - DES Algorithm Implementation

Lab1 contains the implementation of the DES Algorithm written in Go.

### Environment setup

- Operating System: Linux Debian 13
- Programming Language: Go 1.24.4

### Implementation details

The project consist of a local module called `utils` and the module with the DES Algorithm implementation: `myDES`. For comparison purposes, the `crypto/des` module is used inside `myDES_test.go`. Other Go modules are used to handle I/O operations, conversions, time metrics, testing, etc.

- The `utils` module contains functions that help with converting raw strings to binary and the other way around. It also contains the function that helps with the rotation of bits inside the algorithm, and a function to handle user input.
- The `myDES` module contains the actual implementation of the DES Algorithm. There, the different functions in charge of the permutations, fesitel, expansion, S-Boxes mapping, etc., are defined.

When running the program, the user is prompted to enter a string. Then, the encryption and decryption process occurs. Finally, the program outputs the actual string entered by the user, the encryption and decryption results and the time it took to execute both processes.

When running the `myDES_test.go` file, the DES implementation is tested against the one that is already imlemented inside the `crypto/des` module. Both encryption and decryption are tested.

> A known issue is that when the user enters a strign consisting of only numbers, the conversion to strings, later on, generates an ascii character.

#### How to run:

- Clone this repo: `git clone https://github.com/demivn15/Cryptography`.
- Install Go in your system (in Debian 13: `sudo apt install golang-go`).
- Inside the Project directory execute `go run .` to run the programm without compiling it.
- If you want to compile the project, you can run `go build` inside the Project directory. An executable will be generated.
- In order to test the implementation against the DES Algorithm from the `crypto/des` module, execute the following command: `go test -v .` in the myDES directory inside the Project folder.

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

---
