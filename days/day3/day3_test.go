package day3

import "testing"

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
