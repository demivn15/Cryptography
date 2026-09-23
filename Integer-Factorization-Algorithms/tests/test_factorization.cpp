#include <catch2/catch_test_macros.hpp>
#include "trial_division.hpp"
#include "fermat.hpp"
#include "pollards_rho.hpp"
#include "rsa.hpp"
#include "extended_euc.hpp"

// ==========================================
// 1. Trial Division Tests
// ==========================================
TEST_CASE("Trial Division factorization is correct", "[trial_division]") {
    TrialDivision td;

    SECTION("Factoring a composite number") {
        auto result = td.trialDivision(35);
        uint64_t p = result.first;
        uint64_t q = result.second;
        // Verify they multiply back to 35 (factors can be 5, 7 or 7, 5)
        REQUIRE((p * q == 35));
        REQUIRE((p == 5 || p == 7));
    }

    SECTION("Factoring a prime number") {
        auto result = td.trialDivision(13);
        REQUIRE(result.first == 1);
        REQUIRE(result.second == 13);
    }
}

// ==========================================
// 2. Fermat's Factorization Tests
// ==========================================
TEST_CASE("Fermat factorization is correct", "[fermat]") {
    Fermat fermat;

    SECTION("Factoring a composite number with close factors") {
        // 59 * 61 = 3599
        auto result = fermat.fermat(3599);
        uint64_t p = result.first;
        uint64_t q = result.second;
        REQUIRE((p * q == 3599));
        REQUIRE((p == 59 || p == 61));
    }

    SECTION("Handling even numbers") {
        auto result = fermat.fermat(38);
        REQUIRE((result.first * result.second == 38));
    }
}

// ==========================================
// 3. Pollard's Rho Tests
// ==========================================
TEST_CASE("PollardsRho factorization is correct", "[pollards_rho]") {
    PollardsRho rho;

    SECTION("Factoring a composite number") {
        // 10403 = 101 * 103
        auto result = rho.pollardsRho(10403);
        uint64_t p = result.first;
        uint64_t q = result.second;
        REQUIRE((p * q == 10403));
        REQUIRE((p == 101 || p == 103 || q == 101 || q == 103));
    }
}

// ==========================================
// 4. Extended Euclidean Algorithm Tests
// ==========================================
TEST_CASE("Extended Euclidean modular inverse is correct", "[extended_euc]") {
    SECTION("Computing valid modular inverse") {
        // e = 17, phi = 3120 -> d should be 2753
        uint64_t d = RSAUtils::modInverse(17, 3120);
        REQUIRE(d == 2753);
    }

    SECTION("Modular inverse fails for non-coprime inputs") {
        // 18 and 30 share a factor of 6, inverse does not exist
        uint64_t d = RSAUtils::modInverse(18, 30);
        REQUIRE(d == 0);
    }
}

// ==========================================
// 5. RSA Encryption & Decryption Tests
// ==========================================
TEST_CASE("RSA encryption and decryption cycle is correct", "[rsa]") {
    SECTION("Full RSA roundtrip with standard educational keys") {
        uint64_t p = 61;
        uint64_t q = 53;
        uint64_t n = p * q;              // 3233
        uint64_t phi = (p - 1) * (q - 1); // 3120
        uint64_t e = 17;
        uint64_t d = RSAUtils::modInverse(e, phi); // 2753

        uint64_t originalMessage = 42;

        uint64_t ciphertext = RSA::encrypt(originalMessage, e, n);
        uint64_t decryptedMessage = RSA::decrypt(ciphertext, d, n);

        REQUIRE(ciphertext != originalMessage);
        REQUIRE(decryptedMessage == originalMessage);
    }
}