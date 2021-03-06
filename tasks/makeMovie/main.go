package main

import (
	"errors"
	"flag"
	"os"
	"strconv"

	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/tasks"
	"github.com/huydeerpets/tbs/utils/contribution"
	"github.com/huydeerpets/tbs/utils/movie"
	"github.com/huydeerpets/tbs/utils/sound"
)

var followMap map[int]int
var contributionIDList []int
var err error
var logfile *os.File

func init() {
	if err = tasks.SetConfig(); err != nil {
		tasks.Err(err, "makeMovie")
	}
}

func main() {
	id := flag.Int("userContributionId", 0, "user_contribution_id")
	flag.Parse()

	if err := MakeMovie(*id); err != nil {
		contributions.AddOrSaveMovie(*id, "", models.MovieTypeYoutube, models.StatusError)
		tasks.Err(err, "makeMovie")
	}

}

// MakeMovie 
func MakeMovie(id int) error {

	if id == 0 {
		return errors.New("no userContributionId")
	}

	userMovie, err := contributions.GetMovie(id, models.MovieTypeYoutube)
	if err != nil {
		return err
	}

	if userMovie.MovieStatus == models.StatusRunning {
		return errors.New("runnig")
	}

	if userMovie.MovieStatus == models.StatusUploading {
		return errors.New("uploading")
	}

	if userMovie.ID == uint(0) {
		if err = contributions.AddMovie(id, "", models.MovieTypeYoutube, models.StatusRunning); err != nil {
			return err
		}
	} else {
		userMovie.MovieStatus = models.StatusRunning
		if err = userMovie.Save(); err != nil {
			return err
		}
	}

	list, err := contributions.GetSoundDetailListByUserContributionID(id)
	if err != nil {
		return err
	}

	// Audio file creation
	if err = contributions.MakeSoundFile(id, list); err != nil {
		return err
	}

	if err = contributions.UpdateSoundToMakeStatus(id, models.MakeStatusMade); err != nil {
		return err
	}

	if err = sound.ToM4a(strconv.Itoa(id)); err != nil {
		return err
	}

	// Movie file creation
	if err = movie.Make(strconv.Itoa(id)); err != nil {
		return err
	}

	if err = movie.ToFilter(strconv.Itoa(id)); err != nil {
		return err
	}

	userMovie, err = contributions.GetMovie(id, models.MovieTypeYoutube)
	if err != nil {
		return err
	}

	userMovie.MovieStatus = models.StatusMade

	return userMovie.Save()
}
