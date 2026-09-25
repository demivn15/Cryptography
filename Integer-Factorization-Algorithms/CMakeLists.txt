cmake_minimum_required(VERSION 3.14)
project(PrimeFactorizationAndRSA)

set(CMAKE_CXX_STANDARD 17)

include(FetchContent)
FetchContent_Declare(
    Catch2
    URL https://github.com/catchorg/Catch2/archive/refs/tags/v3.4.0.tar.gz
)
FetchContent_MakeAvailable(Catch2)

add_executable(app
    src/fermat_method.cc
    src/main.cc
    src/pollards_rho.cc
    src/rsa.cc
    src/rsa_utils.cc
    src/trial_division.cc
)

target_include_directories(app PRIVATE include)

add_executable(run_tests
    src/fermat_method.cc
    src/pollards_rho.cc
    src/rsa.cc
    src/rsa_utils.cc
    src/trial_division.cc
    tests/test_implementation.cc
)

target_include_directories(run_tests PRIVATE include)

add_executable(run_benchmark 
    src/benchmark.cc 
    src/trial_division.cc 
    src/fermat_method.cc 
    src/pollards_rho.cc
)
target_include_directories(run_benchmark PRIVATE include)

target_link_libraries(run_tests PRIVATE Catch2::Catch2WithMain)