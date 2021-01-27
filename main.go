package main

import (
	"fmt"
	"hello/library"
	"time"
)

type VideoCall struct {
	VideoId           string    `gorm:"column:video_id" json:"video_id"`
	PlayerId          int64     `gorm:"column:req_player_id" json:"req_player_id"`               // 呼叫人
	AnchorId          int64     `gorm:"column:anchor_id" json:"anchor_id"`                       // 主播（同时也是房间id）
	BossId            int64     `gorm:"column:boss_id" json:"boss_id"`                           // 老板
	Charge            int32     `gorm:"column:charge" json:"charge"`                             // 主播价格
	SharePercent      int32     `gorm:"column:share_percent" json:"share_percent"`               // 主播分成比例
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`                     //
	UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at"`                     //
	BossVersion       string    `gorm:"column:boss_version" json:"boss_version"`                 // 老板客戶端版本
	BossDeviceType    int32     `gorm:"column:boss_device_type" json:"boss_device_type"`         // 老板客户端类型
	AnchorVersion     string    `gorm:"column:anchor_version" json:"anchor_version"`             // 主播客戶端版本
	AnchorDeviceType  int32     `gorm:"column:anchor_device_type" json:"anchor_device_type"`     // 主播客戶端类型
	AcceptAt          time.Time `gorm:"column:accept_at" json:"accept_at"`                       // 接听时间
	BossConfirmAt     time.Time `gorm:"column:boss_confirm_at" json:"boss_confirm_at"`           // 老板确认时间
	BossEnterRoomAt   time.Time `gorm:"column:boss_enter_room_at" json:"boss_enter_room_at"`     // 老板进房时间
	BossExitRoomAt    time.Time `gorm:"column:boss_exit_room_at" json:"boss_exit_room_at"`       // 老板退房时间
	BossExitTimeout   bool      `gorm:"column:boss_exit_timeout" json:"boss_exit_timeout"`       // 老板退房是否超时
	AnchorConfirmAt   time.Time `gorm:"column:anchor_confirm_at" json:"anchor_confirm_at"`       // 主播确认时间
	AnchorEnterRoomAt time.Time `gorm:"column:anchor_enter_room_at" json:"anchor_enter_room_at"` // 主播进房时间
	AnchorExitRoomAt  time.Time `gorm:"column:anchor_exit_room_at" json:"anchor_exit_room_at"`   // 主播退房时间
	AnchorExitTimeout bool      `gorm:"column:anchor_exit_timeout" json:"anchor_exit_timeout"`   // 主播退房是否超时
	IMCall            bool      `gorm:"column:im_call" json:"im_call"`                           // 是否使用im呼叫
	OptID             int64     `gorm:"column:opt_id" json:"-"`                                  // 触发下一步骤的一方id (不使用)
	Valid             bool      `gorm:"column:valid" json:"valid"`
}

func (f *VideoCall) TableName() string {
	return "video_calls"
}

func main() {
	target := 6
	arr := []interface{}{}
	arr = append(arr, 1, 2, 3, 4)
	fmt.Println(library.In_array(target, arr))
}
