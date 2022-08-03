package errs

<<<<<<< HEAD
import "net/http"

type AppErr struct {
	Code    int `json:"code,omitempty"`
	Message string `json:"message"`
}

func (e AppErr) AsMessage()*AppErr{
=======
import (
	"net/http"
)

type AppErr struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

func (e AppErr) AsMessage() *AppErr {
>>>>>>> 933acab27e67785d8e63b2f07537dd6e22bf744f
	return &AppErr{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *AppErr {
	return &AppErr{
<<<<<<< HEAD
		Code: http.StatusNotFound,
=======
		Code:    http.StatusNotFound,
>>>>>>> 933acab27e67785d8e63b2f07537dd6e22bf744f
		Message: message,
	}
}

func NewUnexpectedError(message string) *AppErr {
	return &AppErr{
<<<<<<< HEAD
		Code: http.StatusInternalServerError,
		Message: message,
	}	
}

=======
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewBadRequestError(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewValidationError(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}

func NewAutenticationError(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}


>>>>>>> 933acab27e67785d8e63b2f07537dd6e22bf744f
