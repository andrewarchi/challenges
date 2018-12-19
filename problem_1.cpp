#include <iostream>

int problem1_s(unsigned int n, unsigned int d) {
  unsigned int m = (n - 1) / d;
  return d * m * (m + 1) / 2;
}

int problem1(unsigned int n) {
  return problem1_s(n, 3) + problem1_s(n, 5) - problem1_s(n, 15);
}

int main() {
  std::cout << problem1(1000) << std::endl;
  return 0;
}
