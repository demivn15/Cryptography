/**
 * @file rsa.cc
 * @author Demian Viteri
 * @brief Implementation of the RSA encryption and decryption.
 * @license GPLv3
 */

#include "rsa.h"

#include <cstdint>

uint64_t RSA::ModularExponentiation(uint64_t base_element, uint64_t exp, uint64_t mod) {
    if (mod == 1) {
        return 0;
    }
    __int128 result = 1;
    __int128 b = base_element % mod;
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

uint64_t RSA::Encrypt(uint64_t plaintext, uint64_t e, uint64_t n) {
    return ModularExponentiation(plaintext, e, n);
}
uint64_t RSA::Decrypt(uint64_t ciphertext, uint64_t d, uint64_t n){
    return ModularExponentiation(ciphertext, d, n);
}