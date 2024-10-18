package validation

import (
	"bankingApp/middlewares"
	"bankingApp/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func VerifyAdminAuthorization(r *http.Request) (claims *models.Claims, err error) {
	claims, err = middlewares.VerifyJWT(r.Header.Get("Authorization"))
	if err != nil || !claims.IsAdmin {
		return nil, fmt.Errorf("unauthorized")
	}
	return claims, nil
}

func DecodeRequestBody(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return fmt.Errorf("invalid request data")
	}
	return nil
}

// func GetUserIDFromRequest(r *http.Request) (int, error) {
// 	vars := mux.Vars(r)
// 	idStr := vars["UserId"]

// 	userID, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		return 0, fmt.Errorf("invalid UserId parameter")
// 	}
// 	return userID, nil
// }

func GetIDFromRequest(r *http.Request, param string) (int, error) {
    vars := mux.Vars(r)
    idStr := vars[param]

    id, err := strconv.Atoi(idStr)
    if err != nil {
        return 0, fmt.Errorf("invalid %s parameter", param)
    }
    return id, nil
}

func VerifyCustomerAuthorization(r *http.Request) (*models.Claims, error) {
    claims, err := middlewares.VerifyJWT(r.Header.Get("Authorization"))
    if err != nil || !claims.IsCustomer {
        return nil, fmt.Errorf("unauthorized")
    }
    return claims, nil
}
