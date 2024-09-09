package handlers

import (
	"encoding/json"
	"net/http"
)

func EmailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Erro ao processar o formulário", http.StatusBadRequest)
			return
		}

		email := r.FormValue("email")

		response := map[string]string{
			"message": "Email enviado com sucesso!",
			"email":   email,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
