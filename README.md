# Cryptography

## Lab 1 - DES Algorithm Implementation

Lab1 contains the implementation of the DES Algorithm written in `Go`.

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
- the xor output is processed to accomplish nonlinearity by shortening the input from 48 bits down to 32 bits;
- then, the result goes through another permutation and becomes the output of the function.

##### Transformation of subkeys breakdown:

- In order to generate the subkeys for each round the resulting key from **performing permuted choice one** is split into halves (28 bits each);
- a left rotation is performed. The number of positions shifted depends on the current round (one position is shifted for rounds 1, 2, 9, 16, and two are shifted for any other round);
- the resulting 56-bits key is performed a **round key permutation** on. This derives a 48-bit long subkey, the one used as input for the Feistel function.

---
