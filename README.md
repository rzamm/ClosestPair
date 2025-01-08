# Closest Points Finder

This project is designed to find the closest pair of points from a given set of points in a plane.
It includes a Go program for the main functionality, and Python script for verifying correctness.

## Project Structure

- `main.go`: Contains the Go implementation of the closest points finder.
- `bug-finder`: Folder containing A Python script that runs the Go program and compares its output with a known correct solution.
- `test-files`: Files containing test cases for the Go program as well as the correct answer.

## Usage

1. **Run the Closest Points Finder**:
    ```sh
    go run main.go < test-files/cp*_input.txt
    ```

2. **Run the bug finder**:
    ```sh
   cd bug-finder
    ./run.sh
    ```

## Dependencies

- Go
- Python 3
- g++

## How The Bug Finder Works

1. The `run.sh` script builds the Go and C++ programs and runs `bugs-finder.py`.
2. `bug-finder.py` generates test cases and runs both the Go and C++ programs with the generated input.
3. The script then compares the outputs of the two programs to find any discrepancies.

## License

This project is licensed under the MIT License.