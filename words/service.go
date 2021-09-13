package words

import (
	"net/http"

	"github.com/go-chi/chi"
)

/* Rest API */
func RestApiGetWords(w http.ResponseWriter, r *http.Request) {
	kata := chi.URLParam(r, "kata")

	words := GetWordKbbi(kata)
	if s := words; s.Arti != nil {
		HttpResponseSuccess(w, r, words)
		return
	}
	HttpResponseError(w, r, words, 404)
}
