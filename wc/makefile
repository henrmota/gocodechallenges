test1: 
	go run . -c test.txt
	@echo "----"
	wc -c test.txt
	@echo "##################\n"
test2:
	go run . -l test.txt
	@echo "----"
	wc -l test.txt
	@echo "##################\n"
test3:
	go run . -w test.txt
	@echo "----"
	wc -w test.txt
	@echo "##################\n"
test4:
	go run . -m test.txt
	@echo "----"
	wc -m test.txt
	@echo "##################\n"

test5:
	go run . test.txt
	@echo "----"
	wc test.txt
	@echo "##################\n"
test6:
	cat test.txt | go run .
	@echo "----"
	cat test.txt | wc
	@echo "##################\n"

testAll: test1 test2 test3 test4 test5 test6

build: 
	go build -o bin/ccwc main.go

run:
	go run main.go
