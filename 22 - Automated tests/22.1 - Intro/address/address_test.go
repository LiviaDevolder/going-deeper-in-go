package address

import "testing"

type testCase struct {
	testData string
	expected string
}

func TestAddressType(t *testing.T) {
	t.Parallel()

	// Arrange
	testCases := []testCase{
		{"Route 66", "Route"},
		{"HIGHWAY TO HELL", "Highway"},
		{"Any Street", "Invalid Type"},
		{"", "Invalid Type"},
	}

	for _, data := range testCases {
		// Act
		result := AddressType(data.testData)

		// Assert
		if result != data.expected {
			t.Errorf("Error type is different from the expected. Expect %s but got %s", data.expected, result)
		}
	}
}
