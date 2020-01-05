// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	exampleIns := [][]string{{`"zzazz"`}, {`"mbadm"`}, {`"leetcode"`}, {`"g"`}, {`"no"`}}
	exampleOuts := [][]string{{`0`}, {`2`}, {`5`}, {`0`}, {`1`}}
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	if err := testutil.RunLeetCodeFunc(t, minInsertions, exampleIns, exampleOuts); err != nil {
		t.Fatal(err)
	}
}
