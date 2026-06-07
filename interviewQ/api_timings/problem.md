Title: Min, Max, and Mode of API Response Times

Problem statement:
Given a list of recorded API response times (non‑negative integers, in milliseconds), compute and return three statistics:
- Minimum response time
- Maximum response time
- Mode response time (the value that appears most frequently)

If the input list is empty, return a special indication (e.g., null or -1) for each statistic.

Tiebreaking rule for mode:
- If multiple values are tied for highest frequency, return the smallest value among them.

Input format (single test case):
- First line: n — an integer (0 ≤ n ≤ 10^6), the number of recorded response times.
- Second line: n space-separated non-negative integers t_i (0 ≤ t_i ≤ 10^9), the response times in milliseconds.

Output format:
- Three space-separated values: min max mode
- For an empty list (n = 0), output: -1 -1 -1

Examples:
1) Input:
6
120 200 120 150 200 120
Output:
120 200 120

2) Input:
5
50 50 100 100 75
Output:
50 100 50
(Explanation: 50 and 100 both appear twice; choose the smaller => 50)

3) Input:
0

Output:
-1 -1 -1

Constraints:
- Time limit: (typical) 1–2 seconds.
- Memory limit: (typical) 256–512 MB.