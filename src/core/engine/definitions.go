package engine

type Bucket struct {
	CreateAt    string `json:"created_at" gorm:"column:created_at"`
	ClientIdent string `json:"client_ident" gorm:"column:client_ident"`
	Key         string `json:"key" gorm:"primaryKey;column:key"`
	Value       string `json:"value" gorm:"column:value"`
}
