
# Caesar Cipher

## Explain Problem

## Read inputs

```go
var length int
fmt.Scanf("%d\n", &length)
var input string
fmt.Scanf("%s\n", &input)
var rotate int
fmt.Scanf("%d\n", &rotate)
```

## Overall idea to solve this...

- Pretend no uppercase
- pretend no special characters

We want to find a characters spot in the alphabet string:

`abcdefghijklmnopqrstuvwxyz`

Then rotate it and start over at the start of the string if we get to the end.

```
z + 2 = ?

z, a, b
+0 +1 +2

z + 2 = b
```

## More generally...

We can rotate any letter in a slice by N regardless of what letter it is. Eg if our alphabet was only:

```go
alpha = "abcdf"
```

And we wanted to rotate each char in the string `bad` by `3` we get:

```
b + 3 = f
a + 3 = d
d + 3 = b // wraps around!
```

## Writing this in code

You will be tempted to treat your string as a byte slice

This will work in this case, but might not work all the time. Instead, let's treat it as a rune slice.

We can make a rune slice by just converting our string

```go
runes := []rune(someString)
```

This gives us code like:

```go
func rotateRune(s rune, delta int, key []rune) rune {
	idx := -1
	for i, r := range key {
		if s == r {
			idx = i
			break
		}
	}
	if idx < 0 {
		// shouldn't happen!
		panic("couldn't find that rune in the key")
	}
	for i := 0; i < delta; i++ {
		idx++
		if idx >= len(key) {
			idx = 0
		}
	}
	return key[idx]
}
```


## Explain modulo/remainder

5 / 4 = 1 in integer division
Number of WHOLE times numerator (5) goes into denominator (4)


Modulo gives remainder
5 % 4 = 1
6 % 4 = 2
7 % 4 = 3
The remainder AFTER dividing numerator into denominator


## Why does modulo matter?

We are cyclinng through our key a lot. Instead we can use modulo to simplify our code.


```go
func rotateRune(s rune, delta int, key []rune) rune {
	idx := strings.IndexRune(string(key), s)
	if idx < 0 {
		panic("couldn't find that rune in the key")
	}
	idx = (idx + delta) % len(key)
	return key[idx]
}
```

## Back to our original problem - rotating an entire string!

```go
lower = []rune("abcdefghijklmnopqrstuvwxyz")
upper = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// COULD ALSO USe upper/lower strings to find!
strings.IndexRune(r, string(upper)) >= 0
// a little slower, but more generic

func isUpperAlpha(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isLowerAlpha(r rune) bool {
	return r >= 'a' && r <= 'z'
}

// in main
ret := ""
for _, r := range input {
  switch {
  case isUpperAlpha(r):
    ret = ret + rotate(r, delta, upper)
  case isLowerAlpha(r):
    ret = ret + rotate(r, delta, lower)
  default:
    ret = ret + r
  }
}
fmt.Println(ret)
```

## Doing this without any slices

if our range is 80-106 we can't modulo correctly, so we need to rebase

```go
func rotateUpper(r rune, delta int) rune {
  // rebase
  tmp := int(r) - 'A'
  // add delta and modulo
	tmp = (tmp + delta) % 26
  // go back to original base
	return rune(tmp + 'A')
}
```
