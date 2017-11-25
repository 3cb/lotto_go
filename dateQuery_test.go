package main

import (
	"reflect"
	"testing"
)

func TestReformat(t *testing.T) {
	tt := []struct {
		queryIn, want DateQuery
	}{
		{DateQuery{[]string{"2016-03-18T04:00:00.000Z"}}, DateQuery{[]string{"2016-03-18T00:00:00.000"}}},
		{DateQuery{[]string{"2016-05-05T00:05:00.000Z", "2017-03-02T01:00:00.000Z"}}, DateQuery{[]string{"2016-05-05T00:00:00.000", "2017-03-02T00:00:00.000"}}},
		{DateQuery{[]string{"2016-07-22T04:00:00.111Z"}}, DateQuery{[]string{"2016-07-22T00:00:00.000"}}},
	}

	for _, c := range tt {
		(&c.queryIn).Reformat()
		got := c.queryIn
		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("(%q).Reformat() == %q. Wanted %q.", c.queryIn, got, c.want)
		}
	}
}
