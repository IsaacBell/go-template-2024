package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Soapstone-Services/go-template-2024"
	"github.com/Soapstone-Services/go-template-2024/pkg/utl/secure"
	errorUtils "github.com/Soapstone-Services/go-template-2024/pkg/utl/errors"


	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

func main() {
	dbInsert := `INSERT INTO public.companies VALUES (1, now(), now(), NULL, 'admin_company', true);
	INSERT INTO public.locations VALUES (1, now(), now(), NULL, 'admin_location', true, 'admin_address', 1);
	INSERT INTO public.roles VALUES (100, 100, 'SUPER_ADMIN');
	INSERT INTO public.roles VALUES (110, 110, 'ADMIN');
	INSERT INTO public.roles VALUES (120, 120, 'COMPANY_ADMIN');
	INSERT INTO public.roles VALUES (130, 130, 'LOCATION_ADMIN');
	INSERT INTO public.roles VALUES (200, 200, 'USER');`
	var psn = os.Getenv("PG_DB")
	queries := strings.Split(dbInsert, ";")

	u, err := pg.ParseURL(psn)
	errorUtils.CheckErr(err)
	db := pg.Connect(u)
	_, err = db.Exec("SELECT 1")
	errorUtils.CheckErr(err)
	createSchema(db, &template.Company{}, &template.Location{}, &template.Role{}, &template.User{})

	for _, v := range queries[0 : len(queries)-1] {
		_, err := db.Exec(v)
		errorUtils.CheckErr(err)
	}

	sec := secure.New(1, nil)

	userInsert := `INSERT INTO public.users (id, created_at, updated_at, first_name, last_name, username, password, email, active, role_id, company_id, location_id) VALUES (1, now(),now(),'Admin', 'Admin', 'admin', '%s', 'johndoe@mail.com', true, 100, 1, 1);`
	_, err = db.Exec(fmt.Sprintf(userInsert, sec.Hash("admin")))
	errorUtils.CheckErr(err)
}

func createSchema(db *pg.DB, models ...interface{}) {
	for _, model := range models {
		errorUtils.CheckErr(db.CreateTable(model, &orm.CreateTableOptions{
			FKConstraints: true,
		}))
	}
}
