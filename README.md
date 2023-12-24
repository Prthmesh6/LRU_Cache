# LRU Cache Implementation in Golang

Welcome to my LRU (Least Recently Used) Cache implementation in Golang! This project showcases the functionality of an LRU Cache utilizing generics in Golang.

## What is an LRU Cache?

An LRU Cache is a data structure that maintains a limited number of items, discarding the least recently used items when full. It operates on the principle that items least recently accessed are the ones most likely to be removed when the cache reaches its capacity.

## Features

- **LRU Cache:** Implements a generic LRU Cache that stores key-value pairs.
- **Golang with Generics:** Utilizes the power of generics in Golang to create a versatile and type-safe cache.

## Advantages of Using Generics in Golang for LRU Cache Implementation

### 1. Type Safety and Reusability

- Generics allow the creation of a cache that can handle any data type without sacrificing type safety. This results in reusable and adaptable code.

### 2. Reduced Code Duplication

- With generics, it's possible to write a single implementation for the cache logic, avoiding redundant code for different data types.

### 3. Improved Readability and Maintainability

- Generic code tends to be more readable as it eliminates the need for repetitive code structures, making maintenance and understanding easier.

### Prerequisites

- Golang version 1.18 and later

### Getting Started

1. Clone this repository.
2. You can play with this Cache from main.go file itself.

### Using the LRU Cache
- You can generate lru1,lru2 and so on by using Cache.New(), Shown example in main.go
- If you make any changes, make sure to run or add more unit tests
- To instantiate you need to define 2 things,
1) What are the data types of your key-value (can be any data type)
2) Maximum Capacity of LRU Cache


## Future Enhancements

- Will add more cache types like Least Frequently Used (LFU) cache
- This code is scalable to integrate with any number of different caches, will cover them


## Contact Information

For any inquiries or suggestions, feel free to contact me at,

<div id="badges">
  <a href="https://www.linkedin.com/in/prathmeshpatil64/">
    <img src="https://img.shields.io/badge/LinkedIn-blue?style=for-the-badge&logo=linkedin&logoColor=white" alt="LinkedIn Badge"/>
  </a>
</div>
