package p

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type environment struct {
	env    string
	db     string
	dbUser string
	dbPass string
}

var config *environment

func init() {
	err := godotenv.Load(fmt.Sprintf("../%s.env.yml", os.Getenv("GO_ENV")))
	if err != nil {
		// .env読めなかった場合の処理
	}
	config = &environment{
		os.Getenv("ENV"),
		os.Getenv("DB"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
	}
}

func Main(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Message string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		switch err {
		case io.EOF:
			fmt.Fprint(w, "更新できてるかのテスト")
			fmt.Fprint(w, config.env)
			fmt.Fprint(w, config.db)
			fmt.Fprint(w, config.dbUser)
			fmt.Fprint(w, config.dbPass)
			return
		default:
			log.Printf("json.NewDecoder: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}

	if d.Message == "" {
		fmt.Fprint(w, "Hello World!")
		return
	}
	fmt.Fprint(w, html.EscapeString(d.Message))
}
