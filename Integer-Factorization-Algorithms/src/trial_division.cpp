#include "trial_division.hpp"
#include <utility>

TrialDivision::TrialDivision() : m_factorization({0, 0}) {}

std::pair<uint64_t, uint64_t> TrialDivision::trialDivision(uint64_t targetInteger) {
    if (targetInteger <= 1)
        return {0, 0};
    if (targetInteger % 2 == 0)
        return {2, targetInteger / 2};
    if (targetInteger % 3 == 0)
        return {3, targetInteger / 3};
    uint64_t divisor_p = 5;
    while (divisor_p * divisor_p <= targetInteger) {
        if (targetInteger % divisor_p == 0)
            return {divisor_p, targetInteger / divisor_p};
        if (targetInteger % (divisor_p + 2) == 0)
            return {divisor_p + 2, targetInteger / (divisor_p + 2)};
        divisor_p += 6;
    }
    return {1, targetInteger};
}
