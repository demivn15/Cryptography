#pragma once

class RSA {
    public:
        RSA();
        static uint64_t modularExponentiation(uint64_t baseElement, uint64_t exp, uint64_t mod);
        static uint64_t encrypt(uint64_t plaintext, uint64_t e, uint64_t n);
        static uint64_t decrypt(uint64_t ciphertext, uint64_t d, uint64_t n);
    private:
        static uint64_t m_expResult;
        static uint64_t encryption;
        static uint64_t decryption;
};