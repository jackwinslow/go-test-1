# go-test-1

Please enter go run main.go to run the program. The first function executes as a part of Question 1, 
and then right after the second function executes as a part of Question 2. Neither Function incorporates working
functionality for arrays, only slices and maps where applicable.

Question 3 Answer:
I believe my code does resemble the theory about big O mentioned in the Golang Docs, since Stable makes O(n*log(n)*log(n))
to data.Swap instead of O(n*log(n)) like in Sort. This is because it is not altering the order of equal duplicates which
naturally saves time. O(n*log(n)*log(n)) < O(n*log(n)), so the theory is true.
