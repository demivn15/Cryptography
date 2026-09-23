#include <catch2/catch_test_macros.hpp>
#include "trial_division.hpp"
#include "fermat.hpp"
#include "pollards_rho.hpp"

TEST_CASE("TrialDivision factorization is correct", "[trial_division]") {
    TrialDivision solver;

    SECTION("Factoring a composite number") {
        auto [p, q] = solver.trialDivision(3233);
        REQUIRE(p * q == 3233);
        REQUIRE(p > 1);
        REQUIRE(q > 1);
    }

    SECTION("Handling edge cases like primes") {
        auto [p, q] = solver.trialDivision(13);
        REQUIRE(p == 1);
        REQUIRE(q == 13);
    }
}

TEST_CASE("Fermat factorization is correct", "[fermat]") {
    Fermat solver;

    SECTION("Factoring numbers with close factors") {
        auto [p, q] = solver.fermat(3233); // 13 * 17
        REQUIRE(p * q == 3233);
    }
}

TEST_CASE("PollardsRho factorization is correct", "[pollards_rho]") {
    PollardsRho solver;

    SECTION("Factoring a composite number") {
        auto [p, q] = solver.pollardsRho(8051); // 83 * 97
        REQUIRE(p * q == 8051);
    }
}
