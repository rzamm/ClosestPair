import math
import subprocess
import random
import sys

MIN_X = -100
MAX_X = 100
MIN_Y = -100
MAX_Y = 100

def generate_test_case(n):
    """Generate a single test case with n points."""
    points = [
        f"{random.uniform(MIN_X, MAX_X):.2f} {random.uniform(MIN_Y, MAX_Y):.2f}"
        for _ in range(n)
    ]
    return f"{n}\n" + "\n".join(points)

def generate_input(max_points):
    """Generate input with a given number of points."""
    case = []
    n = random.randint(2, max_points)
    case.append(generate_test_case(n))
    case.append("0")  # Terminating line
    return "\n".join(case)

def run_program(program, input_data):
    """Run a program with the given input data and return its output."""
    try:
        result = subprocess.run(
            program,
            input=input_data,
            text=True,
            capture_output=True
        )
        return result.stdout.strip()
    except Exception as e:
        print(f"Error running {program}: {e}")
        sys.exit(1)

def calculate_distance(input_str):
    # Split the input string into a list of numbers
    coords = list(map(float, input_str.split()))

    # Extract the coordinates for both points
    x1, y1, x2, y2 = coords

    # Calculate the Euclidean distance between the two points
    distance = math.sqrt((x2 - x1)**2 + (y2 - y1)**2)

    return distance

def find_bug():
    """Generate inputs and find a discrepancy between ours and the known solution."""
    for i in range(1000):
        input_data = generate_input(max_points=100)
        output1 = run_program("./our-solution", input_data)
        dist1 = calculate_distance(output1)
        output2 = run_program("./known-solution", input_data)
        dist2 = calculate_distance(output2)

        if i % 50 == 0:
            print(f"{i} cases checked")

        if dist1 != dist2:
            print("Bug found!")
            print("Input:")
            print(input_data)
            print("Output of our solution:")
            print(output1)
            print("Claimed distance:")
            print(f"{dist1:.2f}")
            print("Output of known solution:")
            print(output2)
            print("Actual distance:")
            print(f"{dist2:.2f}")
            break

    print("No Bugs found!")

if __name__ == "__main__":
    find_bug()
