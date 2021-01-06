package level2

import (
	"github.com/whosonfirst/go-edtf/tests"
)

var Tests map[string]map[string]*tests.TestResult = map[string]map[string]*tests.TestResult{
	EXPONENTIAL_YEAR: map[string]*tests.TestResult{
		"Y-17E7": tests.NewTestResult(tests.TestResultOptions{}), // TO DO - https://github.com/whosonfirst/go-edtf/issues/5
		"Y10E7":  tests.NewTestResult(tests.TestResultOptions{}), // TO DO
		"Y20E2": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2000-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "2000-01-01T23:59:59Z",
			EndLowerTimeRFC3339:   "2000-12-31T00:00:00Z",
			EndUpperTimeRFC3339:   "2000-12-31T23:59:59Z",
		}),
	},
	SIGNIFICANT_DIGITS: map[string]*tests.TestResult{
		"1950S2": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1900-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "1900-12-31T23:59:59Z",
			EndLowerTimeRFC3339:   "1999-01-01T00:00:00Z",
			EndUpperTimeRFC3339:   "1999-12-31T23:59:59Z",
		}),
		"Y171010000S3": tests.NewTestResult(tests.TestResultOptions{}),
		// -2000/-2999
		"Y-20E2S3": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "-2999-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "-2999-12-31T23:59:59Z",
			EndLowerTimeRFC3339:   "-2000-01-01T00:00:00Z",
			EndUpperTimeRFC3339:   "-2000-12-31T23:59:59Z",
		}),
		"Y3388E2S3": tests.NewTestResult(tests.TestResultOptions{}),
	},
	SUB_YEAR_GROUPINGS: map[string]*tests.TestResult{
		"2001-34": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2001-04-01T00:00:00Z",
			StartUpperTimeRFC3339: "2001-04-01T23:59:59Z",
			EndLowerTimeRFC3339:   "2001-06-30T00:00:00Z",
			EndUpperTimeRFC3339:   "2001-06-30T23:59:59Z",
		}),
		"2019-28": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2019-12-01T00:00:00Z",
			StartUpperTimeRFC3339: "2019-12-01T23:59:59Z",
			EndLowerTimeRFC3339:   "2020-02-01T00:00:00Z",
			EndUpperTimeRFC3339:   "2020-02-29T23:59:59Z",
		}),
		// "second quarter of 2001": tests.NewTestResult(tests.TestResultOptions{}),	// TO DO
	},
	SET_REPRESENTATIONS: map[string]*tests.TestResult{
		"[1667,1668,1670..1672]": tests.NewTestResult(tests.TestResultOptions{
			// THIS FEELS WRONG...LIKE IT'S BACKWARDS
			StartLowerTimeRFC3339: "1667-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "1667-12-31T23:59:59Z",
			EndLowerTimeRFC3339:   "1672-01-01T00:00:00Z",
			EndUpperTimeRFC3339:   "1672-12-31T23:59:59Z",
			// INCLUSIVITY
		}),
		"[..1760-12-03]": tests.NewTestResult(tests.TestResultOptions{
			EndLowerTimeRFC3339: "1760-12-03T00:00:00Z",
			EndUpperTimeRFC3339: "1760-12-03T23:59:59Z",
			StartLowerIsOpen:    true,
			StartUpperIsOpen:    true,
			// INCLUSIVITY
		}),
		"[1760-12..]": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1760-12-01T00:00:00Z",
			StartUpperTimeRFC3339: "1760-12-31T23:59:59Z",
			EndLowerIsOpen:        true,
			EndUpperIsOpen:        true,
			// INCLUSIVITY
		}),

		/*

		FIX ME:
		POSSIBLE [1760-01,1760-02,1760-12..] 1760-01,1760-02,1760-12
		START '1760-01' END '1760-12'
		OPEN START 'false' END 'true'
		INCLUSIVITY 4
		PARSE 1760-01/..
		    set_representation_test.go:38: Results failed tests '[1760-01,1760-02,1760-12..]', Failed StartUpperTimeRFC3339 test, Invalid RFC3339 time, expected '1760-12-31T23:59:59Z' but got '1760-01-31T23:59:59Z'

				"[1760-01,1760-02,1760-12..]": tests.NewTestResult(tests.TestResultOptions{
					StartLowerTimeRFC3339: "1760-01-01T00:00:00Z",
					StartUpperTimeRFC3339: "1760-12-31T23:59:59Z",
					EndLowerIsOpen: true,
					EndUpperIsOpen: true,
					// INCLUSIVITY
				}),
		*/

		"[1667,1760-12]": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1667-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "1667-12-31T23:59:59Z",
			EndLowerTimeRFC3339:   "1760-12-01T00:00:00Z",
			EndUpperTimeRFC3339:   "1760-12-31T23:59:59Z",
			// INCLUSIVITY
		}),

		"[..1984]": tests.NewTestResult(tests.TestResultOptions{
			StartLowerIsOpen:    true,
			StartUpperIsOpen:    true,
			EndLowerTimeRFC3339: "1984-01-01T00:00:00Z",
			EndUpperTimeRFC3339: "1984-12-31T23:59:59Z",
			// INCLUSIVITY
		}),
		"{1667,1668,1670..1672}": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1667-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "1667-12-31T23:59:59Z",
			EndLowerTimeRFC3339:   "1672-01-01T00:00:00Z",
			EndUpperTimeRFC3339:   "1672-12-31T23:59:59Z",
			// INCLUSIVITY
		}),
		"{1960,1961-12}": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1960-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "1960-12-31T23:59:59Z",
			EndLowerTimeRFC3339:   "1961-12-01T00:00:00Z",
			EndUpperTimeRFC3339:   "1961-12-31T23:59:59Z",
			// INCLUSIVITY
		}),
		"{..1984}": tests.NewTestResult(tests.TestResultOptions{
			StartLowerIsOpen:    true,
			StartUpperIsOpen:    true,
			EndLowerTimeRFC3339: "1984-01-01T00:00:00Z",
			EndUpperTimeRFC3339: "1984-12-31T23:59:59Z",
			// INCLUSIVITY
		}),
	},
	GROUP_QUALIFICATION: map[string]*tests.TestResult{
		"2004-06-11%": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"2004-06~-11": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"2004?-06-11": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
	},
	INDIVIDUAL_QUALIFICATION: map[string]*tests.TestResult{
		"?2004-06-~11": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"2004-%06-11": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
	},
	UNSPECIFIED_DIGIT: map[string]*tests.TestResult{
		"156X-12-25": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"15XX-12-25": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		// "XXXX-12-XX": tests.NewTestResult(tests.TestResultOptions{}),	// TO DO
		"1XXX-XX": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"1XXX-12": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"1984-1X": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
	},
	INTERVAL: map[string]*tests.TestResult{
		"2004-06-~01/2004-06-~20": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"2004-06-XX/2004-07-03": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
	},
}
