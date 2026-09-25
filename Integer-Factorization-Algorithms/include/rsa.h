#ifndef RSA_H_
#define RSA_H_

#include <cstdint>

class RSA {
    public:
        static uint64_t Encrypt(uint64_t plaintext, uint64_t e, uint64_t n);
        static uint64_t Decrypt(uint64_t ciphertext, uint64_t d, uint64_t n);
    private:
        static uint64_t ModularExponentiation(uint64_t base_element, uint64_t exp, uint64_t mod);
};

#endif