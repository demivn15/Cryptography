#pragma once
#include <cstdint>

class RSAUtils {
    public:
        RSAUtils();
        static uint64_t modInverse(uint64_t e, uint64_t phi);
    private:
        static int64_t extendedGCD(int64_t a, int64_t b, int64_t &x, int64_t &y);
        static int64_t eucResult;
        static int64_t extendedGCDResult;

};
