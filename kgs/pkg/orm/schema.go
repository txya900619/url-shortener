package orm

type UnusedKey struct {
	Key string `gorm:"primaryKey"`
}

type UsedKey struct {
	Key string `gorm:"primaryKey"`
}
