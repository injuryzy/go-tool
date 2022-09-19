package timetool

import (
	"database/sql/driver"
	"fmt"
	"time"
)

//时区
var TimeLocal TimeZone

type TimeZone struct {
	//时区
	TimeZone string
	//时间偏移
	TimeOffset int
}

// 1. 创建 time.Time 类型的副本 XTime；
type MyTime struct {
	time.Time
}

// 2. 为 Xtime 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
func (t MyTime) MarshalJSON() ([]byte, error) {

	loc := time.FixedZone(TimeLocal.TimeZone, TimeLocal.TimeOffset)
	output := fmt.Sprintf("\"%s\"", t.In(loc).Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

// 3. 为 Xtime 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t MyTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// 4. 为 Xtime 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *MyTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = MyTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t *MyTime) String() string {
	loc := time.FixedZone(TimeLocal.TimeZone, TimeLocal.TimeOffset)
	return t.In(loc).String()
}

func (t *MyTime) GetTime() time.Time {
	loc := time.FixedZone(TimeLocal.TimeZone, TimeLocal.TimeOffset)
	return t.In(loc)
}
