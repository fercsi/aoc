# `aoc` go module

Helper utility for adventofcode challenges

## Installation

```bash
go get github.com/fercsi/aoc
```

## Usage

Example of how to import and use it:

```go
import "github.com/fercsi/aoc"

func main() {
    aoc.ReadInputIntList()
    ...
}
```

Dijkstra module:

```go
import "github.com/fercsi/aoc/dijkstra"

func main() {
    graph := dijkstra.Graph{}
    ...
    distances := dijkstra.Dijkstra(graph, 0, len(Graph))
}
```



## Features

### Handling input data

- `func GetRunType() string`
- `func ReadInput() []string`
- `func ReadInputLine() string`
- `func ReadInputBytes() [][]byte`
- `func ReadInputSplit(delimiter string) [][]string`
- `func ReadInputLineSplit(delimiter string) []string`
- `func ReadInputInt() []int`
- `func ReadInputIntList() [][]int`
- `func ReadInputLineIntList() []int`
- `func ParseLineIntList(line string) []int`

### Show intermediate and final results

- `func ShowResult(values ...any)`
- `func ShowBytesCondensed(area [][]byte)`
- `func ShowIntsCondensed(area [][]int)`
- `func ShowIntsCsv(area [][]int)`

### Manage coordinates

- `type Coord struct`
- `func (coord Coord) Up() Coord`
- `func (coord Coord) Down() Coord`
- `func (coord Coord) Left() Coord`
- `func (coord Coord) Right() Coord`
- `func (coord Coord) Neighbours(width, height int) []Coord`
- `func (coord Coord) AllNeighbours() []Coord`
- `func (coord Coord) Neighbours8(width, height int) []Coord`
- `func (coord Coord) AllNeighbours8() []Coord`
- `func Distance(coord1, coord2 Coord) int`
- `func AtCoord[E any](area [][]E, coord Coord) E`
- `func AtCoordUnlimited[E any](area [][]E, coord Coord, defaultValue E) E`
- `func SetAtCoord[E any](area [][]E, coord Coord, val E)`
- `func Boundaries(coords []Coord) (Coord, Coord)`
- `func GridCoords(tl, br Coord) iter.Seq[Coord]`

### Save results as an image

- `func SaveBytesImage(filename string, image [][]byte, colors map[byte]int)`

### Further tools

- `func Sum[E addable](s []E) E`
- `func SumFunc[S any, E addable](s []S, f func(e S) E) E`
- `func SumSeq[E addable](seq iter.Seq[E]) E`
- `func CountIf[S any](s []S, f func(S) bool) int`
- `func Grid2D(startx, endx, starty, endy int) iter.Seq2[int, int]`

### Dijkstra algprithm

- `type Edge struct`
- `type Graph map[int][]Edge`
- `func Dijkstra(graph Graph, start int, n int) []int`

## Documentation

- [GoDoc reference](https://pkg.go.dev/github.com/fercsi/aoc)

## Testing

How to run tests:

```bash
go test
```

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
