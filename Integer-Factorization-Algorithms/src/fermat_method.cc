/**
 * @file fermat_method.cc
 * @author Demian Viteri
 * @brief Implementation of the Fermat's Algorithm to factor out a number into prime numbers.
 * @license GPLv3
 */

#include "fermat_method.h"

#include <algorithm>
#include <cmath>
#include <utility>

static uint64_t IntegerSqrt(uint64_t n) {
// Helper function to safely compute integer square root for uint64_t without precision loss
    if (n <= 1) return n;
    __int128 low = 1, high = n, result = 1;
    while (low <= high) {
        __int128 mid = low + (high - low) / 2;
        if (mid * mid <= n) {
            result = mid;
            low = mid + 1;
        } else {
            high = mid - 1;
        }
    }
    return static_cast<uint64_t>(result);
}

std::pair<uint64_t, uint64_t> Fermat::ComputeFermat(uint64_t target_integer) {
    // Check if the target integer is suitable to be factored out.
    if (target_integer <= 1) {
        return {0, 0};
    }
    // If the target integer is divisible by 2, return the value without going into the loop.
    if (target_integer % 2 == 0) {
        return {2, target_integer / 2};
    }
    // Compute precise integer square root
    uint64_t root = IntegerSqrt(target_integer);
    
    // Check if it's a perfect square (e.g., 9, 25, 49)
    if (root * root == target_integer) {
        return {root, root};
    }

    __int128 a = root + 1;
    __int128 target = target_integer;
    __int128 b_square;
    uint64_t b;

    // Set a maximum order of iterations to avoid entering an infinite loop.
    uint64_t iterations = 0;
    const uint64_t kMaxIterations = 10000000;

    do {
        b_square = (a * a) - target;
        
        if (b_square >= 0) {
            b = IntegerSqrt(static_cast<uint64_t>(b_square));
            if (static_cast<__int128>(b) * b == b_square) {
                break; 
            }
        }

        a++;
        iterations++;
        if (iterations > kMaxIterations) {
            return {0, 0};
        }
    } while (true);

    uint64_t factor_p = static_cast<uint64_t>(a - b);
    uint64_t factor_q = static_cast<uint64_t>(a + b);
    return {factor_p, factor_q};
}