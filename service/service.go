package service

import (
	"context"
	"log"

	"github.com/nodias/go-ApmExam3/database"

	"github.com/nodias/go-ApmCommon/model"
	"github.com/nodias/go-ApmCommon/response"
	"go.elastic.co/apm"
)

func GetUserInfo(ctx context.Context, id string) (*model.User, *response.ResponseError) {
	span, ctx := apm.StartSpan(ctx, "GetUserInfo", "custom")
	defer span.End()

	db := database.NewOpenDB()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, response.NewResponseError(err, 500)
	}
	user := model.User{}
	defer db.Close()
	row := tx.QueryRowContext(ctx, "SELECT * FROM schema_user.user WHERE id = $1", id)
	err = row.Scan(&user.Id, &user.Name)
	if err != nil {
		log.Printf("database.PostgreAccess.Get - %s", err)
		return nil, response.NewResponseError(err, 500)
	}
	log.Printf("database.PostgreAccess.Get - id: %s, name: %s", user.Id, user.Name)
	return &user, nil
}
