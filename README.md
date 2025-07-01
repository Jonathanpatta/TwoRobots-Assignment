# Game Simulation Project

This project simulates a game scenario as outlined in the assignment specifications. It is structured into modular components to separate concerns and improve maintainability.

## Project Structure

- **main.go**  
  Contains the simulation logic that runs the game scenario provided in the assignment.

- **game_engine.go**  
  Includes the game engine code responsible for managing core execution flow and integration of game components.

- **game_logic.go**  
  Contains the core game logic, such as rules, ability handling, and game state transitions. This helps isolate pure game logic from engine mechanics.

- **game_engine_test.go**  
  Includes unit tests for the `ExecuteAbility` and `ProcessEvent` functions to ensure correctness and robustness.

## Future Improvements

- **Improved Initialization**  
  Enhance the initialization process for the game and game state to make writing and running tests easier and more flexible.

- **Ability Priority Execution**  
  Implement a priority system for ability execution, allowing abilities to be ranked and resolved in a defined order.
