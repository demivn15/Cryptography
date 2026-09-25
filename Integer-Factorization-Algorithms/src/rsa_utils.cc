/**
 * @file rsa_utils.cc
 * @author Demian Viteri
 * @brief Extended Euclidean Algorithm for GDC and extended GCD computing.
 * @license GPLv3
 */

#include "rsa_utils.h"

#include <cstdint>

int64_t RSAUtils::ExtendedGCD(int64_t a, int64_t b, int64_t &x, int64_t &y){
            if (b == 0) {
                x = 1;
                y = 0;
                return a;
            }
            int64_t x1, y1;
            int64_t gcd = ExtendedGCD(b, a % b, x1, y1);
            x = y1;
            y = x1 - (a / b) * y1;
            return gcd;
}

uint64_t RSAUtils::ModInverse(uint64_t e, uint64_t phi) {
            int64_t x = 0;
            int64_t y = 0;
            int64_t gcd = ExtendedGCD(static_cast<int64_t>(e), static_cast<int64_t>(phi), x, y);
            // Inverse does not exist if e and phi are not coprime
            if (gcd != 1) {
                return 0;
            }
            // x might be negative, so we normalize it into the range [0, phi - 1]
            int64_t phi_s = static_cast<int64_t>(phi);
            int64_t result = (x % phi_s + phi_s) % phi_s;
            return static_cast<uint64_t>(result);
}