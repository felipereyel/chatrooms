package models

import "testing"

// func IsCommand(message string) bool {
// 	return message[0] == '/'
// }

// func (c *CommandView) GetStockCode() (string, error) {
// 	regex := regexp.MustCompile(`^\/stock=([a-zA-Z0-9.]+)$`)
// 	matches := regex.FindStringSubmatch(c.Payload)
// 	if len(matches) != 2 {
// 		return "", errors.New("invalid command")
// 	}

// 	return matches[1], nil
// }

var isCommandTestCases = []struct {
	message  string
	expected bool
}{
	{"/stock=GOOG", true},
	{"stock=GOOG", false},
	{"", false},
}

func TestIsCommand(t *testing.T) {
	for _, testCase := range isCommandTestCases {
		actual := IsCommand(testCase.message)
		if actual != testCase.expected {
			t.Fatalf("Expected %v, got %v", testCase.expected, actual)
		}
	}
}

var getStockCodeTestCases = []struct {
	payload  string
	expected string
	err      error
}{
	{"/stock=GOOG", "GOOG", nil},
	{"/stock=", "", ErrInvalidCommand},
	{"", "", ErrInvalidCommand},
}

func TestGetStockCode(t *testing.T) {
	for _, testCase := range getStockCodeTestCases {
		command := CommandView{Payload: testCase.payload}

		stockCode, err := command.GetStockCode()
		if err != testCase.err {
			t.Fatalf("Expected %v, got %v", testCase.err, err)
		}

		if stockCode != testCase.expected {
			t.Fatalf("Expected %v, got %v", testCase.expected, stockCode)
		}
	}
}
