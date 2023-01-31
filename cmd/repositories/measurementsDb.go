package repositories

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/storage"
	"time"
)

func CreateMeasurement(measurement models.Measurements) (models.Measurements, error) {
	db := storage.GetDB()
	sqlStatement := `INSERT INTO measurements (user_id, weight, height, body_fat, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(sqlStatement, measurement.UserId, measurement.Weight, measurement.Height, measurement.BodyFat, time.Now()).Scan(&measurement.Id)
	if err != nil {
		return measurement, err
	}

	return measurement, nil
}

func UpdateMeasurement(measurement models.Measurements, id int) (models.Measurements, error) {
	db := storage.GetDB()
	sqlStatement := `
	  UPDATE measurements
	  SET weight = $2, height = $3, body_fat = $4, created_at = $5
	  WHERE id = $1
	  RETURNING id`
	err := db.QueryRow(sqlStatement, id, measurement.Weight, measurement.Height, measurement.BodyFat, time.Now()).Scan(&id)
	if err != nil {
		return models.Measurements{}, err
	}
	measurement.Id = id
	return measurement, nil
}
