/**
 * @file trial_division.cc
 * @author Demian Viteri
 * @brief Implementation of the Trial Division Algorithm for factorization.
 * @license GPLv3
 */

#include "trial_division.h"

#include <utility>

std::pair<uint64_t, uint64_t> TrialDivision::ComputeTrialDivision(uint64_t target_integer) {
    // Check if the integer value can be factorized. Return a {0, 0} pair if not.
    if (target_integer <= 1) {
        return {0, 0};
    }
    // Check if it is divisible by 2 and return the value immediately.
    if (target_integer % 2 == 0) {
        return {2, target_integer / 2};
    }
    // Check if the integer is divisible by 3 and return the value immediately.
    if (target_integer % 3 == 0) {
        return {3, target_integer / 3};
    }
    // Check divisibility only up to the square root of the target integer. Start at 5 and avoid iterating through multiples of 2 or 3.
    uint64_t divisor_p = 5;
    while (divisor_p * divisor_p <= target_integer) {
        if (target_integer % divisor_p == 0)
            return {divisor_p, target_integer / divisor_p};
        if (target_integer % (divisor_p + 2) == 0)
            return {divisor_p + 2, target_integer / (divisor_p + 2)};
        divisor_p += 6;
    }
    return {0, 0};
}