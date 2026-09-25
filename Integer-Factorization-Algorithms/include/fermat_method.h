#ifndef FERMAT_METHOD_H_
#define FERMAT_METHOD_H_

#include <cstdint>
#include <utility>

class Fermat {
    public:
        std::pair<uint64_t, uint64_t> ComputeFermat(uint64_t target_integer);
};

#endif