package service

import (
	"context"
	"log"
	"strings"
	"unicode"

	"github.com/nodias/go-ApmExam3/database"

	"github.com/nodias/go-ApmCommon/model"
	"go.elastic.co/apm"
)

func GetUserInfo(ctx context.Context, id string) (*model.User, error) {
	span, ctx := apm.StartSpan(ctx, "GetUserInfo", "custom")
	defer span.End()
	db := database.NewOpenDB()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	user := model.User{}
	defer db.Close()
	row := tx.QueryRowContext(ctx, "SELECT * FROM schema_user.user WHERE id = $1", id)
	err = row.Scan(&user.Id, &user.Name)
	if err != nil {
		log.Printf("database.PostgreAccess.Get - %s", err)
		return nil, err
	}
	log.Printf("database.PostgreAccess.Get - id: %s, name: %s", user.Id, user.Name)
	return &user, nil
}

func PanicGenerator(ctx context.Context, name string) (int, error) {
	span, ctx := apm.StartSpan(ctx, "PanicGenerator", "custom")
	defer span.End()
	if strings.IndexFunc(name, func(r rune) bool { return r >= unicode.MaxASCII }) >= 0 {
		panic("non-ASCII name!")
	}
	return 0, nil
}
