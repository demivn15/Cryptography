#pragma once
#include <utility>
#include <cstdint>

class TrialDivision {
    public:
        TrialDivision();
        std::pair<uint64_t, uint64_t> trialDivision(uint64_t targetInteger);
    private:
        std::pair<uint64_t, uint64_t> m_factorization;
};
