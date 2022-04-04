package memo

import "gorm.io/gorm"

type Memo struct {
	gorm.Model
	Text string
}
