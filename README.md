# AoC2017 in go

Bruteforcing my way through Advent of Code 2017 edition with golang.

I have no idea what am I doing.

## Notes

I'm stealing this idea of notes from [Luka](https://github.com/lmuzinic).

After day 13 I went back to day 1 and started optimizing the code.

## Day 1

My first try at day 1 was a [bit of a mess](https://github.com/robertbasic/aoc2017/blob/54a103f5910eaa1e4d5558439914011ba5c24d49/day1.go).

On day 2 I rewrote day 1 to be a bit more nicer to read and added tests for it.

I also moved to reading the day inputs from files, using [`bufio.NewScanner`](https://github.com/robertbasic/aoc2017/blob/f6c810294e331c688c5401e0d15e3b4c7004c1d0/day1.go#L16).

Coming from PHP, juggling types around is [hard yo](https://github.com/robertbasic/aoc2017/blob/f6c810294e331c688c5401e0d15e3b4c7004c1d0/day1.go#L47-L50)! That's what I get for not understanding Go's type system!

### Optimizations

 - using `ioutil.ReadFile` to read in the entire file instead of openning it, scanning with a scanner, creating an input var, passing that around.
 - moved away from strings and converting to integers to do the sum calculations to bytes. Subtracting 48 from the byte to get the actual integer value.
 - benchmarking reports no difference between using `int`s and `int64`, so I left those as `int64`
 - I also see no difference between looping over the `input` bytes with `range` or with a "regular" `for` loop.

```
BenchmarkDay1-4      	   10000	    149501 ns/op	   22660 B/op	    4086 allocs/op
BenchmarkDay1Opt-4   	  200000	      9665 ns/op	    3104 B/op	      10 allocs/op
```

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

A day later, I want back to the 2nd puzzle and managed to solve it completely by code. The solution is not 100%
correct, as it is a bit too "eager", as in it fixes the weight by the required weight difference on **all** of
the towers on the unbalanced branch, but in the end it finds the final unbalanced tower and correctly adjusts
it's weight to give the correct answer, so I can live with that.

## Day 8

I was so happy when I read the first puzzle of the day, immediately had a solution in my head using golang
`struct`s, had the solution typed out soon, the test case was passing, ran the program over the puzzle input,
got the answer, aaaand. It's wrong. WTH?!

Well, I misunderstood the task in the first puzzle. I understood it as to find the largest value ever in the
registries, but the actual task is to find the largest value at the end of the instructions.

My original solution to puzzle 1 was the solution to puzzle 2, so... Yeah. Read and re-read, and re-read some
more, folks!

## Day 9

Going step by step for puzzle 1 was the right way to go:

 - cancel out stuff
 - remove garbage
 - count score

Most problem had with cancelling out stuff, because I kept looking back in the input string, instead of in the
output string.

Turns out going step by step proved to be really helpful for the 2nd puzzle, because I already had the function
for cancelling out stuff, which doesn't count towards the total garbage removed. The rest was counting the number
of characters between the opening `<` and closing `>`.

## Day 10

Ugh, this was a tough one. Again. I guess I'm not good with numbers.

For the first puzzle, my first and biggest mistake is that I kept trying to build a new list out of the existing
list. This ended up in weird bugs as I mixed golang's `append` with index access, as the new list kept growing
beyond the size of the original list. Eventually I deleted all the code I wrote on Sunday and left it for Monday.

On Monday I started from scratch and instead went with replacing items in the existing list. That quickly fell
into place once I fixed all the out of index errors:

 - instead of `sublist := list[position : position+length]` I had `sublist := list[position : length]` which was
   good for the test input, but was broken on the real input,
 - instead fo `for i := position; i < length; i++ {` I had `for i := position; i < position+length; i++ {` which
   made `i` go out of bounds pretty quick. Again, was OK for test input, not so OK for real input.

OK, even though the resulting checksum in puzzle 1 was correct, the resulting list was bad. Pure luck I guess that
the first two elements were what they should have been.

Need to modify a copy of the original list. On Sunday I was making a completely new list, but the real thing is to
work on the copy of the original list. I should probably look into the real reason why, but I'm so tired from this
puzzle, to be honest.

Second puzzle was bad because I HAD THE WRONG VARIABLE AS AN ARGUMENT TO `ElvenHashChecksum`!!! After converting the
lengths to their ASCII counterpart, I ended up still having the original `lengths` variable as an argument. FFS.

**NOTE**: On what is technically day 15 I was traveling to Zagreb, so in that time I rewrote day 10 to be more
manageable, easier to reuse for day 14.

## Day 11

I tried first to figure out movement on my own... Didn't quite get it. I still have another idea I want to try, but
until then I solved it using information from [this AMAZING article on hex grids](https://www.redblobgames.com/grids/hexagons/).
h/t [@lmuzinic](https://github.com/lmuzinic) for finding that article.

## Day 12

Took me a little while to figure out how to recurse here... Biggest problem was that I was setting `seen[to]` in `BuildPointers`
outside/after the main loop, which meant the `seen` was pretty much useless. Once I fixed that, the rest fell into place.

## Day 13

Biggest challenge today was to come up with a solution for `move` that does not include `for` loops
and handles the delay correctly. I don't even know how to explain the math behind it:

``` golang
func (s *Scanner) move(cp int) {
	r := s.rang - 1
	d := cp - r
	div := int(math.Trunc(float64(d / r)))
	m := int(math.Mod(float64(d), float64(r)))

	if math.Mod(float64(div), 2) == 0 {
		s.position = r - m
	} else {
		s.position = m
	}
}
```

Boils down to how many rounds does a scanner do moving around in it's range for the given `cp` 
number of steps. It probably can be improved further, but it gave me the correct answer in under
a minute, so I'm good with that.

Oh, and notice that the range for depth 1 is always 2. That helps in speeding up things.

@lmuzinic had a good optimization idea - when looking for the delay, no need to everytime a scanner catches us, it's enough
to be caught only for the first time!

With all this in place the execution is pretty fast:

```
time ./aoc2017
Severity is:  1476
Delay for:  3937334

real    0m3.146s
user    0m3.139s
sys     0m0.003s
```

## Day 14

Holy shit is the description of today's puzzle confusing. I need to go back and rewrite day 10 to make day 14 nicer.

3 days later... (OK, I was traveling on days 15 and 16). After rewriting the day 10 puzzles, the first puzzle for day 14
was easy. Apparently day 12 can be used for day 14 puzzle 2, but I ended up writing a new solution for the 2nd puzzle
anyway.

For puzzle 2 my first attempt consisted of just simply going one step back and one step up from the current position. That
didn't work, gave a number to high. Then I realised I was missing an edge case of one step forward one step up. But then
that made me realise I was missing an even larger edge case. I ended up drawing a 8x8 grid of ones and zeros that covers
all the possible cases (it's not that hard really), and using that grid solving that puzzle became a lot easier:

``` golang
binrows := []string{
	"01001011",
	"11011011",
	"01010001",
	"01110011",
	"00000110",
	"11100111",
	"00100000",
	"01111111",
}
```

## Day 15

Pretty much bruteforced myself out of today's puzzles. Also, solved on day 17.

## Day 16

Oh, this was fun! The programs rearrange themselves to a..p after 60 iterations, so I only
had to do 1billion mod 60 = 40 iterations to find the correct answer for the 2nd puzzle.

## Day 17

Puzzle 1 taught me this weird slice trick, which I think I understand

``` golang
func insertAt(s []int, p int, v int) []int {
	s = append(s, 0)
	copy(s[p+1:], s[p:])
	s[p] = v
	return s
}
```

It appends a single zero to the end, copies from the position `p` to the end to the position `p+1` to the end, and
sets the value at the position `p` to be `v`. All this while leaving anything below position `p` intact.

Puzzle 2 boils down to track the number that would be added after the zero number, which is always at zero position.
Took me a while to figure that out, but here we are!

## Day 18

Lagging behind with days 16 and 17, so jumping straight to day 18.

Even though I'd like to say the first puzzle was easy, I won't. First, I don't like to say "easy" or "simple" about
any programming task... Second, I made some silly mistakes.

Biggest mistake was with `jgz` is that I tested against the `reg`, and then jumped by it too.

Then I learned that to `break` multiple levels in Go I need to use labels. `break` breaks only one level, but to break
2 or more levels, we need to set labels and then `break foo` to get to the foo label.