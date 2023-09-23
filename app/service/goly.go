package service

import (
	"github.com/Garry028/url-shortener/database"
	"github.com/Garry028/url-shortener/model"
)

func GetAllGolies() ([]model.Goly, error) {
	var golies []model.Goly

	// this find all the matching records
	tx := database.DATABASE.Find(&golies)
	if tx.Error != nil {
		return []model.Goly{}, tx.Error
	}
	return golies, nil
}

func GetGoly(id uint64) (model.Goly, error) {
	var goly model.Goly

	tx := database.DATABASE.Where("id= ?", id).First(&goly)

	// if error then return error and empty goly struct
	if tx.Error != nil {
		return model.Goly{}, tx.Error
	}
	return goly, nil
	// if we got response then return response and nil as error
}

func CreateGol(goly model.Goly) error {
	tx := database.DATABASE.Create(&goly)
	// if error then error will returned else error
	return tx.Error
}

func UpdateGoly(goly model.Goly) error{
	tx := database.DATABASE.Save(&goly)
	return tx.Error
}

func DeleteGoly(id uint64) error {
	tx := database.DATABASE.Unscoped().Delete(&model.Goly{}, id)
	return tx.Error
}

func FindByGolyUrl(url string) (model.Goly,error){
	var goly model.Goly

	tx:= database.DATABASE.Where("goly = ?",url).First(&goly)
	return goly,tx.Error
}