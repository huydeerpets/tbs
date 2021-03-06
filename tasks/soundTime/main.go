package main

import (
	"math"
	"strconv"
	"unicode/utf8"

	"github.com/huydeerpets/tbs/tasks"
	"github.com/huydeerpets/tbs/utils/contribution"
	"github.com/huydeerpets/tbs/utils/sound"
)

var err error

func init() {
	if err = tasks.SetConfig(); err != nil {
		tasks.Err(err, "makeMovie")
	}
}

func main() {
	contribution()
}

// contribution
func contribution() error {
	contributionIDList, err := contributions.GetViewStatusPublicIDList()
	if err != nil {
		tasks.Err(err, "moveTime")
	}

	for _, id := range contributionIDList {
		second, err := sound.GetLength(strconv.Itoa(id))
		if err != nil {
			return err
		}

		if second == 0 {
			continue
		}

		detail, err := contributions.GetDetailByUserContributionID(id)
		if err != nil {
			return err
		}

		s, err := contributions.GetSearchWordBody(detail.Body)
		if err != nil {
			return err
		}

		err = contributions.AddOrSaveSoundLength(id, int(math.Ceil(second)), utf8.RuneCountInString(s))
		if err != nil {
			return err
		}

	}

	return nil
}
