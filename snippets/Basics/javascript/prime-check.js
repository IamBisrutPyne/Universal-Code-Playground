// Prime Check in JavaScript
// This program checks whether a number is prime

// Function to check if a number is prime
function isPrime(n) {
    if (n <= 1) return false;       // 0 and 1 are not prime
    if (n <= 3) return true;        // 2 and 3 are prime

    if (n % 2 === 0 || n % 3 === 0) return false;

    // Check divisibility using 6k ± 1 optimization
    for (let i = 5; i * i <= n; i += 6) {
        if (n % i === 0 || n % (i + 2) === 0) return false;
    }

    return true;
}

// Example usage:
const n= prompt("Enter a number to check if it's prime: ");
const num = parseInt(n);
if (isNaN(number)) {
    console.log("Please enter a valid number!");
} else {
    console.log(`${number} is prime? ${isPrime(number)}`);
}