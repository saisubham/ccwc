# ccwc

A word count utility similar to UNIX wc.
Done as part of this [Coding Challenge](https://codingchallenges.fyi/challenges/challenge-wc)

## Build
```shell
go build ccwc.go
```

## Run

### Count stats of a file given as arguments
```shell
./ccwc test.txt
./ccwc -w test.txt
./ccwc -l -c test.txt
```

### Count stats of a file given from stdin
```shell
cat test.txt | ./ccwc
cat test.txt | ./ccwc -l -w
echo "hello world" | ./ccwc
```

### Count stats of stdin
```shell
./ccwc # press ctrl+d to escape out of giving input
```

## Help
```shell
./ccwc -h
```