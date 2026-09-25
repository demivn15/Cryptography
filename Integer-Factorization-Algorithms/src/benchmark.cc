/**
 * @file benchmark.cc
 * @author Demian Viteri
 * @brief Experimental evaluation benchmark.
 * @license GPLv3
 */

#include "trial_division.h"
#include "fermat_method.h"
#include "pollards_rho.h"

#include <iostream>
#include <chrono>
#include <vector>
#include <numeric>
#include <iomanip>
#include <cmath>

int GetBitSize(uint64_t n) {
// Helper to calculate bit size of a 64-bit integer
    if (n == 0) return 0;
    return 64 - __builtin_clzll(n);
}

template <typename Func>
double MeasureExecutionTime(Func&& func, int trials = 3) {
// Function to measure average execution time over 3 runs (in microseconds)
    double total_duration = 0.0;
    for (int i = 0; i < trials; ++i) {
        auto start = std::chrono::high_resolution_clock::now();
        func();
        auto end = std::chrono::high_resolution_clock::now();
        std::chrono::duration<double, std::micro> elapsed = end - start;
        total_duration += elapsed.count();
    }
    return total_duration / trials;
}

void RunExercise6Benchmarks() {
    std::cout << "=================================================================\n";
    std::cout << "          EXERCISE 6: EXPERIMENTAL BENCHMARK TABLE             \n";
    std::cout << "=================================================================\n";
    std::cout << std::left << std::setw(6) << "Bits" 
              << std::setw(22) << "Modulus (n)" 
              << std::setw(18) << "Trial Div (us)" 
              << std::setw(15) << "Fermat (us)" 
              << std::setw(15) << "Pollard's Rho (us)" << "\n";
    std::cout << "-----------------------------------------------------------------\n";

    // Four moduli of different bit sizes
    std::vector<uint64_t> test_moduli = {
        10403,                // ~14 bits (101 * 103)
        4295098363ULL,        // ~32 bits (65537 * 65539)
        3599,                 // ~12 bits (59 * 61 - close factors)
        9223372019157690407ULL // ~63 bits 
    };

    TrialDivision td;
    Fermat fermat;
    PollardsRho rho;

    for (uint64_t n : test_moduli) {
        int bits = GetBitSize(n);

        double td_time = MeasureExecutionTime([&]() { td.ComputeTrialDivision(n); });
        double fermat_time = MeasureExecutionTime([&]() { fermat.ComputeFermat(n); });
        double rho_time = MeasureExecutionTime([&]() { rho.ComputePollardsRho(n); });

        std::cout << std::left << std::setw(6) << bits 
                  << std::setw(22) << n 
                  << std::setw(18) << std::fixed << std::setprecision(2) << td_time 
                  << std::setw(15) << fermat_time 
                  << std::setw(15) << rho_time << "\n";
    }
    std::cout << "=================================================================\n\n";
}

void RunExercise7FermatAnalysis() {
    std::cout << "=================================================================\n";
    std::cout << "             EXERCISE 7: FERMAT ANALYSIS                         \n";
    std::cout << "=================================================================\n";

    // Case A: p and q close
    uint64_t p_close = 65537;
    uint64_t q_close = 65539;
    uint64_t n_close = p_close * q_close;
    uint64_t diff_close = q_close - p_close;

    // Case B: p and q far apart
    uint64_t p_far = 13;
    uint64_t q_far = 709490155319822339ULL;
    uint64_t n_far = p_far * q_far;
    uint64_t diff_far = q_far - p_far;

    Fermat fermat;

    std::cout << "[Case A: Close Factors]\n";
    std::cout << "  Modulus (n)  = " << n_close << "\n";
    std::cout << "  |p - q|      = " << diff_close << "\n";
    auto start_a = std::chrono::high_resolution_clock::now();
    fermat.ComputeFermat(n_close);
    auto end_a = std::chrono::high_resolution_clock::now();
    std::chrono::duration<double, std::micro> elapsed_a = end_a - start_a;
    std::cout << "  Execution Time = " << elapsed_a.count() << " us\n\n";

    std::cout << "[Case B: Far-Apart Factors]\n";
    std::cout << "  Modulus (n)  = " << n_far << "\n";
    std::cout << "  |p - q|      = " << diff_far << "\n";
    auto start_b = std::chrono::high_resolution_clock::now();
    fermat.ComputeFermat(n_far);
    auto end_b = std::chrono::high_resolution_clock::now();
    std::chrono::duration<double, std::micro> elapsed_b = end_b - start_b;
    std::cout << "  Execution Time = " << elapsed_b.count() << " us (Timed out / Max iterations)\n";
    std::cout << "=================================================================\n";
}

int main() {
    RunExercise6Benchmarks();
    RunExercise7FermatAnalysis();
    return 0;
}