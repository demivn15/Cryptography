/**
 * @file test_implementation.cc
 * @author Demian Viteri
 * @brief Test the different functions from the project.
 * @license GPLv3
 */

#include "fermat_method.h"
#include "pollards_rho.h"
#include "rsa.h"
#include "rsa_utils.h"
#include "trial_division.h"

#include <catch2/catch_test_macros.hpp>

TEST_CASE("Greatest Common Divisor and Extended GCD are correct", "[rsa_utils]") {
    SECTION("Basic GCD calculation via ExtendedGCD") {
        int64_t x = 0, y = 0;
        // GCD(48, 18) = 6
        int64_t gcd = RSAUtils::ExtendedGCD(48, 18, x, y);
        REQUIRE(gcd == 6);
        // Verify Bézout's identity: a*x + b*y = gcd
        REQUIRE((48 * x + 18 * y) == gcd);
    }
    SECTION("GCD with coprime numbers") {
        int64_t x = 0, y = 0;
        // GCD(35, 64) = 1
        int64_t gcd = RSAUtils::ExtendedGCD(35, 64, x, y);
        REQUIRE(gcd == 1);
        REQUIRE((35 * x + 64 * y) == 1);
    }
}

TEST_CASE("Extended Euclidean modular inverse is correct", "[rsa_utils]") {
    SECTION("Computing valid modular inverse") {
        // e = 17, phi = 3120 -> d should be 2753
        uint64_t d = RSAUtils::ModInverse(17, 3120);
        REQUIRE(d == 2753);
    }
    SECTION("Modular inverse fails for non-coprime inputs") {
        // 18 and 30 share a factor of 6, inverse does not exist
        uint64_t d = RSAUtils::ModInverse(18, 30);
        REQUIRE(d == 0);
    }
}

TEST_CASE("Trial Division factorization and verification are correct", "[trial_division]") {
    TrialDivision td;
    SECTION("Factoring a composite number with verification (p * q == n)") {
        uint64_t n = 35;
        auto result = td.ComputeTrialDivision(n);
        uint64_t p = result.first;
        uint64_t q = result.second;
        REQUIRE(p > 0);
        REQUIRE(q > 0);
        REQUIRE((p * q == n));
        REQUIRE((p == 5 || p == 7));
    }
    SECTION("Factoring a prime number returns {0,0}") {
        auto result = td.ComputeTrialDivision(13);
        REQUIRE(result.first == 0);
        REQUIRE(result.second == 0);
    }
    SECTION("Factoring 1 returns {0,0}") {
        auto result = td.ComputeTrialDivision(1);
        REQUIRE(result.first == 0);
        REQUIRE(result.second == 0);
    }
}

TEST_CASE("Fermat factorization and precision-safe square root behavior are correct", "[fermat_method]") {
    Fermat fermat;
    SECTION("Factoring a composite number with close factors and verification") {
        // 59 * 61 = 3599
        uint64_t n = 3599;
        auto result = fermat.ComputeFermat(n);
        uint64_t p = result.first;
        uint64_t q = result.second;
        REQUIRE(p > 0);
        REQUIRE(q > 0);
        REQUIRE((p * q == n));
        REQUIRE((p == 59 || p == 61));
    }
    SECTION("Factoring a composite number with far-apart factors times out safely") {
        // 3 * 1000000007 = 3000000021
        auto result = fermat.ComputeFermat(3000000021);
        REQUIRE(result.first == 0);
        REQUIRE(result.second == 0);
    }
    SECTION("Handling even numbers") {
        uint64_t n = 38;
        auto result = fermat.ComputeFermat(n);
        uint64_t p = result.first;
        uint64_t q = result.second;
        REQUIRE(p > 0);
        REQUIRE(q > 0);
        REQUIRE((p * q == n));
    }
}

TEST_CASE("Pollard's Rho factorization and verification are correct", "[pollards_rho]") {
    PollardsRho rho;
    SECTION("Factoring a composite number with verification") {
        // 10403 = 101 * 103
        uint64_t n = 10403;
        auto result = rho.ComputePollardsRho(n);
        uint64_t p = result.first;
        uint64_t q = result.second;
        REQUIRE(p > 0);
        REQUIRE(q > 0);
        REQUIRE((p * q == n));
        REQUIRE((p == 101 || p == 103 || q == 101 || q == 103));
    }
    SECTION("Factoring a prime number returns {0,0}") {
        auto result = rho.ComputePollardsRho(13);
        REQUIRE(result.first == 0);
        REQUIRE(result.second == 0);
    }
    SECTION("Factoring 1 returns {0,0}") {
        auto result = rho.ComputePollardsRho(1);
        REQUIRE(result.first == 0);
        REQUIRE(result.second == 0);
    }
}

TEST_CASE("RSA encryption and decryption consistency is correct", "[rsa]") {
    SECTION("Full RSA roundtrip consistency") {
        uint64_t p = 104729;
        uint64_t q = 1299709;
        uint64_t n = p * q;
        uint64_t phi = (p - 1) * (q - 1);
        uint64_t e = 65537;
        uint64_t d = RSAUtils::ModInverse(e, phi);
        
        REQUIRE(d != 0);

        uint64_t originalMessage = 987654321;
        uint64_t ciphertext = RSA::Encrypt(originalMessage, e, n);
        uint64_t decryptedMessage = RSA::Decrypt(ciphertext, d, n);

        REQUIRE(ciphertext != originalMessage);
        REQUIRE(decryptedMessage == originalMessage);
    }
}