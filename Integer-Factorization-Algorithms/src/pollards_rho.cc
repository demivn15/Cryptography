/**
 * @file pollards_rho.cc
 * @author Demian Viteri
 * @brief Implementation of the Pollard's Rho Factorization Algorithm.
 * @license GPLv3
 */

#include "pollards_rho.h"

#include <numeric>
#include <utility>

std::pair<uint64_t, uint64_t> PollardsRho::ComputePollardsRho(uint64_t target_integer) {
    if (target_integer <= 1) { // Check for suitable number to factor.
        return {0, 0};
    }
    if (target_integer % 2 == 0) { // Check for even numbers.
        return {2, target_integer / 2};
    }
    if (target_integer % 3 == 0) { // Check for numbers divisible by 3 (Optimization).
        return {3, target_integer / 3};
    }
    uint64_t x = 2;
    uint64_t y = 2;
    uint64_t c = 1;
    uint64_t d = 1;

    auto rho = [&](uint64_t v, uint64_t mod) { // Lambda function with __int128 to safely handle multiplication/squaring without overflow
        __int128 val = v;
        return static_cast<uint64_t>((val * val + c) % mod);
    };
    int attempts = 0;
    while (d == 1) {
        x = rho(x, target_integer);
        y = rho(rho(y, target_integer), target_integer);
        uint64_t diff = (x > y) ? (x - y) : (y - x);
        d = std::gcd(diff, target_integer);
        if (d == target_integer) {
            c++;
            x = 2;
            y = 2;
            d = 1;
            attempts++;
            // Avoid endless looping.
            if (attempts > 50) break;
        }
    }
    if (d > 1 && d < target_integer) {
        return {d, target_integer / d};
    }
    return {0, 0};
}