
# gosort

**Name:** Nithya Santhosh  
**Student ID:** 241ADB038

## How to Run

go mod init gosort

 1: Random Numbers

go run . -r 20 or go build
./gosort -r 20
Generates 15 random integers (0-999) and sorts them.

 2: Input File

go run . -i input.txt or ./gosort -i input.txt
Reads integers from a file .

 3: Directory

go run . -d incoming or ./gosort -d incoming
Sorts all .txt files in the folder, saves to incoming_sorted_nithya_santhosh_241ADB038.

## Design Decisions

Chunking: Minimum 4 chunks, otherwise ceil(sqrt(n)) chunks.

Concurrency: Each chunk sorted in separate goroutine using sync.WaitGroup.

No Race Conditions: Copies chunks before sorting to avoid conflicts.

Merging: Efficient pairwise merging (not flatten-and-resort).

Random Range: 0 to 999 inclusive.

Error Handling: Validates inputs and provides clear error messages.