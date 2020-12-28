package level1

import (
	"github.com/whosonfirst/go-edtf"
	"errors"
	"fmt"
	"strings"
	"strconv"
)

/*

Unspecified digit(s) from the right

The character 'X' may be used in place of one or more rightmost digits to indicate that the value of that digit is unspecified, for the following cases:

    A year with one or two (rightmost) unspecified digits in a year-only expression (year precision)
    Example 1       ‘201X’
    Example 2       ‘20XX’
    Year specified, month unspecified in a year-month expression (month precision)
    Example 3       ‘2004-XX’
    Year and month specified, day unspecified in a year-month-day expression (day precision)
    Example 4       ‘1985-04-XX’
    Year specified, day and month unspecified in a year-month-day expression  (day precision)
    Example 5       ‘1985-XX-XX’


*/

func IsUnspecifiedDigits(edtf_str string) bool {
	return re_unspecified.MatchString(edtf_str)
}

func ParseUnspecifiedDigits(edtf_str string) (*edtf.EDTFDate, error) {

	m := re_unspecified.FindStringSubmatch(edtf_str)

	if len(m) != 4 {
		return nil, errors.New("Invalid Level 1 unspecified digits string")
	}
	
	yyyy := m[1]
	mm := m[2]
	dd := m[3]

	start_yyyy := yyyy
	end_yyyy := yyyy
	
	/*
	start_mm := mm
	start_dd := dd


	end_mm := mm
	end_dd := dd
	*/
	
	if strings.HasSuffix(yyyy, "X"){

		start_m := int64(0)
		end_m := int64(0)
		
		start_c:= int64(0)
		end_c := int64(900)

		start_d := int64(0)
		end_d := int64(90)

		start_y := int64(0)
		end_y := int64(9)
		
		if string(yyyy[0]) == "X" {
			return nil, errors.New("X millenium not implemented")
		} else {
			
			m, err := strconv.ParseInt(string(yyyy[0]), 10, 32)

			if err != nil {
				return nil, err
			}

			start_m = m * 1000
			end_m = start_m
		}

		if string(yyyy[1]) != "X" {

			c, err := strconv.ParseInt(string(yyyy[1]), 10, 32)

			if err != nil {
				return nil, err
			}

			start_c = c * 100
			end_c = start_c
		}

		if string(yyyy[2]) != "X" {

			d, err := strconv.ParseInt(string(yyyy[2]), 10, 32)

			if err != nil {
				return nil, err
			}

			start_d = d * 10
			end_d = start_d
		}

		if string(yyyy[3]) != "X" {
			
			y, err := strconv.ParseInt(string(yyyy[3]), 10, 32)

			if err != nil {
				return nil, err
			}

			start_y = y * 1
			end_y = start_y
		}

		start_ymd := start_m + start_c + start_d + start_y
		end_ymd := end_m + end_c + end_d + end_y

		// fmt.Printf("OMG '%s' '%d' '%d' '%d' '%d' '%d'\n", yyyy, start_m, start_c, start_d, start_y, start_ymd)
		// fmt.Printf("WTF '%s' '%d' '%d' '%d' '%d' '%d'\n", yyyy, end_m, end_c, end_d, end_y, end_ymd)

		start_yyyy = strconv.FormatInt(start_ymd, 10)
		end_yyyy = strconv.FormatInt(end_ymd, 10)		
	}

	if strings.HasSuffix(mm, "X"){

		
	}

	if strings.HasSuffix(dd, "X"){

	}

	fmt.Println(start_yyyy, end_yyyy)
	return nil, nil
}
