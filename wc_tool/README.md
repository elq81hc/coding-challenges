# [Write Your Own wc Tool](https://codingchallenges.fyi/challenges/challenge-wc)
This is a Go solution to the challenge.
```bash
Usage: ccwc [OPTION]... [FILE]...
  -c    The number of bytes in each input file is written to the standard output
  -l    The number of lines in each input file is written to the standard output
  -m    The number of characters in each input file is written to the standard output
  -w    The number of words in each input file is written to the standard output
```

# How to build
To build the application, run the following command:

```bash
$ make build
go build -o bin/ccwc main.go
```
This will create a binary named `ccwc` in the bin directory.

# How to run
```bash
# The default behavior is to print the number of lines, words, and bytes of the input file.
$ ./ccwc test.txt
  7143  58164  335040  test.txt

# To get only the number of lines
$ ./ccwc -l test.txt
  7143  test.txt
```

# How to run tests
```bash
$ make test
=== RUN   TestWcTestData
--- PASS: TestWcTestData (0.00s)
=== RUN   TestWcFile
--- PASS: TestWcFile (0.01s)
PASS
ok      _/Users/admin/learns/coding-challenges/wc_tool  0.401s
```

# Benchmark
```bash
$ make benchmark
BenchmarkWc-8            3087156               358.4 ns/op          4096 B/op          1 allocs/op
PASS
ok      _/Users/admin/learns/coding-challenges/wc_tool  1.858s
```