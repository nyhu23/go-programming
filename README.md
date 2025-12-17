
# gosort

**Name:** Nithya Santhosh  
**Student ID:** 241ADB038

## Run:
go mod init gosort
go run . -r 20        # random numbers
go run . -i incoming/input.txt # from file # used the same file from the folder
go run . -d incoming    # directory of .txt files

## Design:
- Splits into ceil(sqrt(n)) chunks, min 4
- Each chunk sorted in own goroutine
- Merges with pairwise merging (no re-sort)
- Random numbers: 0-999
- Directory mode creates folder with name-ID