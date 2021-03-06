package models

import (
	"github.com/huydeerpets/tbs/tests"

	. "gopkg.in/check.v1"
)

type TestUserContribution struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserContribution{}
	t.SetTableNameList([]string{
		"user_contributions",
	})

	var _ = Suite(t)
}

func (t *TestUserContribution) TestGetIDAndAdd(c *C) {
	u := &UserContribution{
		UserID: 100,
		Title:  "ああああ",
	}

	r, _ := u.GetIDAndAdd()

	c.Check(r, Equals, uint(3))
}

func (t *TestUserContribution) TestSave(c *C) {
	u := &UserContribution{}
	uc, _, _ := u.GetByID(1)
	uc.Title = "test"

	r := uc.Save()

	c.Check(r, Equals, nil)
}

func (t *TestUserContribution) TestDelete(c *C) {
	u := &UserContribution{}
	uc, _, _ := u.GetByID(1)
	r := uc.Delete()

	c.Check(r, Equals, nil)
}

func (t *TestUserContribution) TestGetByID(c *C) {
	u := &UserContribution{}

	r, _, _ := u.GetByID(1)

	c.Check(r.ID, Equals, uint(1))
}

func (t *TestUserContribution) TestGetListByUserID(c *C) {
	u := &UserContribution{}

	r, _, _ := u.GetListByUserID(1, "ID", 10, 0)

	c.Check(r[0].ID, Equals, uint(1))
}

func (t *TestUserContribution) TestGetCountByUserID(c *C) {
	u := &UserContribution{}

	r, _ := u.GetCountByUserID(1, "ID")

	c.Check(r, Equals, 1)
}

func (t *TestUserContribution) TestGetByTop(c *C) {
	u := &UserContribution{}

	r, _, _ := u.GetByTop(0, 1)

	c.Check(r[0].ID, Equals, uint(2))
}

func (t *TestUserContribution) TestGetListByIDList(c *C) {
	u := &UserContribution{}

	r, _, _ := u.GetListByIDList([]int{1, 2})

	c.Check(r[0].ID, Equals, uint(1))
	c.Check(r[1].ID, Equals, uint(2))
}

func (t *TestUserContribution) TestGetListByViewStatusPublic(c *C) {
	u := &UserContribution{}

	r, _, _ := u.GetListByViewStatusPublic()

	c.Check(r[0].ID, Equals, uint(1))
	c.Check(r[1].ID, Equals, uint(2))
}
