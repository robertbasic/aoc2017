# AoC2017 in go

Bruteforcing my way through Advent of Code 2017 edition with golang.

I have no idea what am I doing.

## Notes

I'm stealing this idea of notes from [Luka](https://github.com/lmuzinic).


## Day 1

My first try at day 1 was a [bit of a mess](https://github.com/robertbasic/aoc2017/blob/54a103f5910eaa1e4d5558439914011ba5c24d49/day1.go).

On day 2 I rewrote day 1 to be a bit more nicer to read and added tests for it.

I also moved to reading the day inputs from files, using [`bufio.NewScanner`](https://github.com/robertbasic/aoc2017/blob/f6c810294e331c688c5401e0d15e3b4c7004c1d0/day1.go#L16).

Coming from PHP, juggling types around is [hard yo](https://github.com/robertbasic/aoc2017/blob/f6c810294e331c688c5401e0d15e3b4c7004c1d0/day1.go#L47-L50)! That's what I get for not understanding Go's type system!

## Day 2

I don't understand why do I need to re-read a file [after scanning it once](https://github.com/robertbasic/aoc2017/blob/f6c810294e331c688c5401e0d15e3b4c7004c1d0/day2.go#L13-L24)? Or is there maybe a way to "reset" a scanner?

There's a difference between a [`bufio.Scanner`](https://github.com/robertbasic/aoc2017/blob/f6c810294e331c688c5401e0d15e3b4c7004c1d0/day2.go#L27) and a `scanner.Scanner`. I first type-hinted the `scanner` against `scanner.Scanner`, took me a while to figure out why all of a sudden `scanner.Scan()` wanted to return 3 values instead of a `bool`.

Good thing about this [`bufio.NewScanner`](https://github.com/robertbasic/aoc2017/blob/f6c810294e331c688c5401e0d15e3b4c7004c1d0/day2_test.go#L35) is that I can pass in a different `io.Reader` - a string reader in tests, a file in actual code.

[`strconv.Atoi`](https://github.com/robertbasic/aoc2017/blob/f6c810294e331c688c5401e0d15e3b4c7004c1d0/day2.go#L58) seems like a nicer way to convert strings to integers.

[`strings.Fields`](https://github.com/robertbasic/aoc2017/blob/f6c810294e331c688c5401e0d15e3b4c7004c1d0/day2.go#L75) is nice, splits a string by whitespace.

So far the smallest thing I constantly trip up on is `=` vs `:=`. I'll learn it, eventually.

## Day 3

Puzzle 1 was tough as balls.

I cheated with puzzle 2 and solved it with [OEIS](https://oeis.org/A141481).

## Day 4

Nothing special about puzzle 1 for today. It's just wierd that Go has no "unique" function.

For puzzle 2, I first named the `resort` function as `sort`. That tripped up Go and couldn't
import the `sort` package.