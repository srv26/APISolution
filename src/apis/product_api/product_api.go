package product_api

import (
	"encoding/json"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"main.go/src/config"
	"main.go/src/entities"
	"main.go/src/models"
)

func QuerySelectFromGivenColorAndType(response http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	colortype := query["color"]
	cartype := query["type"]
	db, err := config.GetDB()
	if err != nil {
		responseWihError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		obj, err := productModel.QuerySelectFromGivenColorAndType(colortype[0], cartype[0]) // Please refer for comments in this api
		res := entities.Response{Success: true, Data: obj}
		if err != nil {
			responseWihError(response, http.StatusBadRequest, err.Error())
		} else {
			responseWithJson(response, http.StatusOK, res)
		}
	}
}

func QuerySelectEvertythingFromGivenColor(response http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	colortype := query["color"]
	db, err := config.GetDB()
	if err != nil {
		responseWihError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		obj, err := productModel.QuerySelectNameFromGivenColor(colortype[0]) // Please refer for comments in this api
		res := entities.Response{Success: true, Data: obj}
		if err != nil {
			responseWihError(response, http.StatusBadRequest, err.Error())
		} else {
			responseWithJson(response, http.StatusOK, res)
		}
	}
}

func DecribeTable(response http.ResponseWriter, r *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		responseWihError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		obj, err := productModel.DescribeTable() // Please refer for comments in this api
		res := entities.Response{Success: true, Data: obj}
		if err != nil {
			responseWihError(response, http.StatusBadRequest, err.Error())
		} else {
			responseWithJson(response, http.StatusOK, res)
		}
	}
}

func QuerySelectNameFromGivenColor(response http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	colortype := query["color"]
	db, err := config.GetDB()
	if err != nil {
		responseWihError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		obj, err := productModel.QuerySelectNameFromGivenColor(colortype[0]) // Please refer for comments in this api
		res := entities.Response{Success: true, Data: obj}
		if err != nil {
			responseWihError(response, http.StatusBadRequest, err.Error())
		} else {
			responseWithJson(response, http.StatusOK, res)
		}
	}
}

func FindInnerJoin(response http.ResponseWriter, r *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		responseWihError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		joinedRes, err := productModel.FindInnerJoin() // Please refer for comments in this api
		res := entities.Response{Success: true, Data: joinedRes}
		if err != nil {
			responseWihError(response, http.StatusBadRequest, err.Error())
		} else {
			responseWithJson(response, http.StatusOK, res)
		}
	}

}

func FindLeftJoin(response http.ResponseWriter, r *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		responseWihError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		joinedRes, err := productModel.FindLeftJoin() // Please refer for comments in this api
		res := entities.Response{Success: true, Data: joinedRes}
		if err != nil {
			responseWihError(response, http.StatusBadRequest, err.Error())
		} else {
			responseWithJson(response, http.StatusOK, res)
		}
	}

}

func FindRightJoin(response http.ResponseWriter, r *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		responseWihError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		joinedRes, err := productModel.FindRightJoin() // Please refer for comments in this api
		res := entities.Response{Success: true, Data: joinedRes}
		if err != nil {
			responseWihError(response, http.StatusBadRequest, err.Error())
		} else {
			responseWithJson(response, http.StatusOK, res)
		}
	}
}

func FindFullJoin(response http.ResponseWriter, r *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		responseWihError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		joinedRes, err := productModel.FindFullJoin() // Please refer for comments in ths api
		res := entities.Response{Success: true, Data: joinedRes}
		if err != nil {
			responseWihError(response, http.StatusBadRequest, err.Error())
		} else {
			responseWithJson(response, http.StatusOK, res)
		}
	}

}

//Sends an error response
func responseWihError(response http.ResponseWriter, code int, msg string) {
	responseWithJson(response, code, map[string]string{"error": msg})
}

//Sends a json response with valid data
func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Contenet-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
