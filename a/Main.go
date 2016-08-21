package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		c := map[byte]int{'N': 0, 'S': 0, 'W': 0, 'E': 0}
		bs := sc.Bytes()
		for _, b := range bs {
			num, ok := c[b]
			if ok {
				c[b] = num + 1
			}
		}
		n, _ := c['N']
		s, _ := c['S']
		w, _ := c['W']
		e, _ := c['E']

		if n == 0 && s == 0 && w == 0 && e == 0 {
			// ?
			fmt.Println("No")
		} else {
			if ((n == 0) && (s == 0) || (n > 0) && (s > 0)) &&
				((w == 0) && (e == 0) || (w > 0) && (e > 0)) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}
