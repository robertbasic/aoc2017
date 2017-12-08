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

## Day 5

With puzzle 1, I first tried to use a `map` for `instructions`, realised that a map is unordered, so I had to
switch to a slice.

Then I kept comparing `pos > l` and wondering why I keep getting an index out of range error.

For puzzle 2, what caught me off guard is that the slices are "just pointers" to the underlying array. See this
[goplay](https://play.golang.org/p/AIzHYPB07F)

## Day 6

For puzzle 1, I made a mistake in generating the `k` for the `blocks` map. I used only the value from the slice,
which caused incorrect result in the end, as it breaks when, for example, we have 11 & 1, and 1 & 11.

For puzzle 2, I just dumped all the `k`s from the `blocks` map into a file, found the line number on which
the last line first appears/repeats, substracted it from the last line number, voila.

## Day 7

Solved the first puzzle with an "array diff". Built two slices, one with all the programs that are holding
other programs, and another slice with all the programs that are being held by other programs. The difference
between those two slices is the program that holds all the other programs.

The second puzzle was harder. I didn't even solve it properly.

Knowing the "root" program from the first puzzle, I built a tree of towers from the bottom up. Every tower
has the name, the weight, the carried weight and the total weight.

Then the idea was to go through this tree and find the out of balance tower. Ended up just dumping the whole
tower tree into a file and then visually look for the branches that are out of balance until I spotted the
one that has the wrong weight.

## Day 8

I was so happy when I read the first puzzle of the day, immediately had a solution in my head using golang
`struct`s, had the solution typed out soon, the test case was passing, ran the program over the puzzle input,
got the answer, aaaand. It's wrong. WTH?!

Well, I misunderstood the task in the first puzzle. I understood it as to find the largest value ever in the
registries, but the actual task is to find the largest value at the end of the instructions.

My original solution to puzzle 1 was the solution to puzzle 2, so... Yeah. Read and re-read, and re-read some
more, folks!