#pragma once
#include <utility>
#include <cstdint>

class Fermat {
    public:
        Fermat();
        std::pair<uint64_t, uint64_t> fermat(uint64_t targetInteger);
    private:
        std::pair<uint64_t, uint64_t> m_factorization;
};
