#include "pollards_rho.hpp"
#include <numeric>
#include <utility>

PollardsRho::PollardsRho() : m_factorization({0, 0}) {}

std::pair<uint64_t, uint64_t> PollardsRho::pollardsRho(uint64_t targetInteger) {
    if (targetInteger <= 1) return {1, targetInteger};
    if (targetInteger % 2 == 0) return {2, targetInteger / 2};
    if (targetInteger % 3 == 0) return {3, targetInteger / 3};

    uint64_t x = 2;
    uint64_t y = 2;
    uint64_t c = 1;
    uint64_t d = 1;

    // Lambda function with __int128 to safely handle multiplication/squaring without overflow
    auto f = [&](uint64_t v, uint64_t mod) {
        __int128 val = v;
        return static_cast<uint64_t>((val * val + c) % mod);
    };

    int attempts = 0;
    while (d == 1) {
        x = f(x, targetInteger);
        y = f(f(y, targetInteger), targetInteger);

        uint64_t diff = (x > y) ? (x - y) : (y - x);
        d = std::gcd(diff, targetInteger);

        if (d == targetInteger) {
            c++;
            x = 2;
            y = 2;
            d = 1;
            attempts++;
            if (attempts > 50) break; // Safeguard against endless looping
        }
    }

    if (d > 1 && d < targetInteger) {
        return {d, targetInteger / d};
    }
    
    return {1, targetInteger}; // Fallback if no non-trivial factor found
}
