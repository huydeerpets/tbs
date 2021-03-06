package main

import (
	"os"

	"github.com/huydeerpets/tbs/tasks"
	"github.com/huydeerpets/tbs/utils/contribution"
	"github.com/huydeerpets/tbs/utils/log"
	"github.com/huydeerpets/tbs/utils/tag"
)

var followMap map[int]int
var contributionIDList []int
var err error
var logfile *os.File

func init() {
	if err = tasks.SetConfig(); err != nil {
		tasks.Err(err, "contributionSearch")
	}
}

func main() {
	logs.Batch("start", "contributionSearch")

	if err = ResetSearch(); err != nil {
		tasks.Err(err, "contributionSearch")
		return
	}

	logs.Batch("finish", "contributionSearch")
}

// ResetSearch 
func ResetSearch() error {
	if err = contributions.TruncateSearch(); err != nil {
		return err
	}

	contributionIDList, err = contributions.GetViewStatusPublicIDList()
	if err != nil {
		return err
	}

	for _, id := range contributionIDList {
		u, err := contributions.GetByUserContributionID(id)
		if err != nil {
			return err
		}

		t, err := tags.GetTagNameJoin(id)
		if err != nil {
			return err
		}

		detail, err := contributions.GetDetailByUserContributionID(id)
		if err != nil {
			return err
		}

		b, err := contributions.GetSearchWordBody(detail.Body)
		if err != nil {
			return err
		}

		searchWord := contributions.SearchWord{
			Title: u.Title,
			Body:  b,
			Tag:   t,
		}

		s := contributions.JoinSearchWord(searchWord)
		if err := contributions.AddSearch(id, s); err != nil {
			return err
		}
	}

	return nil
}
