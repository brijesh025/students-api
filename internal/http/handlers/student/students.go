package students

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/brijesh025/students-api/internal/types"
	"github.com/brijesh025/students-api/internal/utils/response"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating a Student")
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		// if err!=nil {
		// 	slog.Error("decoding failed")
		// }
		if errors.Is(err, io.EOF){
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		if(err != nil){
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		// ------------------> we should always validate our request made by user at our end
		
		response.WriteJson(w, http.StatusCreated, map[string] interface{} {"sucsess": "OK", "student": student})
	}
}