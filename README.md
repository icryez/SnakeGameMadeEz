
# Snake Game

This project is a self-playing Snake game where the snake autonomously navigates the grid to locate and consume bait. Utilizing a concurrent search algorithm, the snake seeks out bait.


## Purpose
The primary goal of this project is to demonstrate the capabilities of concurrent programming in Go using Go routines and knowledge of data structures and algorithms. 

## Features

- Autonomous Gameplay: The snake operates independently, using a concurrent search algorithm to locate bait on the grid.

- Concurrent Search Algorithm: Utilizes go routines to search for bait, ensuring fast search performance.
- Dynamic Movement: Once the bait is found, the snake dynamically adjusts its movement to reach and consume the bait, utilizing pathfinding algorithms.
- Continuous Play: As the snake consumes bait, it continuously searches for the next appearing bait, ensuring seamless gameplay without interruptions.
- Scalability: The game is designed to handle varying levels of complexity and grid sizes, providing scalability and adaptability to different environments.


## Benefits

- Educational Value: Provides insights into concurrent programming concepts and algorithms through easy no brainer gameplay.

- Technical Showcase: Demonstrates the practical application of concurrent programming techniques in game development, highlighting their effectiveness in solving complex problems.
