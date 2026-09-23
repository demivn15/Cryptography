#include <iostream>
#include "rsa.hpp"

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
