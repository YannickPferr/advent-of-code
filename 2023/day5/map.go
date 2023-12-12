package main

type Map struct {
	conversionRules []ConversionRule
}

func (m *Map) AddRule(conversionRule ConversionRule) {
	m.conversionRules = append(m.conversionRules, conversionRule)
}

func (m *Map) Convert(input int) int {
	for _, conversionRule := range m.conversionRules {
		if convertedValue := conversionRule.Convert(input); convertedValue != input {
			return convertedValue
		}
	}
	return input
}

func (m *Map) ConvertRange(inputRanges []Range) []Range {
	var outputRanges []Range
	for _, conversionRule := range m.conversionRules {
		inputRanges = conversionRule.ConvertRanges(inputRanges)
	}
	for _, inputRange := range inputRanges {
		multiplier := 1
		// we marked ranges that are already transformed as negatives, so now we need to unmark them
		if inputRange.start < 0 || inputRange.end < 0 {
			multiplier = -1
		}
		outputRanges = append(outputRanges, Range{
			start: inputRange.start * multiplier,
			end:   inputRange.end * multiplier,
		})
	}
	return outputRanges
}

//25 95
//18 + (81 - 25)
