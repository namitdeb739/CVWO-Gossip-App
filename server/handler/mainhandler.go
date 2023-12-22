package handler

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/namitdeb739/cvwo-gossip-app/database"
)

func CreateEntry(c *fiber.Ctx, tableType interface{}) error {
	db := database.DB.Db
	entryType := reflect.TypeOf(tableType)
	entry := reflect.New(entryType).Interface()

	err := c.BodyParser(entry)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error",
											"message": "Invalid input",
											"data": entryType})
	}

	err = db.Create(entry).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error",
											"message": "Could not create entry",
											"data": "err"})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success",
										"message": "Entry added created",
										"data": entry})
}

func GetAllEntries(c* fiber.Ctx, tableType interface{}) error {
	db := database.DB.Db
	entrySliceType := reflect.SliceOf(reflect.TypeOf(tableType))
	entries := reflect.New(entrySliceType).Interface()

	db.Find(entries)

	entrySliceVal := reflect.ValueOf(entries).Elem()

	if entrySliceVal.Len() == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error",
											"message": "Entries not found",
											"data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success",
										"message": "Entries found",
										"data": entries})
}

func GetSingleEntry(c *fiber.Ctx, tableType interface{}, primaryKeyFieldName string) error {
	db := database.DB.Db

	entryPrimaryKeyVal := c.Params(primaryKeyFieldName)
	entryType := reflect.TypeOf(tableType)
	entry := reflect.New(entryType).Interface()

	search := db.Where(primaryKeyFieldName + " = ?", entryPrimaryKeyVal).First(entry)
	if search.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error",
											"message": "Entry " + entryPrimaryKeyVal +  " not found",
											"data": nil})
	}
	
	val := reflect.ValueOf(entry).Elem()
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        if field.Kind() == reflect.Slice && field.Type().Elem().Kind() == reflect.Struct {
            associationName := val.Type().Field(i).Name
            db.Preload(associationName).Find(entry)
        }
    }

	return c.Status(200).JSON(fiber.Map{"status": "success",
										"message": "Entry found",
										"data": entry})
}

func UpdateEntry(c *fiber.Ctx, tableType interface{}, primaryKeyFieldName string) error {
	db := database.DB.Db

	entryType := reflect.TypeOf(tableType)
	entry := reflect.New(entryType).Interface()

	entryPrimaryKeyVal := c.Params(primaryKeyFieldName)
	search := db.Where(primaryKeyFieldName + " = ?", entryPrimaryKeyVal).First(entry)
	if search.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error",
											"message": "Entry " + entryPrimaryKeyVal +  " not found",
											"data": nil})
	}

	var updateEntry map[string]interface{}
	err := c.BodyParser(&updateEntry)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error",
											"message": "Invalid input",
											"data": updateEntry})
	}

	err = db.Model(entry).Updates(updateEntry).Error
	if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error",
												"message": "Could not update entry",
												"data": "err"})
		}

	return c.Status(200).JSON(fiber.Map{"status": "success",
										"message": "Record updated",
										"data": entry})
}

func DeleteEntry(c *fiber.Ctx, tableType interface{}, primaryKeyFieldName string) error {
	db := database.DB.Db
	entryType := reflect.TypeOf(tableType)
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
											"message": "Failed to delete entry",
											"data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "error",
										"message": "Entry " + entryPrimaryKeyVal + " deleted"})
}