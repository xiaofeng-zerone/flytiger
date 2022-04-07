package ddosmechanism

type Hello struct {
	Name   string `json:"name" gorm:"comment:名称"`   //名称
	Remark string `json:"remark" gorm:"comment:备注"` //备注
}
