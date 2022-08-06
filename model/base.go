package model

//const (
//	timeFormart = "2006-01-02 15:04:05"
//)

type BaseEntity struct {
	ID uint64 `gorm:"column:id;primary_key;"` // 主键id bigint
	//CreateTime time.Time `gorm:"column:create_time;"`    // 创建时间；必填 datetime
	//UpdateTime time.Time `gorm:"column:update_time;"`    // 修改时间；必填datetime
}
