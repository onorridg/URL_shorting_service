package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"time"
)

import (
	db "main/internal/database"
	"main/utils"
)

import (
	"github.com/gorilla/mux"
)

type RequestBody struct {
	Url string `json:"url"`
}

func bodyUrl(r *http.Request) string {
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return ""
	}
	var body RequestBody
	err = json.Unmarshal(b, &body)
	if err != nil {
		return ""
	}
	return body.Url
}

func realUrlExist(realUrl string, database *sql.DB) *db.DataRow {
	res := db.GetRow(db.REALURL, realUrl, database)
	if res != nil {
		return res
	}
	return nil
}

func getRedirectToRealUrl(w http.ResponseWriter, r *http.Request) {
	urlPath := path.Base(r.URL.Path)
	fmt.Println(urlPath)
	database := db.OpenDB()
	defer database.Close()
	data := db.GetRow(db.SHORTURL, urlPath, database)
	if data != nil {
		http.Redirect(w, r, fmt.Sprintf("http://%s", data.RealUrl), http.StatusSeeOther)
	} else {
		http.NotFound(w, r)
	}
}

func postCrateShortUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	database := db.OpenDB()
	defer database.Close()

	realUrl := bodyUrl(r)
	if realUrl == "" {
		http.Error(w, fmt.Sprintf("wrong real URL: %s", realUrl), 500)
		return
	}

	existUrl := realUrlExist(realUrl, database)
	if existUrl != nil {
		w.WriteHeader(http.StatusOK)
		result := fmt.Sprintf("{\"short_url\": \"%s\", \"real_url\": \"%s\"}",
			existUrl.ShortUrl, existUrl.RealUrl)
		w.Write([]byte(result))
		return
	}

	var shortUrl string
	for {
		shortUrl = utils.UrlGenerator(10)
		exist := db.GetRow(db.SHORTURL, shortUrl, database)
		if exist != nil {
			continue
		}
		break
	}
	db.InsertRow(realUrl, shortUrl, database)

	w.WriteHeader(http.StatusCreated)
	result := fmt.Sprintf("{\"short_url\": \"%s\", \"real_url\": \"%s\"}",
		shortUrl, realUrl)
	w.Write([]byte(result))
}

func initApiHandler(router *mux.Router) {
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", postCrateShortUrl).Methods(http.MethodPost)
	api.HandleFunc("/{key}", getRedirectToRealUrl).Methods(http.MethodGet)
}

func InitServer(addr string, port string) {
	router := mux.NewRouter()
	initApiHandler(router)

	server := &http.Server{
		Handler:      router,
		Addr:         addr + ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
