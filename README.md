# [freeCodeCamp.org](https://www.freecodecamp.org/) | Coderbyte - Curso de programação dinâmica - Implementação em Golang

<img src="freecodecamp-logo.jpg" alt="freecodecamp-logo" width="48"/>

Esse repositório possui todo o conteudo do curso **Dynamic Programming - Learn to Solve Algorithmic Problems & Coding Challenges** implementado em **Go 1.16**.

Além disso, possui também a página **[webassembly.html](https://guilhermerodrigues680.github.io/freeCodeCamp-coderbyte-dynamic-programming-course-golang-implementation/webassembly.html)** que executa as versões das implementações em Go compiladas para módulos WebAssembly direto do navegador de internet.


## 📚 Link do curso no YouTube 📚

Dynamic Programming - Learn to Solve Algorithmic Problems & Coding Challenges:
[www.youtube.com/watch?v=oBt53YbR9Kk](https://www.youtube.com/watch?v=oBt53YbR9Kk)

## ⭐️ Implementações ⭐️
TODO: Fazer todas implementações
TODO: Criar links para descrições e soluções
TODO: Renomear normal para `brute force`
TODO: Escrever complexidade `Big O` de `tempo` e `espaço`
- 💻 fib 
    - **[Descrição do problema](#-fib)**
    - [normal](./01-fib/main.go)
    - [memoization](./01A-fib-memoization/main.go)
    - tabulation

- 💻 gridTraveler
    - **[Descrição do problema](#-gridTraveler)**
    - [normal](./02-gridTraveler/main.go)
    - [memoization](./02A-gridTraveler-memoization/main.go)
    - tabulation

- 💻 canSum
    - **[Descrição do problema](#-canSum)**
    - [normal](./03-canSum/main.go)
    - [memoization](./03A-canSum-memoization/main.go)
    - tabulation

- 💻 howSum
    - **[Descrição do problema](#-howSum)**
    - [normal](./04-howSum/main.go)
    - [memoization](./04A-howSum-memoization/main.go)
    - tabulation

- 💻 bestSum
    - **[Descrição do problema](#-bestSum)**
    - [normal](./05-bestSum/main.go)
    - [memoization](./05A-bestSum-memoization/main.go)
    - tabulation

- 💻 canConstruct
    - **[Descrição do problema](#-canConstruct)**
    - [normal](./06-canConstruct/main.go)
    - [memoization](./06A-canConstruct-memoization/main.go)
    - tabulation

- 💻 countConstruct
    - **[Descrição do problema](#-countConstruct)**
    - [normal](./07-countConstruct/main.go)
    - [memoization](./07A-countConstruct-memoization/main.go)
    - tabulation

- 💻 allConstruct
    - **[Descrição do problema](./)**
    - normal
    - memoization
    - tabulation

## ⭐️ Descrição dos problemas ⭐️
TODO: Preencher descrição dos problemas
### 💻 fib
Write a function `fib(n)` that takes in a number as an argument.
The function should return the n-th number of the Fibonacci sequence.

The 1st and 2nd number of the sequence is 1.
To generate the next number of the sequence, we sum the previous two.

```txt
n:      1, 2, 3, 4, 5, 6, 7,  8,  9,  ...
fib(n): 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
```

### 💻 gridTraveler
Say that you are a traveler on a 2D grid. You begin in the top-left corner and your goal is to travel to bottom-right corner.
You may only move down or right.

In how many ways can you travel to the goal on a grid with dimensions m * n?

Write a function `gridTraveler(m, n)` that calculates this.

### 💻 canSum
Write a function `canSum(targetSum, numbers)` that takes in a targetSum and an array of numbers as arguments.

The function should return a boolean indicating whether or not it is possible to generate the targetSum using numbers fro the array.

You may use an element of the array as many times as needed.

You way assume that all input numbers are nonnegative.

### 💻 howSum
Write a function `howSum(targetSum, numbers)` that takes in a targetSum and an array of numbers as arguments.

The function should return an array containing any combination of elements that add up to exactly the targetSum. If there is no combination that adds up to the targetSum, the return null.

If there are multiple combinations possible, you may return any single one.

### 💻 bestSum
Write a function `bestSum(targetSum, numbers)` that takes in a targetSum and an array of numbers as arguments.

The function should return an array containing the `shortest` combination of numbers that add up to exactly the targetSum.

If there is a tie for the shortest combination, you may return any one of the shortest.

### 💻 canConstruct
Write a function `canConstruct(target, wordBank)` that accepts a target string and an array of strings.

The function should return a boolean indicating whether or not the `target` can be constructed by concatenating elements of the `wordBank` array.

You may reuse elements of `wordBank` as many times as needed.

### 💻 countConstruct
Write a function `countConstruct(target, wordBank)` that accepts a target string and an array of strings.

The function should return the `number of ways` that the `target` can be constructed by concatenating elements of the `wordBank` array.

You may reuse elements of `wordBank` as many times as needed.

### 💻 allConstruct
