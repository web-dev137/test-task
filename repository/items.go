package repository

import (
	"database/sql"
	"reflect"

	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/web-dev137/test-task/models"
)

type ItemRepo interface {
	GetItems(ids []int) []models.Items
}

func (r *Repo) GetItems(ids []int) ([]models.Items, error) {
	q := "SELECT * FROM items WHERE id = ANY($1)" //query
	var items []models.Items
	rows, err := r.db.Query(q, pq.Array(ids)) //executing query and taking data from db
	if err != nil {
		if err == sql.ErrNoRows {
			log.WithFields(log.Fields{"error": err}).Error("Not found")
		} else {
			log.WithFields(log.Fields{"error": err}).Error("Internal error")
		}

		return nil, err
	}
	defer rows.Close()
	item := models.Items{}
	/*
	*Here we need get values for saving data from query
	 */
	fields := reflect.ValueOf(&item).Elem()
	numFields := fields.NumField()
	columns := make([]interface{}, numFields)
	for i := 0; i < numFields; i++ {
		field := fields.Field(i)
		columns[i] = field.Addr().Interface()
	}
	for rows.Next() {
		err = rows.Scan(columns...) //save into item
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error("Internal error")
			return items, err
		}
		items = append(items, item)
	}
	return items, nil
}
