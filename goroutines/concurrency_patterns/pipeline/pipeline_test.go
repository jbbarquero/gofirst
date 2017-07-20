package pipeline

import "testing"

func TestLaunchPipeline(t *testing.T) {
	tableTest := [][]int{
		{3, 14},
		{5, 55},
	}

	var res int
	for _, test := range tableTest {
		res = LaunchPipeline(test[0])
		if res != test[1] {
			t.Fatalf("Should receive a %d but had a %d", test[1], res)
		}
		t.Logf("%d == %d", res, test[1])
	}

}
