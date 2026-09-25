#ifndef TRIAL_DIVISION_H_
#define TRIAL_DIVISION_H_

#include <cstdint>
#include <utility>

class TrialDivision {
    public:
        static std::pair<uint64_t, uint64_t> ComputeTrialDivision(uint64_t target_integer);
};

#endif