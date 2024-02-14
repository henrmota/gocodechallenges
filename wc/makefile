assert1: 
	wc -c test.txt 

assert2:
	wc -l test.txt

assert3:
	wc -w test.txt

assert4:
	wc -m test.txt
	
assert5: 
	wc test.txt

assert6: 
	cat test.txt | wc test.txt

test1: assert1
	go run ccwc.go -c test.txt 
	
test2: assert2
	go run ccwc.go -l test.txt

test3: assert3
	go run ccwc.go -w test.txt

test4: assert4
	go run ccwc.go -m test.txt

test5: assert5
	go run ccwc.go test.txt

test6: assert6
	cat test.txt | go run ccwc.go test.txt


all: test1 test2 test3 test4 test5 test6

build: 
	go build -o bin/ccwc ccwc.go

run:
	go run ccwc.go
