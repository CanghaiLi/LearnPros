package model

func migration() {
	DB.Set("gorm:table_options", "charset=utf8mb4")
	DB.AutoMigrate(&User{}).AutoMigrate(&Task{})
	DB.Model(&Task{}).AddForeignKey("Uid", "User(id)", "CASCADE", "CASCADE")
}
