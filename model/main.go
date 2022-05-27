package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	g "github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/olivere/elastic/v7"
	"net/http"
	"reflect"
	"reportapi/contrib/helper"
	"runtime"
	"strings"
	"time"
)

type MetaTable struct {
	MerchantRedis *redis.Client
	MerchantTD    *sqlx.DB
	DorisDB       *sqlx.DB
	PullDorisDB   *sqlx.DB
	ReportDB      *sqlx.DB
	SlaveDB       *sqlx.DB
	ReportEs      *elastic.Client
	Prefix        string
	EsPrefix      string
	PullPrefix    string
	Program       string
	Lang          string
}

var (
	loc               *time.Location
	meta              *MetaTable
	ctx               = context.Background()
	dialect           = g.Dialect("mysql")
	colReport         = helper.EnumFields(Report{})
	colRealTimeReport = helper.EnumFields(RealTimeReport{})
	accessCol         = helper.EnumFields(MonitorItem{})
	colsMember        = helper.EnumFields(Member{})
)

type AccessKey struct {
	ReportTime    int64  `json:"report_time" db:"report_time"`
	ReportTimeStr string `json:"report_time_str" db:"report_time_str"`
	Prefix        string `json:"prefix" db:"prefix"`
	LastHour      string `json:"last_hour" db:"last_hour"`
	Yesterday     string `json:"yesterday" db:"yesterday"`
}

func Constructor(mt *MetaTable) {

	meta = mt
	if meta.Lang == "cn" {
		loc, _ = time.LoadLocation("Asia/Shanghai")
	} else if meta.Lang == "vn" || meta.Lang == "th" {
		loc, _ = time.LoadLocation("Asia/Bangkok")
	}
}

func pushLog(err error, code string) error {

	_, file, line, _ := runtime.Caller(1)
	paths := strings.Split(file, "/")
	l := len(paths)
	if l > 2 {
		file = paths[l-2] + "/" + paths[l-1]
	}
	path := fmt.Sprintf("%s:%d", file, line)

	ts := time.Now()
	id := helper.GenId()

	fields := g.Record{
		"id":       id,
		"content":  err.Error(),
		"project":  meta.Program,
		"flags":    code,
		"filename": path,
		"ts":       ts.In(loc).UnixMilli(),
	}

	query, _, _ := dialect.Insert("goerror").Rows(&fields).ToSQL()
	//fmt.Println(query)
	_, err1 := meta.MerchantTD.Exec(query)
	if err1 != nil {
		fmt.Println("insert SMS = ", err1.Error(), fields)
	}

	note := fmt.Sprintf("Hệ thống lỗi %s", id)
	return errors.New(note)
}

func Close() {
	_ = meta.DorisDB.Close()
	_ = meta.PullDorisDB.Close()
	_ = meta.ReportDB.Close()
	_ = meta.SlaveDB.Close()
	_ = meta.MerchantRedis.Close()
}

func IsToday(s string, loc *time.Location) bool {

	t, err := time.Parse(http.TimeFormat, s)
	if err != nil {
		return false
	}

	y, m, d := t.In(loc).Date()
	ty, tm, td := time.Now().Date()

	if y == ty && m == tm && d == td {
		return true
	}

	return false
}

func EnumFields(obj interface{}) []interface{} {

	rt := reflect.TypeOf(obj)
	if rt.Kind() != reflect.Struct {
		return nil
	}

	var fields []interface{}
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		if field := f.Tag.Get("db"); field != "" && field != "-" {
			fields = append(fields, field)
		}
	}

	return fields
}

func MemberMCache(names []string) (map[string]Member, error) {

	data := map[string]Member{}

	if len(names) == 0 {
		return data, errors.New(helper.ParamNull)
	}

	var mbs []Member
	ex := g.Ex{
		"username": names,
	}
	query, _, _ := dialect.From("tbl_members").Where(ex).Select(colsMember...).Limit(uint(len(names))).ToSQL()
	err := meta.SlaveDB.Select(&mbs, query)
	if err != nil && err != sql.ErrNoRows {
		return data, pushLog(err, helper.DBErr)
	}

	if len(mbs) > 0 {
		for _, v := range mbs {
			if v.Username != "" {
				data[v.Username] = v
			}
		}
	}

	return data, nil
}

func MemberCache(id string) (Member, string, error) {

	var data Member

	if len(id) == 0 {
		return data, "", errors.New("id is null")
	}
	query := fmt.Sprintf(`select uid	,username,password,prefix,regip,reg_device,created_at,last_login_ip	,last_login_at,
source_id,first_bet_at,first_bet_amount,first_deposit_at,first_deposit_amount,top_uid,top_name,parent_uid,parent_name,bankcard_total,
last_login_device,last_login_source,remarks,balance,lock_amount,commission,state from tbl_members where uid = %s limit 1`, id)
	err := meta.SlaveDB.Get(&data, query)
	if err != nil {
		fmt.Println(id)
		return data, "db", err
	}

	return data, "", nil
}
