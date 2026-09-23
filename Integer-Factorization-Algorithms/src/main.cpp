// Main program for Cryptography & Integer Factorization Toolkit.

#include "fermat.hpp"
#include "pollards_rho.hpp"
#include "rsa.hpp"
#include "extended_euc.hpp"
#include "trial_division.hpp"
#include <cstdint>
#include <iostream>

void printMenu() {
    std::cout << "\n====================================\n";
    std::cout << "       CRYPTOGRAPHY TOOLKIT         \n";
    std::cout << "====================================\n";
    std::cout << "0. Test RSA Encryption/Decryption & Key Gen\n";
    std::cout << "------------------------------------\n";
    std::cout << "      INTEGER FACTORIZATION TOOL    \n";
    std::cout << "------------------------------------\n";
    std::cout << "1. Trial Division\n";
    std::cout << "2. Fermat's Factorization Method\n";
    std::cout << "3. Pollard's Rho Algorithm\n";
    std::cout << "4. Exit\n";
    std::cout << "Choose an option (0-4): ";
}

int main() {
    uint64_t target = 0;
    int choice = -1;

    while (true) {
        printMenu();
        if (!(std::cin >> choice)) {
            std::cout << "Invalid input. Please enter a valid number.\n";
            std::cin.clear();
            std::cin.ignore(10000, '\n');
            continue;
        }

        if (choice == 4) {
            std::cout << "Exiting program. Goodbye!\n";
            break;
        }

        if (choice < 0 || choice > 4) {
            std::cout << "Invalid choice. Please select between 0 and 4.\n";
            continue;
        }

        // RSA Demonstration & Private Key Generation (Choice 0)
        if (choice == 0) {
            std::cout << "\n====================================\n";
            std::cout << "         RSA DEMONSTRATION          \n";
            std::cout << "====================================\n";
            
            uint64_t p = 0, q = 0, e = 0, message = 0;
            std::cout << "Enter prime number p (e.g., 61): ";
            std::cin >> p;
            std::cout << "Enter prime number q (e.g., 53): ";
            std::cin >> q;
            std::cout << "Enter public exponent e (e.g., 17): ";
            std::cin >> e;
            std::cout << "Enter numeric message to encrypt (must be < n): ";
            std::cin >> message;

            uint64_t n = p * q;
            uint64_t phi = (p - 1) * (q - 1);
            
            // Computes modular inverse using Extended Euclidean Algorithm
            uint64_t d = RSAUtils::modInverse(e, phi);

            if (d == 0) {
                std::cout << "[-] Error: 'e' and phi(n) are not coprime! Invalid RSA keys.\n";
                continue;
            }

            std::cout << "\n[+] RSA Key Pair Generated:\n";
            std::cout << "    Modulus (n)       = " << n << "\n";
            std::cout << "    Totient phi(n)    = " << phi << "\n";
            std::cout << "    Public Key (e, n) = (" << e << ", " << n << ")\n";
            std::cout << "    Private Key (d)   = " << d << "\n";

            uint64_t ciphertext = RSA::encrypt(message, e, n);
            uint64_t decrypted = RSA::decrypt(ciphertext, d, n);

            std::cout << "\n[+] Cryptographic Operations:\n";
            std::cout << "    Original Message  = " << message << "\n";
            std::cout << "    Encrypted Cipher  = " << ciphertext << "\n";
            std::cout << "    Decrypted Message = " << decrypted << "\n";
            std::cout << "====================================\n";
            continue;
        }

        // Factorization Choices (1 - 3)
        std::cout << "Enter a positive integer to factorize: ";
        if (!(std::cin >> target)) {
            std::cout << "Invalid number format.\n";
            std::cin.clear();
            std::cin.ignore(10000, '\n');
            continue;
        }

        if (target <= 1) {
            std::cout << "[-] Number must be greater than 1 to factorize.\n";
            continue;
        }

        std::pair<uint64_t, uint64_t> result = {0, 0};
        std::cout << "\n[*] Running factorization on " << target << "...\n";

        switch (choice) {
            case 1: {
                TrialDivision td;
                result = td.trialDivision(target);
                break;
            }
            case 2: {
                Fermat fermatSolver;
                result = fermatSolver.fermat(target);
                break;
            }
            case 3: {
                PollardsRho rhoSolver;
                result = rhoSolver.pollardsRho(target);
                break;
            }
        }

        std::cout << "------------------------------------\n";
        if (result.first != 0 && result.second != 0) {
            std::cout << "[+] Factors found successfully!\n";
            std::cout << "    Factor p: " << result.first << "\n";
            std::cout << "    Factor q: " << result.second << "\n";
            std::cout << "    Verification: " << result.first << " * " << result.second 
                      << " = " << (result.first * result.second) << "\n";
        } else {
            std::cout << "[-] Factorization failed or no non-trivial factors found.\n";
        }
        std::cout << "------------------------------------\n";
    }

    return 0;
}