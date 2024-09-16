# Having GO fun with Katas

What the hell, the goal here is to rely on a TDD approach, Test First code later

ok went crazy with Unit tests using idea of `Generators` in Go Testing

# Instructions
Write a function that takes positive integers and outputs their string representation.

Your function should comply with the following additional rules:

If the number is a multiple of three, return the string "Fizz".
If the number is a multiple of five, return the string "Buzz".
If the number is a multiple of both three and five, return the string "FizzBuzz".
For example, given the numbers from 1 to 15 in order, the function would return:

```text
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
```

Reference: https://beta.cyber-dojo.org/creator/home

# Tesing with GoConvey

Check https://github.com/smartystreets/goconvey for how to

To install into new project

```shell
$ go get -u github.com/smartystreets/goconvey
```

When working locally with Tool (Live TDD Testing)

```shell
$ go install github.com/smartystreets/goconvey
```

For Windows (ensure GOLANG bin in path) - else just check `goconvey` site

```shell
goconvey.exe
```

Open Browser and point to: http://locahost:8080