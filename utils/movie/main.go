package movie

import (
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/huydeerpets/tbs/utils"

	"github.com/astaxie/beego"

	youtube "google.golang.org/api/youtube/v3"
)

// Upload upload
type Upload struct {
	UserContributionID string
	Title              string
	Description        string
	CategoryID         string
	VideoStatus        string
}

// getRootPath 
func getRootPath() (string, error) {
	p, err := utils.GetAppPath()
	if err != nil {
		return "", err
	}

	return p + "/", nil
}

// Make 
func Make(file string) error {
	path, err := getRootPath()
	if err != nil {
		return err
	}

	movie := path + "static/movie/input.mp4"
	sound := path + "static/files/tmp/sound/" + file + ".m4a"
	dist := path + "static/files/tmp/movie/" + file + ".mp4"

	cmd := "ffmpeg -y -i " + movie + " -i " + sound + " -map 0:0 -map 1:0 -movflags faststart -vcodec libx264 -acodec copy " + dist
	_, err = exec.Command("sh", "-c", cmd).Output()

	return err
}

// ToFilter 
func ToFilter(file string) error {
	path, err := getRootPath()
	if err != nil {
		return err
	}

	filter := path + "static/movie/complex.mp4"
	src := path + "static/files/tmp/movie/" + file + ".mp4"
	dist := path + "static/files/movie/" + file + ".mp4"

	cmd := "ffmpeg -y -i " + src + " -i " + filter + " -filter_complex 'concat=n=2:v=1:a=1' " + dist
	_, err = exec.Command("sh", "-c", cmd).Output()

	return err
}

// ExecMakeMovie 
func ExecMakeMovie(id int) error {
	path, err := getRootPath()
	if err != nil {
		return err
	}

	cmd := "ENV_CONF_BATCH=" + beego.AppConfig.String("runmode") + " " + path + "tasks/makeMovie/makeMovie -userContributionId=" + strconv.Itoa(id)

	return exec.Command("sh", "-c", cmd).Start()
}

// UploadToYoutube 
func UploadToYoutube(client *http.Client, u Upload) (string, error) {
	if utils.IsTest() {
		return "", nil
	}

	path, err := getRootPath()
	if err != nil {
		return "", err
	}

	filename := path + "static/files/movie/" + u.UserContributionID + ".mp4"

	service, err := youtube.New(client)
	if err != nil {
		return "", err
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       u.Title,
			Description: u.Description,
			CategoryId:  u.CategoryID,
		},
		Status: &youtube.VideoStatus{PrivacyStatus: u.VideoStatus},
	}

	upload.Snippet.Tags = []string{"test", "upload", "api"}

	call := service.Videos.Insert("snippet,status", upload)

	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return "", err
	}

	response, err := call.Media(file).Do()
	if err != nil {
		return "", err
	}

	return response.Id, nil
}

// RemoveFile 
func RemoveFile(file string) error {
	path, err := getRootPath()
	if err != nil {
		return err
	}

	tmp := path + "static/files/tmp/movie/" + file + ".mp4"
	if err := os.Remove(tmp); err != nil {
		return err
	}

	mp4 := path + "static/files/movie/" + file + ".mp4"
	if err := os.Remove(mp4); err != nil {
		return err
	}

	return nil
}
