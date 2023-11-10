package testify

import (
	"testing"

	"github.com/denist-huma/testify/v2/suite"
)

type Suite struct{}

func (s *Suite) TearDownSuite(t *suite.T) {
	t.Log(">> suite tear down")
}

func (s *Suite) TearDownTest(t *suite.T) {
	t.Log(">> single test tear down")
}

func (s *Suite) TestOne(t *suite.T) {
	for _, v := range []string{"sub1", "sub2", "sub3"} {
		t.Run(v, func(t *suite.T) {
			t.Parallel()
		})
	}
}

func (s *Suite) TestTwo(t *suite.T) {
	for _, v := range []string{"sub1", "sub2", "sub3"} {
		t.Run(v, func(t *suite.T) {
			t.Parallel()
		})
	}
}

// TestSuiteParallelSubTests shows that the issue with the parallel test is solved.
// https://github.com/stretchr/testify/issues/934
func TestSuiteParallelSubTests(t *testing.T) {
	suite.Run(t, &Suite{})
	// Output:
	// === RUN   TestLogic
	// === RUN   TestLogic/All
	// === RUN   TestLogic/All/TestOne
	// === RUN   TestLogic/All/TestOne/sub1
	// === PAUSE TestLogic/All/TestOne/sub1
	// === RUN   TestLogic/All/TestOne/sub2
	// === PAUSE TestLogic/All/TestOne/sub2
	// === RUN   TestLogic/All/TestOne/sub3
	// === PAUSE TestLogic/All/TestOne/sub3
	// === NAME  TestLogic/All/TestOne
	//     suite.go:64: >> single test tear down
	// === CONT  TestLogic/All/TestOne/sub1
	// === CONT  TestLogic/All/TestOne/sub3
	// === CONT  TestLogic/All/TestOne/sub2
	// === RUN   TestLogic/All/TestTwo
	// === RUN   TestLogic/All/TestTwo/sub1
	// === PAUSE TestLogic/All/TestTwo/sub1
	// === RUN   TestLogic/All/TestTwo/sub2
	// === PAUSE TestLogic/All/TestTwo/sub2
	// === RUN   TestLogic/All/TestTwo/sub3
	// === PAUSE TestLogic/All/TestTwo/sub3
	// === NAME  TestLogic/All/TestTwo
	//     suite.go:64: >> single test tear down
	// === CONT  TestLogic/All/TestTwo/sub1
	// === CONT  TestLogic/All/TestTwo/sub2
	// === CONT  TestLogic/All/TestTwo/sub3
	// === NAME  TestLogic
	//     suite.go:64: >> suite tear down
	// --- PASS: TestLogic (0.00s)
	//     --- PASS: TestLogic/All (0.00s)
	//         --- PASS: TestLogic/All/TestOne (0.00s)
	//             --- PASS: TestLogic/All/TestOne/sub1 (0.00s)
	//             --- PASS: TestLogic/All/TestOne/sub3 (0.00s)
	//             --- PASS: TestLogic/All/TestOne/sub2 (0.00s)
	//         --- PASS: TestLogic/All/TestTwo (0.00s)
	//             --- PASS: TestLogic/All/TestTwo/sub1 (0.00s)
	//             --- PASS: TestLogic/All/TestTwo/sub2 (0.00s)
	//             --- PASS: TestLogic/All/TestTwo/sub3 (0.00s)
	// PASS
}
