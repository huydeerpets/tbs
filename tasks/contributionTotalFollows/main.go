package main

import (
	"os"

	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/tasks"
	"github.com/huydeerpets/tbs/utils/contribution"
	"github.com/huydeerpets/tbs/utils/follow"
	"github.com/huydeerpets/tbs/utils/log"
)

var followMap map[int]int
var contributionIDList []int
var err error
var logfile *os.File

func init() {
	if err = tasks.SetConfig(); err != nil {
		tasks.Err(err, "contributionTotalFollows")
	}
}

func main() {
	logs.Batch("start", "contributionTotalFollows")

	tx := models.Begin()

	if err = AddContributionTotalFollows(); err != nil {
		models.Rollback(tx)
		tasks.Err(err, "contributionTotalFollows")
		return
	}

	models.Commit(tx)

	if err = SaveUserContributionSearchToFollowCount(); err != nil {
		tasks.Err(err, "contributionTotalFollows")
		return
	}

	logs.Batch("finish", "contributionTotalFollows")
}

// AddContributionTotalFollowsAdd a Follow number
func AddContributionTotalFollows() error {
	contributionIDList, err = contributions.GetViewStatusPublicIDList()
	if err != nil {
		return err
	}

	followList, err := follows.GetListByUserContributionIDList(contributionIDList)
	if err != nil {
		return err
	}

	followMap = follows.GetFollowCountMap(followList)

	follows.TruncateTotal()

	if err = follows.AddTotalMap(followMap); err != nil {
		return err
	}

	return nil
}

// SaveUserContributionSearchToFollowCount Update Search Follows
func SaveUserContributionSearchToFollowCount() error {
	search, err := contributions.GetSearchListByUserContributionIDList(contributionIDList)
	if err != nil {
		return err
	}

	return contributions.SaveToFollowCount(search, followMap)
}
