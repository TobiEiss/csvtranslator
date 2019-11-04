package transformer

import (
	"strings"
)

// TransformPerRow apllies template per row
func TransformPerRow(source, template [][]string) [][]string {
	result := [][]string{}

	// create a map of source-header -> index
	sourceHeader := map[string]int{}
	for sourceHeaderIndex := 0; sourceHeaderIndex < len(source[0]); sourceHeaderIndex++ {
		sourceHeader[source[0][sourceHeaderIndex]] = sourceHeaderIndex
	}

	// add header
	result = append(result, template[0])

	// remove header of template to iterate over content
	template = append(template[:0], template[1:]...)

	// line by line of source
	for sourceLineIndex := 1; sourceLineIndex < len(source); sourceLineIndex++ {
		// line by line of template
		for templateLineIndex := 0; templateLineIndex < len(template); templateLineIndex++ {
			// make a copy of this line
			line := append([]string(nil), template[templateLineIndex]...)

			// iterate all fields
			for lineFieldIndex := 0; lineFieldIndex < len(line); lineFieldIndex++ {
				// replace value if there is a link to source
				if strings.HasPrefix(line[lineFieldIndex], "!") {
					key := strings.TrimPrefix(line[lineFieldIndex], "!")
					line[lineFieldIndex] = source[sourceLineIndex][sourceHeader[key]]
				}
			}

			// append line to result
			result = append(result, line)
		}
	}
	return result
}
