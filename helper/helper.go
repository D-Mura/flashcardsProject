package helper

import (
	"time"
)

func MakeFileNamePrefix() string {

	// ファイル名のプレフィックスを作成
	t := time.Now()
	layout := "2006_01_02_15_04_05"
	return t.Format(layout) + "_"

}
