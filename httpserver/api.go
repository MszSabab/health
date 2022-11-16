package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HealthModel struct {
	Status string `json:"status"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	hm := &HealthModel{}
	hm.Status = "ok"
	val, err := json.Marshal(hm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("error", err)
		return
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Println("its ok")
	_, _ = w.Write(val)
}
