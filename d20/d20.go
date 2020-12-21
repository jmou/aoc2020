package main

import (
	"bufio"
	"fmt"
	"os"
)

const TileDimension = 12

const (
	BorderTop = iota
	BorderRight
	BorderBottom
	BorderLeft
)

type Tile struct {
	bitmap              [][]bool
	unnormalizedBorders [4]uint32
	borders             [4]uint32
}

func (t Tile) transformedBorder(transform uint8, cardinal int) uint32 {
	if transform&TransformCW90 != 0 {
		cardinal += 3
	}
	if transform&TransformCW180 != 0 {
		cardinal += 2
	}
	if transform&TransformFlipX != 0 && cardinal%2 == 0 {
		cardinal += 2
	}
	border := t.unnormalizedBorders[cardinal%4]
	if transform&TransformFlipX != 0 {
		return flip(border)
	}
	return border
}

func transformCoord(r, c int, transform uint8) (int, int) {
	if transform&TransformCW90 != 0 {
		r, c = 9-c, r
	}
	if transform&TransformCW180 != 0 {
		r, c = 9-r, 9-c
	}
	if transform&TransformFlipX != 0 {
		r = 9 - r
	}
	return r, c
}

func transformPatch(bitmap [][]bool, transform uint8) [][]bool {
	var patch [][]bool
	for r := 1; r < 9; r++ {
		var row []bool
		for c := 1; c < 9; c++ {
			tr, tc := transformCoord(r, c, transform)
			row = append(row, bitmap[tr][tc])
		}
		patch = append(patch, row)
	}
	return patch
}

const (
	TransformCW90  = 1 << 0
	TransformCW180 = 1 << 1
	TransformFlipX = 1 << 2
	TransformEnd   = 1 << 3
)

type Placement struct {
	id        int
	transform uint8
}

func bitsliceToUint32(slice []bool, reverse bool) uint32 {
	var bits uint32
	for i, value := range slice {
		if value {
			if reverse {
				bits |= 1 << (len(slice) - 1 - i)
			} else {
				bits |= 1 << i
			}
		}
	}
	return bits
}

func transposeColumn(bitmap [][]bool, column int, reverse bool) uint32 {
	var bits uint32
	for i := 0; i < len(bitmap[0]); i++ {
		row := i
		if reverse {
			row = len(bitmap[0]) - 1 - i
		}
		if bitmap[row][column] {
			bits |= 1 << i
		}
	}
	return bits
}

func flip(original uint32) uint32 {
	var reversed uint32
	for i := 0; i < 10; i++ {
		if original&(1<<i) != 0 {
			reversed |= 1 << (9 - i)
		}
	}
	return reversed
}

func normalizeFlip(original uint32) uint32 {
	reversed := flip(original)
	if reversed < original {
		return reversed
	}
	return original
}

func printBitmap(bitmap [][]bool) {
	for _, row := range bitmap {
		for _, bit := range row {
			if bit {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	tiles := make(map[int]Tile)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var id int
		fmt.Sscanf(scanner.Text(), "Tile %d:", &id)
		var tile Tile
		for scanner.Scan() {
			if scanner.Text() == "" {
				break
			}
			row := make([]bool, len(scanner.Text()))
			for i, char := range scanner.Text() {
				if char == '#' {
					row[i] = true
				}
			}
			tile.bitmap = append(tile.bitmap, row)
		}
		tile.unnormalizedBorders = [4]uint32{
			bitsliceToUint32(tile.bitmap[0], false),
			transposeColumn(tile.bitmap, len(tile.bitmap[0])-1, false),
			bitsliceToUint32(tile.bitmap[len(tile.bitmap)-1], true),
			transposeColumn(tile.bitmap, 0, true),
		}
		tile.borders = [4]uint32{
			normalizeFlip(tile.unnormalizedBorders[0]),
			normalizeFlip(tile.unnormalizedBorders[1]),
			normalizeFlip(tile.unnormalizedBorders[2]),
			normalizeFlip(tile.unnormalizedBorders[3]),
		}
		tiles[id] = tile
	}
	if scanner.Err() != nil {
		panic(scanner.Err)
	}

	tileBorders := make(map[uint32][]int)
	for id, tile := range tiles {
		for _, border := range tile.borders {
			tileBorders[border] = append(tileBorders[border], id)
		}
	}

	edges := make(map[int]bool)
	var corners []int
	for _, borders := range tileBorders {
		if len(borders) == 1 {
			if _, ok := edges[borders[0]]; ok {
				corners = append(corners, borders[0])
			} else {
				edges[borders[0]] = true
			}
		}
	}

	var image [TileDimension][TileDimension]Placement
	image[0][0] = Placement{corners[0], 0}
	placed := make(map[int]bool)
	placed[corners[0]] = true
	for y := 1; y < TileDimension; y++ {
		before := image[0][y-1]
		for _, border := range tiles[before.id].borders {
			if len(tileBorders[border]) > 1 {
				for _, id := range tileBorders[border] {
					if !placed[id] && edges[id] {
						image[0][y] = Placement{id, 0}
					}
				}
			}
		}
		placed[image[0][y].id] = true
	}
	for x := 1; x < TileDimension; x++ {
		before := image[x-1][0]
		for _, border := range tiles[before.id].borders {
			if len(tileBorders[border]) > 1 {
				for _, id := range tileBorders[border] {
					if !placed[id] && edges[id] {
						image[x][0] = Placement{id, 0}
						placed[id] = true
					}
				}
			}
		}
	}

	for x := 1; x < TileDimension; x++ {
		for y := 1; y < TileDimension; y++ {
			tileHit := make(map[int]bool)
			for _, border := range tiles[image[x-1][y].id].borders {
				for _, id := range tileBorders[border] {
					if !placed[id] && id != image[x-1][y].id {
						tileHit[id] = true
					}
				}
			}
			var found int
			for _, border := range tiles[image[x][y-1].id].borders {
				for _, id := range tileBorders[border] {
					if tileHit[id] && id != image[x][y-1].id {
						found = id
					}
				}
			}
			image[x][y] = Placement{found, 0}
			placed[found] = true
		}
	}

	fmt.Println(image[0][0].id * image[0][TileDimension-1].id *
		image[TileDimension-1][0].id * image[TileDimension-1][TileDimension-1].id)

	originBorders := tiles[image[0][0].id].borders
	var transformkey [4]int
	for i, border := range originBorders {
		for _, rightBorder := range tiles[image[0][1].id].borders {
			if border == rightBorder {
				transformkey[i] = 1
				break
			}
		}
		for _, downBorder := range tiles[image[1][0].id].borders {
			if border == downBorder {
				transformkey[i] = 2
				break
			}
		}
	}
	switch transformkey {
	case [4]int{0, 1, 2, 0}:
		image[0][0].transform = 0
	case [4]int{1, 2, 0, 0}:
		image[0][0].transform = TransformCW90
	case [4]int{2, 0, 0, 1}:
		image[0][0].transform = TransformCW180
	case [4]int{0, 0, 1, 2}:
		image[0][0].transform = TransformCW180 | TransformCW90
	case [4]int{2, 1, 0, 0}:
		image[0][0].transform = TransformFlipX
	case [4]int{0, 2, 1, 0}:
		image[0][0].transform = TransformFlipX | TransformCW90
	case [4]int{0, 0, 2, 1}:
		image[0][0].transform = TransformFlipX | TransformCW180
	case [4]int{1, 0, 0, 2}:
		image[0][0].transform = TransformFlipX | TransformCW180 | TransformCW90
	default:
		panic("could not determine origin transform")
	}

	for x := 0; x < TileDimension; x++ {
		for y := 0; y < TileDimension; y++ {
			if x == 0 && y == 0 {
				continue
			}
			var transform uint8
			for ; transform < TransformEnd; transform++ {
				borderTop := tiles[image[x][y].id].transformedBorder(transform, BorderTop)
				if x == 0 {
					if len(tileBorders[normalizeFlip(borderTop)]) > 1 {
						continue
					}
				} else {
					if borderTop != flip(tiles[image[x-1][y].id].transformedBorder(image[x-1][y].transform, BorderBottom)) {
						continue
					}
				}
				borderLeft := tiles[image[x][y].id].transformedBorder(transform, BorderLeft)
				if y == 0 {
					if len(tileBorders[normalizeFlip(borderLeft)]) > 1 {
						continue
					}
				} else {
					if borderLeft != flip(tiles[image[x][y-1].id].transformedBorder(image[x][y-1].transform, BorderRight)) {
						continue
					}
				}
				break
			}
			if transform == TransformEnd {
				panic("could not determine transform")
			}
			image[x][y].transform = transform
		}
	}

	var bitmap [][]bool
	for x := 0; x < TileDimension; x++ {
		for i := 0; i < 8; i++ {
			bitmap = append(bitmap, []bool{})
		}
		for y := 0; y < TileDimension; y++ {
			patch := transformPatch(tiles[image[x][y].id].bitmap, image[x][y].transform)
			for r, row := range patch {
				for _, bit := range row {
					bitmap[x*8+r] = append(bitmap[x*8+r], bit)
				}
			}
		}
	}

	seamonsters := make([][]bool, len(bitmap))
	for i, row := range bitmap {
		seamonsters[i] = make([]bool, len(row))
	}

	var seamonsterString = [3]string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}

	var transformed [][]bool
	for _, line := range seamonsterString {
		var row []bool
		for _, char := range line {
			if char == '#' {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		transformed = append(transformed, row)
	}
	for i := 0; i < 2; i++ {
		var flipped [][]bool
		for x := len(transformed) - 1; x >= 0; x-- {
			flipped = append(flipped, transformed[x])
		}
		transformed = flipped
		for j := 0; j < 4; j++ {
			rotated := make([][]bool, len(transformed[0]))
			for x, _ := range rotated {
				rotated[x] = make([]bool, len(transformed))
			}
			for x := 0; x < len(transformed); x++ {
				for y := 0; y < len(transformed[x]); y++ {
					rotated[y][len(transformed)-1-x] = transformed[x][y]
				}
			}
			transformed = rotated
			seamonsters = markSeamonsters(bitmap, seamonsters, transformed)
		}
	}

	originalBits, seamonsterBits := 0, 0
	for _, row := range bitmap {
		for _, bit := range row {
			if bit {
				originalBits++
			}
		}
	}
	for _, row := range seamonsters {
		for _, bit := range row {
			if bit {
				seamonsterBits++
			}
		}
	}
	fmt.Println(originalBits - seamonsterBits)
}

func markSeamonsters(bitmap, seamonsters, needle [][]bool) [][]bool {
	for x := 0; x < len(bitmap)-len(needle)+1; x++ {
		for y := 0; y < len(bitmap[x])-len(needle[0])+1; y++ {
			found := true
			// search:
			for r, line := range needle {
				for c, bit := range line {
					if bit && !bitmap[x+r][y+c] {
						found = false
						// break search
					}
				}
			}
			if found {
				for r, line := range needle {
					for c, bit := range line {
						if bit {
							seamonsters[x+r][y+c] = true
						}
					}
				}
			}
		}
	}
	return seamonsters
}
