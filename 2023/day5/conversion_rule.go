package main

type Range struct {
	start int
	end   int
}

func NewSeedRange(start, length int) Range {
	return Range{
		start: start,
		end:   start + (length - 1),
	}
}

type ConversionRule struct {
	input  Range
	output Range
}

func NewConversionRule(srcStart, dstStart, rangeLength int) ConversionRule {
	return ConversionRule{
		input: Range{
			start: srcStart,
			end:   srcStart + rangeLength - 1,
		},
		output: Range{
			start: dstStart,
			end:   dstStart + rangeLength - 1,
		},
	}
}

func (r ConversionRule) Convert(input int) int {
	if input < r.input.start || input > r.input.end {
		return input
	}
	return r.output.start + (input - r.input.start)
}

func (r ConversionRule) ConvertRanges(inputRanges []Range) []Range {
	var outputRanges []Range
	for _, input := range inputRanges {
		// skip ranges that are negative, because those are already transformed
		if input.start < 0 || input.end < 0 {
			outputRanges = append(outputRanges, input)
			continue
		}
		if !r.Overlaps(input) {
			outputRanges = append(outputRanges, input)
			continue
		}

		var ranges []Range
		if input.start < r.input.start {
			ranges = append(ranges, Range{
				start: input.start,
				end:   r.input.start - 1,
			})
		}

		start := max(input.start, r.input.start)
		end := min(input.end, r.input.end)
		// mark already transformed ranges as negative ranges
		ranges = append(ranges, Range{
			start: -(r.output.start + (start - r.input.start)),
			end:   -(r.output.start + (end - r.input.start)),
		})

		if input.end > r.input.end {
			ranges = append(ranges, Range{
				start: r.input.end + 1,
				end:   input.end,
			})
		}
		outputRanges = append(outputRanges, ranges...)
	}
	return outputRanges
}

func (r ConversionRule) Overlaps(input Range) bool {
	if input.start > r.input.end || r.input.start > input.end {
		return false
	}
	return true
}
