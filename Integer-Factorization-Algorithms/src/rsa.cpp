// RSA Implementation.

#include <cstdint>
#include <iostream>
#include "rsa.hpp"

RSA:: RSA() {}

uint64_t RSA::modularExponentiation(uint64_t baseElement, uint64_t exp, uint64_t mod) {
    if (mod == 1) return 0;
    __int128 result = 1;
    __int128 b = baseElement % mod;
    __int128 e = exp;
    __int128 m = mod;
    while (e > 0) {
        if (e % 2 == 1) {
            result = (result * b) % m;
        }
        b = (b * b) % m;
        e /= 2;
    }
    return static_cast<uint64_t>(result);
}

uint64_t RSA::encrypt(uint64_t plaintext, uint64_t e, uint64_t n) {
    return modularExponentiation(plaintext, e, n);
}
uint64_t RSA::decrypt(uint64_t ciphertext, uint64_t d, uint64_t n){
    return modularExponentiation(ciphertext, d, n);
}

/*

int main() {
    // Example small key parameters (Educational)
    uint64_t p = 61;
    uint64_t q = 53;
    uint64_t n = p * q;            // 3233 (Modulus)
    uint64_t phi = (p - 1) * (q - 1); // 3120
    
    uint64_t e = 17;               // Public exponent
    uint64_t d = 2753;             // Private exponent (modular inverse of e mod phi)

    uint64_t originalMessage = 65; // Message to encrypt (must be < n)

    // Encrypt
    uint64_t ciphertext = RSA::encrypt(originalMessage, e, n);
    std::cout << "Encrypted Ciphertext: " << ciphertext << "\n";

    // Decrypt
    uint64_t decryptedMessage = RSA::decrypt(ciphertext, d, n);
    std::cout << "Decrypted Message: " << decryptedMessage << "\n";

    return 0;
}

*/