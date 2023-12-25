package handler

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/namitdeb739/cvwo-gossip-app/database"
)

// Generic functions for Create, GetAll, GetSingle, Update, and Delete (i.e. CRUD) 
// Using abstraction so that functions dont need to be redefined for each individual model
// ChatGPT4 Used to assist with:
// - Go's reflect package as I wasn't too familiar with how to use it
// - Preloading associations

func CreateEntry(c *fiber.Ctx, tableType interface{}) error {
	db := database.DB.Db
	entryType := reflect.TypeOf(tableType)
	entryTypeName := entryType.Name()
	entry := reflect.New(entryType).Interface()

	err := c.BodyParser(entry)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error",
											"message": "Invalid input: " + err.Error(),
											"data": entryType})
	}

	err = db.Create(entry).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error",
											"message": "Could not create " + entryTypeName + ": " + err.Error(),
											"data": "err"})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success",
										"message": entryTypeName + " added",
										"data": entry})
}

func GetAllEntries(c* fiber.Ctx, tableType interface{}) error {
	db := database.DB.Db
	entrySliceType := reflect.SliceOf(reflect.TypeOf(tableType))
	entryTypeName := reflect.TypeOf(tableType).Name() + "s"
	entries := reflect.New(entrySliceType).Interface()

	db.Find(entries)

	entrySliceVal := reflect.ValueOf(entries).Elem()

	if entrySliceVal.Len() == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error",
											"message": entryTypeName + " not found",
											"data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success",
										"message": entryTypeName + " found",
										"data": entries})
}

func GetSingleEntry(c *fiber.Ctx, tableType interface{}, primaryKeyFieldName string) error {
	db := database.DB.Db

	entryPrimaryKeyVal := c.Params(primaryKeyFieldName)
	entryType := reflect.TypeOf(tableType)
	entryTypeName := entryType.Name()
	entry := reflect.New(entryType).Interface()

	search := db.Where(primaryKeyFieldName + " = ?", entryPrimaryKeyVal).First(entry)
	if search.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error",
											"message": entryTypeName + " " + entryPrimaryKeyVal +  " not found",
											"data": nil})
	}
	
	// Preload all associations of the entry (show relations)
	val := reflect.ValueOf(entry).Elem()
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        if field.Kind() == reflect.Slice && field.Type().Elem().Kind() == reflect.Struct {
            associationName := val.Type().Field(i).Name
            db.Preload(associationName).Find(entry)
        }
    }

	return c.Status(200).JSON(fiber.Map{"status": "success",
										"message": entryTypeName + " found",
										"data": entry})
}

func UpdateEntry(c *fiber.Ctx, tableType interface{}, primaryKeyFieldName string) error {
	db := database.DB.Db

	entryType := reflect.TypeOf(tableType)
	entryTypeName := entryType.Name()
	entry := reflect.New(entryType).Interface()

	entryPrimaryKeyVal := c.Params(primaryKeyFieldName)
	search := db.Where(primaryKeyFieldName + " = ?", entryPrimaryKeyVal).First(entry)
	if search.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error",
											"message": entryTypeName + " " + entryPrimaryKeyVal +  " not found",
											"data": nil})
	}

	var updateEntry map[string]interface{}
	err := c.BodyParser(&updateEntry)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error",
											"message": "Invalid input: " + err.Error(),
											"data": updateEntry})
	}

	err = db.Model(entry).Updates(updateEntry).Error
	if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error",
												"message": "Could not update " + entryTypeName + ": " + err.Error(),
												"data": "err"})
		}

	return c.Status(200).JSON(fiber.Map{"status": "success",
										"message": entryTypeName + " updated",
										"data": entry})
}

func DeleteEntry(c *fiber.Ctx, tableType interface{}, primaryKeyFieldName string) error {
	db := database.DB.Db
	entryType := reflect.TypeOf(tableType)
	entryTypeName := entryType.Name()
	entry := reflect.New(entryType).Interface()

	entryPrimaryKeyVal := c.Params(primaryKeyFieldName)

	search := db.Where(primaryKeyFieldName + " = ?", entryPrimaryKeyVal).First(entry)
	if search.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error",
											"message": "Entry " + entryPrimaryKeyVal +  " not found",
											"data": nil})
	}

	err := db.Delete(entry, primaryKeyFieldName + " = ?", entryPrimaryKeyVal).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error",
											"message": "Failed to delete " + entryTypeName + ": " + err.Error(),
											"data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "error",
										"message": entryTypeName + " " + entryPrimaryKeyVal + " deleted"})
}