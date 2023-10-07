# Coding Challenge: Vending Machine

## Objective:
Create a program that dispenses items based on the input nominal with specific denominations.

## Task Guidelines:
Imagine a vending machine with the following product catalog:
```
------------------------------
|  Item   |   Price   |
------------------------------
| Aqua    |   2000    |
| Sosro   |   5000    |
| Cola    |   7000    |
| Milo    |   9000    |
| Coffee  |  12000    |
------------------------------
```

The vending machine only accepts denominations of Rp. 5000 and Rp. 2000. Products are selected from the most expensive to the cheapest. If the inserted amount is not in the specified denominations (5000 & 2000), an 'invalid denomination' error will occur. Type 'exit' to exit the program.

## Test Cases:
- `[2000]` → Output: 1 Aqua
- `[2000, 2000]` → Output: 2 Aqua
- `[5000, 2000]` → Output: 1 Cola
- `[1000]` → Output: Invalid denomination
- `[5000, 5000]` → Output: 1 Milo
- `[5000, 5000, 5000, 2000]` → Output: 1 Coffee, 1 Sosro
