#pragma once
#include <cstdint>
#include <utility>

class PollardsRho {
    public:
        PollardsRho();
        std::pair<uint64_t, uint64_t> pollardsRho(uint64_t targetInteger);
        uint64_t rho(uint64_t x, uint64_t targetInteger, uint64_t constant);
    private:
        std::pair<uint64_t, uint64_t> m_factorization;
};
