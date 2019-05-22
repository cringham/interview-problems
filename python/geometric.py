##### Finding Geometric Sequences

# Part A
# Given a list (L), positive integers sorted in ascending order
# And a ratio (R) >= 1
# Find the number of groups of 3 indices (i,j,k) in the list such that: 
# 1. i < j < k
# 2. {L[i], L[j], L[k]} is a geometric sequence with a common ratio r
#    i.e. R*L[i] == L[j], R*L[j] == L[k] or R^2 * L[i] == R * L[j] == L[k]

# Example:
# L = [1,1,5,25,25,125,625]
# R = 5

# triplets = [
#     (0, 2, 3)
#     (0, 2, 4),
#     (1, 2, 3),
#     (1, 2, 4),
#     (2, 3, 5),
#     (2, 4, 5),
#     (3, 5, 6),
#     (4, 5, 6)
# ]
# count = 8

# Part B
# How would your solution change if the list were unsorted? What if the ratio and list items did not need to be positive?
#  
# Given a list (L) of integers (unsorted, positive or negative)
# And a ratio (R) that is a non-zero real number:
# Find the same sets of indices as indicated in Part A
#
# You can use the second set of test cases to test a solution to this part.
# To test your solution with the unsorted test cases, add the unsorted=True argument to the case loop
# 
# for case in construct_test_cases(unsorted=True):
#     l, r, output = case
#     print(find_geo_seq(l, r) == output)

import random

def find_geo_seq(l, r):
    # code here
    return 0

def construct_test_cases(unsorted=False):
    """
    This function constructs the set of test cases used to test the geo_seq function.
    The initial set contains short, sorted lists, with the long case dynamically generated and added on.
    If the unsorted parameter is set to True, the test_cases will also contains unsorted lists (for Part B)
    """

    # initial sorted cases
    test_cases = [
        ([1, 2, 2, 4], 2, 2),
        ([1, 1, 5, 25, 25, 125, 625], 5, 8),
        ([125, 125, 25, 25, 5], 5, 0),
        ([1, 3, 9, 9, 9, 9, 9, 10, 27, 81], 3, 15),
        ([345]*10000, 1, 166616670000),
        ([1, 1, 1, 1] + [3, 3, 3, 3], 1, 8),
    ]

    # unsorted cases
    if unsorted:
        test_cases += [
        ([1, 3, 9, 9, 9, 9, 9, 10, 9, 27, 81], 3, 18),
        ([1, 3, 9, 9, 9, 4, 9, 9, 10, 12, 27, 81, 36, 81], 3, 21),
        ([4, 2, 1], .5, 1),
        ([1, 3, 9, 1, 3, 9], 3, 4),
        ([4, 2, 1, 4, 2, 1], .5, 4),
        ([1, 3, 9, 1, 1, 3, 9, 3], 3, 5),
        ([1, 3, 9, 1, 1, 3, 9, 3, 9], 3, 12),
        ([4, 2, 2, 1], .5, 2),
        ([1, -3, 9], -3, 1),
        ([1, -3, 9, 9, -27, 27], -3, 4),
        ([-4, 2, 2, -1], -.5, 2)
    ]

    # one long case, dynamically generated by incrementing powers of four and inserting a random number of each into the list
    long_case = []
    count = 0
    prev = 0
    prev_prev = 0
    for n in range(0, 10000):
        rand_count = random.randrange(4)
        count += prev_prev * prev * rand_count
        prev_prev = prev
        prev = rand_count
        for i in range(rand_count):
            long_case.append(4**n)

    test_cases.append((long_case, 4, count))

    return test_cases

for case in construct_test_cases():
    l, r, output = case
    print(find_geo_seq(l, r) == output)