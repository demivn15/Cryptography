// Implementation of the Fermat's Algorithm.

/*
    The objective is to factor out an integer value by rewriting it as the difference of two squares, i.e., n = a ^ 2 - b ^ 2.
    The Fermat's Algorithm allows to find the values for a and b.
*/

#include "fermat.hpp"
#include <cmath>
#include <utility>

Fermat::Fermat() : m_factorization({0, 0}) {}

std::pair<uint64_t, uint64_t> Fermat::fermat(uint64_t targetInteger) {
    // Check if the target integer is suitable to be factored out.
    if (targetInteger <= 1) return {0, 0};
    // If the target integer is divisible by 2, return the value without going into the loop.
    if (targetInteger % 2 == 0) return {2, targetInteger / 2};
    // Fermat's Algorithm:
    uint64_t a = ceil(sqrt(targetInteger));
    uint64_t b_square;
    uint64_t b;
    uint64_t iterations = 0;
    const uint64_t MAX_ITERATIONS = 10000000;
    do {
        b_square = a * a - targetInteger;
        b = static_cast<uint64_t>(std::sqrt(b_square));
        if (b * b == b_square) {
            break; 
        }
        a++;
        iterations++;
        if (iterations > MAX_ITERATIONS) return {1, targetInteger};
    } while (true);
    uint64_t factor_p = a - b;
    uint64_t factor_q = a + b;
    return {factor_p, factor_q};
}