package MySQL

import (
	MySQlModels "github.com/XDcobra/go_license_key_api_template/model/MySQL"
	"gorm.io/gorm"
)

func Automigration(db *gorm.DB) error {
	// Auto-migrate User (and more models if added in the future)
	err := db.AutoMigrate(&MySQlModels.User{})

	return err
}
