#include "fermat.hpp"
#include <cmath>
#include <utility>

Fermat::Fermat() : m_factorization({0, 0}) {}

std::pair<uint64_t, uint64_t> Fermat::fermat(uint64_t targetInteger) {
    uint64_t a = ceil(sqrt(targetInteger));
    uint64_t b_square;
    uint64_t b;
    do {
        b_square = a * a - targetInteger;
        b = sqrt(b_square);
        a += 1;
    } while ((b * b) != b_square);
    uint64_t factor_p = a - 1 - b;
    uint64_t factor_q = a - 1 + b;
    return {factor_p, factor_q};
}
