// Package aoc provides utility functions and types
// to assist with solving Advent of Code puzzles and
// other similar algorithmic challenges.
//
// It includes:
//   - Input reading and parsing utilities
//   - Simple output formatting (for 2D arrays)
//   - Coordinate manipulation (type Coord)
//   - Saving 2D byte arrays as images
//
// Example usage:
//
//  input := aoc.ReadInputInt()
//  sum := 0
//  for _, value := range input {
//      sum += value
//  }
//
//  coord := aoc.Coord{X: 0, Y: 0}
//  upper := coord.Up()
//
//  caveMap := aoc.ReadInputBytes()
//  aoc.SaveBytesImage("output.ppm", caveMap, map[byte]int{
//      '.': 0xFFFFFF,
//      '#': 0x000000,
//  })
//
// This package aims to provide fast and simple helpers
// for common Advent of Code tasks, minimizing boilerplate code.
package aoc
