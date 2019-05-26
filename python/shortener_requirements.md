####### Shortener

Given two strings s, t, determine if string s can be changed into string t using the following rule:

Remove any number of lowercase letters from s (or none at all), then capitalize all remaining letters of s.

Example of the rule:

s = AAbcD

Possible outputs using the rule:
- Remove none, capitalize -> AABCD
- Remove c, capitalize -> AABD
- Remove b, capitalize -> AACD
- Remove b and c, capitalize -> AAD

If it is possible to create the string t by processing s using the rule, then the function should return True, otherwise return False.