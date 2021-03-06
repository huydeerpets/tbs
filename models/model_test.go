package models

import (
	"github.com/huydeerpets/tbs/tests"

	. "gopkg.in/check.v1"
)

type TestModel struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestModel{}
	t.SetTableNameList([]string{
		"user_masters",
	})

	var _ = Suite(t)
}

// UserMasterTest ユーザーテスト情報
type UserMasterTest struct {
	ID             uint
	Name           string
	Email          string `validate:"required,email"`
	ProfileImageID int    `json:"profile_image_id"`
}

func (t *TestModel) TestGetWhere(c *C) {
	u := UserMaster{}
	whereList := []map[string]interface{}{
		{"ID": 1},
	}
	option := map[string]interface{}{
		"order":  "ID desc",
		"offset": 0,
		"limit":  1,
	}

	db, _ := GetWhere(&u, "ID = :ID", whereList, option)

	r := u
	c.Check(r.ID, Equals, uint(1))
	c.Check(r.Name, Equals, "abc")
	c.Check(r.Email, Equals, "test@tedt.com")

	ut := UserMasterTest{}
	db.Table("user_masters").Scan(&ut)

	c.Check(ut.ID, Equals, uint(1))
}

func (t *TestModel) TestGetWhereRecordNotFound(c *C) {
	u := UserMaster{}
	whereList := []map[string]interface{}{
		{"ID": 999},
	}
	option := map[string]interface{}{
		"order":  "ID desc",
		"offset": 0,
		"limit":  1,
	}

	_, err := GetWhere(&u, "ID = :ID", whereList, option)
	c.Check(err, IsNil)
	r := u
	c.Check(r.ID, Equals, uint(0))

}

func (t *TestModel) TestGetLisWhere(c *C) {
	u := []UserMaster{}
	whereList := []map[string]interface{}{
		{"ID": 1},
	}
	option := map[string]interface{}{
		"order":  "ID desc",
		"offset": 0,
		"limit":  1,
		"select": "id, name, email",
	}

	db, _ := GetListWhere(&u, "ID = :ID", whereList, option)

	r := u

	c.Check(r[0].ID, Equals, uint(1))
	c.Check(r[0].Name, Equals, "abc")
	c.Check(r[0].Email, Equals, "test@tedt.com")

	ut := []UserMasterTest{}
	db.Table("user_masters").Scan(&ut)

	c.Check(ut[0].ID, Equals, uint(1))
}

func (t *TestModel) TestUpdate(c *C) {
	u := []UserMaster{}
	whereList := []map[string]interface{}{
		{"ID": 1},
	}
	option := make(map[string]interface{})

	p := []interface{}{
		"name",
		"hellow",
	}

	Update(&u, p, "ID = :ID", whereList, option)

	r := UserMaster{}
	GetWhere(&r, "ID = :ID", whereList, option)

	c.Check(r.Name, Equals, "hellow")
}

func (t *TestModel) TestUpdates(c *C) {
	u := []UserMaster{}
	whereList := []map[string]interface{}{
		{"ID": 1},
	}
	option := make(map[string]interface{})

	p := UserMaster{
		Name:  "Hellow",
		Email: "test100@a.com",
	}

	Updates(&u, p, "ID = :ID", whereList, option)

	r := UserMaster{}
	GetWhere(&r, "ID = :ID", whereList, option)

	c.Check(r.Name, Equals, "Hellow")
	c.Check(r.Email, Equals, "test100@a.com")
}

func (t *TestModel) TestGetLisWhereRecordNotFound(c *C) {
	u := []UserMaster{}
	whereList := []map[string]interface{}{
		{"ID": 999},
	}
	option := map[string]interface{}{
		"order":  "ID desc",
		"offset": 0,
		"limit":  1,
	}

	_, err := GetListWhere(&u, "ID = :ID", whereList, option)
	c.Check(err, IsNil)
}

func (t *TestModel) TestCreate(c *C) {
	u := UserMaster{
		Name:           "abcdef",
		Email:          "abc@com",
		Password:       "xxxx",
		ProfileImageID: 1,
	}

	Create(&u)

	whereList := []map[string]interface{}{
		{"ID": 3},
	}
	option := make(map[string]interface{})

	GetWhere(&u, "ID = :ID", whereList, option)

	r := u
	c.Check(r.ID, Equals, uint(3))
	c.Check(r.Name, Equals, "abcdef")
}

func (t *TestModel) TestSave(c *C) {
	whereList := []map[string]interface{}{
		{"ID": 1},
	}
	option := make(map[string]interface{})
	u := UserMaster{}

	GetWhere(&u, "ID = :ID", whereList, option)

	u.Name = "xyz"

	Save(&u)

	GetWhere(&u, "ID = :ID", whereList, option)

	r := u
	c.Check(r.ID, Equals, uint(1))
	c.Check(r.Name, Equals, "xyz")
}

func (t *TestModel) TestInsertBatch(c *C) {
	Truncate("user_masters")

	m := []map[string]interface{}{
		{
			"id":               1,
			"name":             "abc",
			"email":            "test@tedt.com",
			"password":         "abc",
			"profile_image_id": 1,
			"created_at":       "2015-01-01 10:00:00",
			"updated_at":       "2015-01-01 10:00:00",
		},
		{
			"id":               2,
			"name":             "def",
			"email":            "test@tedt.com",
			"password":         "abc",
			"profile_image_id": 1,
			"created_at":       "2015-01-01 10:00:00",
			"updated_at":       "2015-01-01 10:00:00",
		},
		{
			"id":               3,
			"name":             "abcdef",
			"email":            "test@tedt.com",
			"password":         "abc",
			"profile_image_id": 1,
			"created_at":       "2015-01-01 10:00:00",
			"updated_at":       "2015-01-01 10:00:00",
		},
	}

	InsertBatch("user_masters", m)

	u := UserMaster{}
	whereList := []map[string]interface{}{
		{"ID": 3},
	}
	option := make(map[string]interface{})

	GetWhere(&u, "ID = :ID", whereList, option)

	r := u
	c.Check(r.ID, Equals, uint(3))
	c.Check(r.Name, Equals, "abcdef")
}

func (t *TestModel) TestRollback(c *C) {
	tx := Begin()

	u := UserMaster{
		Name:           "abcdef",
		Email:          "abc@com",
		Password:       "xxxx",
		ProfileImageID: 1,
	}

	Create(&u)

	Rollback(tx)

	whereList := []map[string]interface{}{
		{"ID": 3},
	}
	option := make(map[string]interface{})

	r := UserMaster{}
	GetWhere(&r, "ID = :ID", whereList, option)

	c.Check(r.ID, Equals, uint(0))
}

func (t *TestModel) TestCommit(c *C) {
	tx := Begin()

	u := UserMaster{
		Name:           "abcdef",
		Email:          "abc@com",
		Password:       "xxxx",
		ProfileImageID: 1,
	}

	Create(&u)

	Commit(tx)

	whereList := []map[string]interface{}{
		{"ID": 3},
	}
	option := make(map[string]interface{})

	r := UserMaster{}
	GetWhere(&r, "ID = :ID", whereList, option)

	c.Check(r.ID, Equals, uint(3))
}

func (t *TestModel) TestLock(c *C) {
	tx := Begin()

	Lock("user_masters", 1)

	Commit(tx)
}
