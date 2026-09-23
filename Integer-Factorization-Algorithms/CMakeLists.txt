cmake_minimum_required(VERSION 3.14)
project(FactorizationProject)

set(CMAKE_CXX_STANDARD 17)

# 1. Automatically fetch Catch2 v3 for tests
include(FetchContent)
FetchContent_Declare(
    Catch2
    URL https://github.com/catchorg/Catch2/archive/refs/tags/v3.4.0.tar.gz
)
FetchContent_MakeAvailable(Catch2)

# 2. Define the main application executable
add_executable(factor_app
    src/main.cpp
    src/fermat.cpp
    src/trial_division.cpp
    src/pollards_rho.cpp
    src/rsa.cpp
)

# Tell CMake where to look for header files (#include "...")
target_include_directories(factor_app PRIVATE include)

# 3. Define the unit test executable
add_executable(run_tests
    tests/test_factorization.cpp
    src/fermat.cpp
    src/trial_division.cpp
    src/pollards_rho.cpp
    src/rsa.cpp
)

target_include_directories(run_tests PRIVATE include)
target_link_libraries(run_tests PRIVATE Catch2::Catch2WithMain)
