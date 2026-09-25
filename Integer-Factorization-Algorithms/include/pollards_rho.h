#ifndef POLLARDS_RHO_H_
#define POLLARDS_RHO_H_

#include <cstdint>
#include <utility>

class PollardsRho {
    public:
        std::pair<uint64_t, uint64_t> ComputePollardsRho(uint64_t target_integer);
};

#endif