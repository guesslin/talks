package main

import (
	"bytes"
	"fmt"
)

// start DGA OMIT
func generate(year, month, day int) string {
	var buffer bytes.Buffer
	for i := 0; i < 16; i++ {
		year = (((year ^ 8) * year) >> 11) ^ ((year & 0xFFFFFFF0) << 17)
		month = ((month ^ 4*month) >> 25) ^ 16*(month&0xFFFFFFF8)
		day = ((day ^ (day << 13)) >> 19) ^ ((day & 0xFFFFFFFE) << 12)
		buffer.WriteString(string(((year ^ month ^ day) % 25) + 97)) // HL
	}
	return buffer.String()
}

// end DGA OMIT

func main() {
	fmt.Println(generate(2015, 07, 20))
}
