package main

import (
	"os"
	"strconv"

	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/tasks"
	"github.com/huydeerpets/tbs/utils/contribution"
	"github.com/huydeerpets/tbs/utils/log"
	"github.com/huydeerpets/tbs/utils/movie"
	"github.com/huydeerpets/tbs/utils/sound"
)

var followMap map[int]int
var contributionIDList []int
var err error
var logfile *os.File

// RemoveDays Delete日数
const RemoveDays = 3

func init() {
	if err = tasks.SetConfig(); err != nil {
		tasks.Err(err, "contributionTotalFollows")
	}
}

func main() {
	logs.Batch("start", "removeContribution")

	if err := RemoveSoundDetail(); err != nil {
		tasks.Err(err, "removeContribution")
	}

	if err := RemoveJoinFile(); err != nil {
		tasks.Err(err, "removeContribution")
	}

	logs.Batch("finish", "removeContribution")
}

// RemoveSoundDetail 音声詳細をDeleteする
func RemoveSoundDetail() error {
	list, err := contributions.GetSoundDetailListByMakeStatusMade()
	if err != nil {
		return err
	}

	list = contributions.GetSoudDetailListBySpecifiedDays(list, RemoveDays)

	for _, v := range list {
		file := strconv.Itoa(v.UserContributionID) + "_" + strconv.Itoa(v.Priority)
		sound.RemoveDetailFile(file)

		v.MakeStatus = models.MakeStatusUncreated
		v.Save()
	}

	return nil
}

// RemoveJoinFile 連結ファイルをDeleteする
func RemoveJoinFile() error {
	list, err := contributions.GetMovieListByMovieStatusPublic()
	if err != nil {
		return err
	}

	list = contributions.GetMovieListBySpecifiedDays(list, RemoveDays)

	for _, v := range list {
		if !contributions.ExistsMovie(v.UserContributionID) {
			continue
		}

		sound.RemoveJoinFile(strconv.Itoa(v.UserContributionID))
		movie.RemoveFile(strconv.Itoa(v.UserContributionID))
	}

	return nil
}
