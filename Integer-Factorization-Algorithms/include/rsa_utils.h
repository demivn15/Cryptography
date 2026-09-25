#ifndef RSA_UTILS_H_
#define RSA_UTILS_H_

#include <cstdint>

class RSAUtils {
    public:
        static uint64_t ModInverse(uint64_t e, uint64_t phi);
        static int64_t ExtendedGCD(int64_t a, int64_t b, int64_t &x, int64_t &y);
};

#endif