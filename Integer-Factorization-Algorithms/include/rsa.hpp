#pragma once
#include <cstdint>

class RSA {
public:
    static uint64_t modExp(uint64_t base, uint64_t exp, uint64_t mod) { //
        if (mod == 1) return 0;
        
        __int128 result = 1;
        __int128 b = base % mod;
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

    // Encrypts a numeric message: c = (m^e) % n
    static uint64_t encrypt(uint64_t message, uint64_t e, uint64_t n) {
        return modExp(message, e, n);
    }

    // Decrypts a numeric ciphertext: m = (c^d) % n
    static uint64_t decrypt(uint64_t ciphertext, uint64_t d, uint64_t n) {
        return modExp(ciphertext, d, n);
    }
};
