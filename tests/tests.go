package tests

import (
	"fmt"
	"github.com/whosonfirst/go-edtf"
	"time"
)

type TestResult struct {
	options TestResultOptions
}

type TestResultOptions struct {
	StartLowerTimeRFC3339 string
	StartUpperTimeRFC3339 string
	EndLowerTimeRFC3339   string
	EndUpperTimeRFC3339   string
	EndLowerTimeUnix      int64
	StartUpperTimeUnix    int64
	StartLowerTimeUnix    int64
	EndUpperTimeUnix      int64
}

func NewTestResult(opts TestResultOptions) *TestResult {

	r := &TestResult{
		options: opts,
	}

	return r
}

func (r *TestResult) TestDate(d *edtf.EDTFDate) error {

	/*

		if d.Start.Lower.Time != nil {
			fmt.Printf("[%s][start.lower] %s %d\n", d.String(), d.Start.Lower.Time.Format(time.RFC3339), d.Start.Lower.Time.Unix())
		}

		if d.Start.Upper.Time != nil {
			fmt.Printf("[%s][start.upper] %s %d\n", d.String(), d.Start.Lower.Time.Format(time.RFC3339), d.Start.Lower.Time.Unix())
		}

		if d.End.Lower.Time != nil {
			fmt.Printf("[%s][end.lower] %s %d\n", d.String(), d.End.Lower.Time.Format(time.RFC3339), d.End.Lower.Time.Unix())
		}

		if d.End.Upper.Time != nil {
			fmt.Printf("[%s][end.upper] %s %d\n", d.String(), d.End.Lower.Time.Format(time.RFC3339), d.End.Lower.Time.Unix())
		}

	*/

	err := r.testRFC3339All(d)

	if err != nil {
		return err
	}

	err = r.testUnixAll(d)

	if err != nil {
		return err
	}

	return nil
}

func (r *TestResult) testRFC3339All(d *edtf.EDTFDate) error {

	if r.options.StartLowerTimeRFC3339 != "" {

		err := r.testRFC3339(r.options.StartLowerTimeRFC3339, d.Start.Lower.Time)

		if err != nil {
			return fmt.Errorf("Failed StartLowerTimeRFC3339 test, %v", err)
		}
	}

	if r.options.StartUpperTimeRFC3339 != "" {

		err := r.testRFC3339(r.options.StartUpperTimeRFC3339, d.Start.Upper.Time)

		if err != nil {
			return fmt.Errorf("Failed StartUpperTimeRFC3339 test, %v", err)
		}
	}

	if r.options.EndLowerTimeRFC3339 != "" {

		err := r.testRFC3339(r.options.EndLowerTimeRFC3339, d.End.Lower.Time)

		if err != nil {
			return fmt.Errorf("Failed EndLowerTimeRFC3339 test, %v", err)
		}
	}

	if r.options.EndUpperTimeRFC3339 != "" {

		err := r.testRFC3339(r.options.EndUpperTimeRFC3339, d.End.Upper.Time)

		if err != nil {
			return fmt.Errorf("Failed EndUpperTimeRFC3339 test, %v", err)
		}
	}

	return nil
}

func (r *TestResult) testRFC3339(expected string, t *time.Time) error {

	if t == nil {
		return fmt.Errorf("Missing time.Time instance")
	}

	t_str := t.Format(time.RFC3339)

	if t_str != expected {
		return fmt.Errorf("Invalid RFC3339 time, expected '%s' but got '%s'", expected, t_str)
	}

	return nil
}

func (r *TestResult) testUnixAll(d *edtf.EDTFDate) error {

	if r.options.StartLowerTimeUnix != 0 {

		err := r.testUnix(r.options.StartLowerTimeUnix, d.Start.Lower.Time)

		if err != nil {
			return fmt.Errorf("Failed StartLowerTimeUnix test, %v", err)
		}
	}

	if r.options.StartUpperTimeUnix != 0 {

		err := r.testUnix(r.options.StartUpperTimeUnix, d.Start.Upper.Time)

		if err != nil {
			return fmt.Errorf("Failed StartUpperTimeUnix test, %v", err)
		}
	}

	if r.options.EndLowerTimeUnix != 0 {

		err := r.testUnix(r.options.EndLowerTimeUnix, d.End.Lower.Time)

		if err != nil {
			return fmt.Errorf("Failed EndLowerTimeUnix test, %v", err)
		}
	}

	if r.options.EndUpperTimeUnix != 0 {

		err := r.testUnix(r.options.EndUpperTimeUnix, d.End.Upper.Time)

		if err != nil {
			return fmt.Errorf("Failed EndUpperTimeUnix test, %v", err)
		}
	}

	return nil
}

func (r *TestResult) testUnix(expected int64, t *time.Time) error {

	if t == nil {
		return fmt.Errorf("Missing time.Time instance")
	}

	ts := t.Unix()

	if ts != expected {
		return fmt.Errorf("Invalid Unix time, expected '%d' but got '%d'", expected, ts)
	}

	return nil
}