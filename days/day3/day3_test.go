package day3

import (
	"reflect"
	"testing"
)

func TestMul(t *testing.T) {

	num1, num2 := 5, 4

	actual := mul(num1, num2)
	expected := 20

	if actual != expected {
		t.Errorf("Actual: %d is not equal to Expected: %d", actual, expected)
	}
}

func TestGetResult(t *testing.T) {
	cases := []struct {
		mulExpression string
		expected      int
	}{
		{
			mulExpression: "mul(5,5)",
			expected:      25,
		},
		{
			mulExpression: "mul(301,784)",
			expected:      235984,
		},
	}

	for _, v := range cases {
		actual := getResult(v.mulExpression)

		if actual != v.expected {
			t.Errorf("Actual: %d is not equal to Expected: %d", actual, v.expected)
		}
	}

}

func TestSplitByDont(t *testing.T) {
	cases := []struct {
		str      string
		expected []string
	}{
		{
			str:      "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			expected: []string{"xmul(2,4)&mul[3,7]!^", "_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"},
		}, {
			str:      "[#from())when()/}+%mul(982,733)mul(700,428)}}don't(){:,$+mul(395,45)[;don't()_mul(5,5)+mul(32,64]",
			expected: []string{"[#from())when()/}+%mul(982,733)mul(700,428)}}", "{:,$+mul(395,45)[;", "_mul(5,5)+mul(32,64]"},
		},
		{
			str:      "xmul(2,4)&mul[3,7]!^)_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			expected: []string{"xmul(2,4)&mul[3,7]!^)_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"},
		},
	}
	for _, v := range cases {
		actual := splitByDont(v.str)

		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Actual doesn't match expected\nActual:\t\t%v\nExpected:\t%v\n", actual, v.expected)
		}
	}
}

func TestGetMulEx(t *testing.T) {
	cases := []struct {
		str      string
		expected [][]string
	}{
		{
			str:      "xmul(2,4)&mul[3,7]!^",
			expected: [][]string{{"mul(2,4)"}},
		},
		{
			str:      "mul[3,7]!^_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			expected: [][]string{{"mul(5,5)"}, {"mul(11,8)"}, {"mul(8,5)"}},
		},
		{
			str:      "&mul[3,7]!^mul( 1, 1)+mul(32,64]do()?:mul$3,4%",
			expected: nil,
		},
	}

	for _, v := range cases {
		actual := getMulEx(v.str)

		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Actual doesn't match expected\nActual:\t\t%v\nExpected:\t%v\n", actual, v.expected)
		}
	}
}

func TestSplitByDo(t *testing.T) {
	cases := []struct {
		str      string
		expected []string
	}{
		{
			str:      "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			expected: []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)un", "?mul(8,5))"},
		},
		{
			str:      "xmul(2,4)&mul[3,7]!^don't()",
			expected: []string{"xmul(2,4)&mul[3,7]!^don't()"},
		},
		{
			str:      "undo()?mul(8,5))who()mul(846,294)mul(382,709)]do()mul(965,528)mul(89,614)",
			expected: []string{"un", "?mul(8,5))who()mul(846,294)mul(382,709)]", "mul(965,528)mul(89,614)"},
		},
	}

	for _, v := range cases {
		actual := splitByDo(v.str)

		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("Actual doesn't match expected\nActual:\t\t%v\nExpected:\t%v\n", actual, v.expected)
		}
	}
}

func TestSolvePartTwo(t *testing.T) {
	cases := []struct {
		str    string
		result int
	}{
		{
			str:    "xmul(2,4)%&mul[3,7]!",
			result: 8,
		}, {
			str:    "xmul(2,4)xdon't()mul(2,4)%&mul[3,7]!",
			result: 8,
		}, {
			str:    "don't()xmul(2,4)xdo()mul(2,4)%&mul[3,7]!don't()",
			result: 8,
		}, {
			str:    "don't()xmul(2,4)%&mul[3,7]!",
			result: 0,
		},
	}

	for _, v := range cases {
		actual, _ := SolvePartTwo(v.str)

		if actual != v.result {
			t.Errorf("Actual not equal to expected\nActual:\t\t%v\nExpected:\t%v\n%s\n", actual, v.result, v.str)
		}
	}
}
