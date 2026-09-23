#include "extended_euc.hpp"
#include <cstdint>

RSAUtils::RSAUtils() {}

int64_t RSAUtils::extendedGCD(int64_t a, int64_t b, int64_t &x, int64_t &y){
            if (b == 0) {
                x = 1;
                y = 0;
                return a;
            }
            int64_t x1, y1;
            int64_t gcd = extendedGCD(b, a % b, x1, y1);
            x = y1;
            y = x1 - (a / b) * y1;
            return gcd;
}

uint64_t RSAUtils::modInverse(uint64_t e, uint64_t phi) {
            int64_t x = 0;
            int64_t y = 0;
            int64_t gcd = extendedGCD(static_cast<int64_t>(e), static_cast<int64_t>(phi), x, y);
            if (gcd != 1) {
                return 0; // Inverse does not exist if e and phi are not coprime
            }
            // x might be negative, so we normalize it into the range [0, phi - 1]
            int64_t phi_s = static_cast<int64_t>(phi);
            int64_t result = (x % phi_s + phi_s) % phi_s;
            return static_cast<uint64_t>(result);
}
