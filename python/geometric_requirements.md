#### Finding Geometric Sequences

Part A
Given a list (L), positive integers sorted in ascending order
And a ratio (R) >= 1
Find the number of groups of 3 indices (i,j,k) in the list such that: 
1. i < j < k
2. {L[i], L[j], L[k]} is a geometric sequence with a common ratio r
   i.e. R*L[i] == L[j], R*L[j] == L[k] or R^2 * L[i] == R * L[j] == L[k]

Example:
L = [1,1,5,25,25,125,625]
R = 5

triplets = [
    (0, 2, 3)
    (0, 2, 4),
    (1, 2, 3),
    (1, 2, 4),
    (2, 3, 5),
    (2, 4, 5),
    (3, 5, 6),
    (4, 5, 6)
]
count = 8

Part B
How would your solution change if the list were unsorted? What if the ratio and list items did not need to be positive?

Given a list (L) of integers (unsorted, positive or negative)
And a ratio (R) that is a non-zero real number:
Find the same sets of indices as indicated in Part A

You can use the second set of test cases to test a solution to this part.
To test your solution with the unsorted test cases, add the unsorted=True argument to the case loop

for case in construct_test_cases(unsorted=True):
    l, r, output = case
    print(find_geo_seq(l, r) == output)