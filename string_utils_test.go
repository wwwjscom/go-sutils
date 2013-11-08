package sutils

import "testing"

type testcase struct {
    test string
    expected string
}

var testcases []testcase
//var mysql_conn *Mysql

func init() {
    //mysql_conn = new(Mysql)
    testcases = []testcase{
        {
            "this is a simple test",
            "this is a simple test",
        },
        {
            "this test is a tad bit harder 'yo",
            "this test is a tad bit harder 'yo",
        },
        {
            "this test is \"real hard",
            "this test is \\\"real hard",
        },
    }
}

func TestEscapeQuotes(t *testing.T) {
    for _, tc := range testcases {
        res := EscapeAllQuotes(tc.test)
        if tc.expected != res {
            t.Errorf("\"%s\" was expected, but \"%s\" was returned", tc.expected, res)
        }
    }
}
