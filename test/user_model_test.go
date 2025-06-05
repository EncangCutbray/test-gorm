package test_test

import (
	"cutbray/test-gorm/db"
	"cutbray/test-gorm/models"
	"cutbray/test-gorm/request"
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestUserModel(t *testing.T) {

	db.InitDB()

	db.DB.DisableForeignKeyConstraintWhenMigrating = true

	err := db.DB.Migrator().DropTable(&models.UserModel{}, &models.NoteModel{}, &models.CreditModel{})

	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	db.InitDB()
	db.DB.DisableForeignKeyConstraintWhenMigrating = false

	t.Run("create a new user", func(t *testing.T) {

		validate := validator.New(validator.WithRequiredStructEnabled())

		input := request.CreateUserRequest{
			Username: "Encang Cutbray", Email: "cabe2cabean214@gmail.com", Password: "secret214",
		}

		err := validate.Struct(input)

		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			for _, ve := range validationErrors {
				fmt.Printf("Field '%s' failed on '%s' tag\n", ve.Field(), ve.Tag())
			}
		}

		result := db.DB.Create(input.ToUserModel())

		assert.Nil(t, result.Error)
	})

	t.Run("create multiple user", func(t *testing.T) {

		validate := validator.New(validator.WithRequiredStructEnabled())

		var validUsers []*models.UserModel

		inputs := []request.CreateUserRequest{
			{Username: "Jinzhu", Email: "cabe2cabean214@1gmail.com", Password: "secret214"},
			{Username: "Jinzhu", Email: "cabe2cabean214@2gmail.com", Password: "secret214"},
			{Username: "Jackson", Email: "cabe2cabean214", Password: "secret214"},
		}

		for _, input := range inputs {
			err := validate.Struct(input)

			if err != nil {
				validationErrors := err.(validator.ValidationErrors)
				for _, ve := range validationErrors {
					fmt.Printf("Field '%s' failed on '%s' tag\n", ve.Field(), ve.Tag())
				}
				continue
			}

			validUsers = append(validUsers, input.ToUserModel())
		}

		result := db.DB.Create(validUsers)

		assert.EqualValues(t, result.RowsAffected, len(validUsers))

		assert.Nil(t, result.Error)
	})

	t.Run("found user by id", func(t *testing.T) {
		var user models.UserModel

		result := db.DB.Where("id = ?", "1").First(&user)

		assert.Nil(t, result.Error)
		assert.NotNil(t, user)
	})

	t.Run("update user by id", func(t *testing.T) {
		var user models.UserModel

		updateResult := db.DB.Model(&user).
			Where("id = ?", 1).
			Updates(map[string]any{
				"username": "hello",
				"password": "myPassword",
			})

		assert.Nil(t, updateResult.Error)
		assert.NotNil(t, user)

	})

	t.Run("soft delete user by id", func(t *testing.T) {
		var user models.UserModel

		// deletedUser := db.DB.Model(&user).
		// 	Where("id = ?", 2).
		// 	Delete(&user)

		deletedUser := db.DB.
			Where("id = ?", 2).
			Delete(&user)

		db.DB.Unscoped().Find(&user, 2)

		fmt.Println(user.Username)
		assert.Nil(t, deletedUser.Error)

	})

	t.Run("create a new user with note and credit", func(t *testing.T) {

		user1 := models.UserModel{
			Username: "john_doe",
			Email:    "john@example.com",
			Password: "securepassword123",
			Note: []models.NoteModel{
				{Content: "Note 1"},
				{Content: "Note 2"},
			},
			Credit: models.CreditModel{
				Number: "1234-5678-9012-3456",
			},
		}

		user2 := models.UserModel{
			Username: "Ras Muhammad",
			Email:    "ras@muhammad.com",
			Password: "securepassword123",
			Note: []models.NoteModel{
				{Content: "Note 1"},
				{Content: "Note 2"},
				{Content: "Note 3"},
				{Content: "Note 4"},
				{Content: "Note 5"},
				{Content: "Note 6"},
			},
			Credit: models.CreditModel{
				Number: "7777-7777-7777-7777",
			},
		}

		result1 := db.DB.Create(&user1)
		result2 := db.DB.Create(&user2)

		assert.Nil(t, result1.Error)
		assert.Nil(t, result2.Error)

	})

	t.Run("preload note by user id", func(t *testing.T) {
		var user models.UserModel

		result := db.DB.Preload("Note").Where("email = ?", "john@example.com").First(&user)

		assert.Nil(t, result.Error)

		fmt.Println(len(user.Note))
	})

	t.Run("preload credit by user id", func(t *testing.T) {
		var user models.UserModel

		result := db.DB.Preload("Credit").Where("email = ?", "john@example.com").First(&user)

		assert.Nil(t, result.Error)

		fmt.Println(user.Credit.Number)
	})

	t.Run("preload user by credit id", func(t *testing.T) {
		var credit models.CreditModel

		result := db.DB.Preload("User").First(&credit, 1)

		assert.Nil(t, result.Error)

		assert.NotNil(t, credit.User)

	})

	t.Run("preload user by note id", func(t *testing.T) {

		var notes []models.NoteModel

		result := db.DB.Where("id = ?", 4).Preload("User.Credit").Find(&notes)

		for _, note := range notes {
			fmt.Println(note.User.Username)
			fmt.Println(note.User.Credit.Number)

		}

		assert.Nil(t, result.Error)

	})

}
