// Code generated by "stringer -type=Noun"; DO NOT EDIT.

package config

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Classifier-0]
	_ = x[Extractor-1]
}

const _Noun_name = "ClassifierExtractor"

var _Noun_index = [...]uint8{0, 10, 19}

func (i Noun) String() string {
	if i < 0 || i >= Noun(len(_Noun_index)-1) {
		return "Noun(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Noun_name[_Noun_index[i]:_Noun_index[i+1]]
}
