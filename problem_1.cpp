#include <iostream>

int problem1(unsigned int max) {
  const unsigned int n3 = (max - 1) / 3;
  const unsigned int n5 = (max - 1) / 5;
  const unsigned int n15 = (max - 1) / 15;
  return (
    3 * n3 * (n3 + 1)
    + 5 * n5 * (n5 + 1)
    - 15 * n15 * (n15 + 1)
  ) / 2;
}

int main() {
  std::cout << problem1(1000) << std::endl;
  return 0;
}
