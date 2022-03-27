package model

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"strconv"
	"time"
)

type total struct {
	ReportTime string `db:"report_time" json:"report_time"`
	Prefix     string `db:"prefix" json:"prefix"`
	Count      uint64 `db:"count" json:"count"`
}

type User struct {
	ReportTime string `db:"report_time" json:"report_time"`
	Prefix     string `db:"prefix" json:"prefix"`
	Uid        string `db:"uid" json:"uid"`
}

func esQuery(rangeParam map[string][]interface{}, param map[string]interface{}) *elastic.BoolQuery {

	boolQuery := elastic.NewBoolQuery()
	terms := make([]elastic.Query, 0)
	filters := make([]elastic.Query, 0)

	if len(rangeParam) > 0 {
		for k, v := range rangeParam {
			if v == nil {
				continue
			}

			if len(v) == 2 {
				filters = append(filters, elastic.NewRangeQuery(k).Gte(v[0]).Lte(v[1]))
			}
		}
	}

	if len(param) > 0 {
		for k, v := range param {
			if v == nil {
				continue
			}

			if vv, ok := v.([]interface{}); ok {
				filters = append(filters, elastic.NewTermsQuery(k, vv...))
				continue
			}

			terms = append(terms, elastic.NewTermQuery(k, v))
		}
	}

	boolQuery.Filter(filters...)
	boolQuery.Must(terms...)

	return boolQuery
}

// parseTime format实时投注监控时间格式
// 例： 1617951600 -> 2021-04-09 15~16
func parseTime(s string) string {

	t, _ := strconv.ParseInt(s, 10, 64)
	ts := time.Unix(t, 0).In(loc)

	return fmt.Sprintf("%s %d~%d", ts.Format("2006-01-02"), ts.Hour(), ts.Add(60*time.Minute).Hour())
}

func parseDay(s string) string {

	t, _ := strconv.ParseInt(s, 10, 64)
	ts := time.Unix(t, 0).In(loc)

	return fmt.Sprintf("%s", ts.Format("2006-01-02"))
}

func parseMonth(s string) string {

	t, _ := strconv.ParseInt(s, 10, 64)
	ts := time.Unix(t, 0).In(loc)

	return fmt.Sprintf("%s", ts.Format("2006-01"))
}

func parsePart(startTime, endTime int64, flag string) string {

	ts := time.Unix(startTime, 0)

	te := time.Unix(endTime, 0).In(loc)

	result := ""
	if flag == "m" {
		result = fmt.Sprintf("%s ~ %s", ts.Format("2006-01"), te.Format("01"))
	}

	if flag == "d" {
		result = fmt.Sprintf("%s ~ %s", ts.Format("2006-01-02"), te.Format("01-02"))
	}

	return result
}

func parseHourTime(s string) string {

	t, _ := strconv.ParseInt(s, 10, 64)
	ts := time.Unix(t, 0).In(loc)

	return fmt.Sprintf("%d~%d", ts.Hour(), ts.Add(60*time.Minute).Hour())
}
