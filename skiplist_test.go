package skiplist

import (
	"testing"
)

type TestType int

func (t TestType) Less(tt interface{}) bool {
	return t < tt.(TestType)
}

func (t TestType) Equal(tt interface{}) bool {
	return t == tt.(TestType)
}

func TestSkipListInsertSearch(t *testing.T) {
	s := NewSkipList()

	s.Insert(TestType(3))
	s.Insert(TestType(6))
	s.Insert(TestType(9))
	s.Insert(TestType(8))
	s.Insert(TestType(5))
	s.Insert(TestType(4))

	tests := []struct{
		TestInput TestType
		Success bool
	}{
		{
			TestType(3),
			true,
		},
		{
			TestType(6),
			true,
		},
		{
			TestType(9),
			true,
		},
		{
			TestType(8),
			true,
		},
		{
			TestType(7),
			false,
		},
		{
			TestType(2),
			false,
		},
	}

	for _, test := range tests {
		out := s.Search(test.TestInput)
		if out != test.Success {
			t.Errorf("Search failed for %d. Expected: %t, actual: %t",
				test.TestInput,
				test.Success,
				out)
		}
	}
}

func TestSkipListInsertDelete(t *testing.T) {
	s := NewSkipList()

	s.Insert(TestType(3))
	s.Insert(TestType(6))
	s.Insert(TestType(9))
	s.Insert(TestType(8))
	s.Insert(TestType(5))
	s.Insert(TestType(4))

	unknownNumber := 7

	tests := []struct{
		TestType
		Success bool
	}{
		{
			TestType(3),
			true,
		},
		{
			TestType(unknownNumber),
			false,
		},
		{
			TestType(6),
			true,
		},
		{
			TestType(9),
			true,
		},
		{
			TestType(3),
			false,
		},
		{
			TestType(8),
			true,
		},
		{
			TestType(5),
			true,
		},
		{
			TestType(4),
			true,
		},
		{
			TestType(4),
			false,
		},
	}

	for _, test := range tests {
		out := s.Delete(test.TestType)
		if test.Success != out {
			t.Errorf("Failed for %d. Expected: %t, actual: %t",
				test.TestType,
				test.Success,
				out)
		}
	}
}
