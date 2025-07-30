package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type DBJSON []byte

// Scan 从数据库读取值 到 DBJSON 类型的字段, 接口 sql.Scanner
func (j *DBJSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("无效的数据源")
	}
	*j = append((*j)[0:0], s...)
	return nil
}

// Value 将结构体字段写入数据库 DBJSON , 接口 driver.Valuer
func (j DBJSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return string(j), nil
}

// MarshalJSON 转换为 DBJSON 输出, 接口 json.Marshaler
func (j DBJSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

// UnmarshalJSON 从 DBJSON 数据填充字段 字符串 -> 结构体 , 接口 json.Unmarshaler
func (j *DBJSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("异常 null")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

// Decode 反序列化到结构体
func (j *DBJSON) Decode(i interface{}) error {
	err := json.Unmarshal(*j, i)
	return err
}
