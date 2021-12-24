# go-round
round float with interface like java
```
const (
	UP RoundMode = iota
	DOWN
	CEILING
	FLOOR
	HALF_UP
	HALF_DOWN
	HALF_EVEN
)
input := 1.2345
Round(input, 1, UP)
Round(input, 1, DOWN)
Round(input, 1, CEILING)
Round(input, 1, FLOOR)
Round(input, 1, HALF_UP)
Round(input, 1, HALF_DOWN)
Round(input, 1, HALF_EVEN)
```
