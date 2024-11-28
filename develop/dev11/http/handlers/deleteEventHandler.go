package handlers

import (
	"L2_tasks/develop/dev11/http/cache"
	"L2_tasks/develop/dev11/http/domain"
	"L2_tasks/develop/dev11/http/models"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

func DeleteEventHandler(w http.ResponseWriter, r *http.Request, c *cache.Cache) {
	if r.Method != http.MethodPost {
		domain.ErrorLogger(w, errors.New("method error"))
		return
	}

	decoder := json.NewDecoder(r.Body)
	var decoded models.Event

	if decodingBodyErr := decoder.Decode(&decoded); decodingBodyErr != nil {
		domain.ErrorLogger(w, decodingBodyErr)
		return
	}

	dateQuery := decoded.Date
	timeQuery := decoded.Time

	if _, errParse := time.Parse("2006-01-02", dateQuery); errParse != nil {
		domain.ErrorLogger(w, errParse)
		return
	}

	if _, errParseTime := time.Parse("15:00", timeQuery); errParseTime != nil {
		domain.ErrorLogger(w, errParseTime)
		return
	}

	c.Delete(dateQuery, timeQuery)

	domain.ResponseLogger(w, "event deleted")
}
