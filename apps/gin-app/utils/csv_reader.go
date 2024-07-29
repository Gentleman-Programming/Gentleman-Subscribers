package utils

import (
	"apps/gin-app/models"
	subscriberModel "apps/gin-app/services/subscribers/models"
	"encoding/csv"
	"net/http"
	"os"
)

func CsvReader(filePath string) (*[]subscriberModel.Subscriber, *models.AppError) {
	// abro el archivo
	file, err := os.Open(filePath)
	if err != nil {
		return nil, &models.AppError{
			Err:     err,
			Message: "No se pudo abrir el archivo",
			Code:    http.StatusInternalServerError,
		}
	}

	// si termina todo, cierra el archivo
	defer file.Close()

	// creo la lectura del archivo y leo sus propiedades
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, &models.AppError{
			Err:     err,
			Message: "No se pudo leer el archivo",
			Code:    http.StatusInternalServerError,
		}
	}

	// creo el resultado para luego modificar
	var result []subscriberModel.Subscriber

	// recorro los resultados y los mapeo en la property result
	for _, record := range records[1:] {
		result = append(result, subscriberModel.Subscriber{
			ID:      record[0],
			Name:    record[1],
			Email:   record[2],
			Country: record[6],
		})
	}

	// retorno el resultado
	return &result, nil
}
