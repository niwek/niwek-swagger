package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	requestform "github.com/niwek/niwek-swagger/form/request"
	responseform "github.com/niwek/niwek-swagger/form/response"

	"github.com/niwek/niwek-swagger/log"
)

// CreateUser ...
func CreateUser(c *gin.Context) {
	const moduleName = "controller.user.CreateUser"
	var createUserRequest requestform.CreateUserRequest

	if err := c.Bind(&createUserRequest); err != nil {
		log.Error(err, moduleName)
		c.AbortWithStatusJSON(http.StatusBadRequest, responseform.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			ModuleName: moduleName,
		})
		return
	}

	// Mock the Response
	userResponse := responseform.UserResponse{
		ID:        uuid.New(),
		FirstName: createUserRequest.FirstName,
		LastName:  createUserRequest.LastName,
		Address: responseform.AddressResponse{
			Address1: createUserRequest.Address.Address1,
			Address2: createUserRequest.Address.Address2,
			City:     createUserRequest.Address.City,
			State:    createUserRequest.Address.State,
			Country:  createUserRequest.Address.Country,
			Zip:      createUserRequest.Address.Zip,
		},
		Dob:       createUserRequest.Dob,
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
	}

	c.JSON(http.StatusOK, userResponse)
	return
}

// GetUserByID ...
func GetUserByID(c *gin.Context) {
	const moduleName = "controller.user.GetUserByID"
	var err error
	var ID uuid.UUID

	idString := c.Param("id")
	if ID, err = uuid.Parse(idString); err != nil {
		log.Error(err, moduleName)
		c.AbortWithStatusJSON(http.StatusBadRequest, responseform.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			ModuleName: moduleName,
		})
		return
	}

	// Mock the Response
	userResponse := responseform.UserResponse{
		ID:        ID,
		FirstName: "firstName",
		LastName:  "lastNAme",
		Address: responseform.AddressResponse{
			Address1: "address1",
			Address2: "address2",
			City:     "addressCity",
			State:    "addressState",
			Country:  "addressCountry",
			Zip:      "addressZip",
		},
		Dob:       time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
	}

	c.JSON(http.StatusOK, userResponse)
	return
}
