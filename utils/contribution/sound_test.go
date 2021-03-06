package contributions

import (
	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/models/csv_models"
	"github.com/huydeerpets/tbs/tests"
	"github.com/huydeerpets/tbs/utils"
	"time"

	. "gopkg.in/check.v1"
)

type TestSound struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestSound{}
	t.SetTableNameList([]string{
		"user_contributions",
		"user_contribution_sounds",
		"user_contribution_sound_details",
		"user_contribution_sound_lengths",
	})

	var _ = Suite(t)
}

func (t *TestSound) TestGetByUserContributionID(c *C) {
	r, _ := GetSoundByUserContributionID(1)

	c.Check(r.UserContributionID, Equals, 1)
}

func (t *TestSound) TestGetSoundListByUserContributionIDList(c *C) {
	r, _ := GetSoundListByUserContributionIDList([]int{1})

	c.Check(r[0].UserContributionID, Equals, 1)
}

func (t *TestSound) TestGetSoundMapByUserContributionIDList(c *C) {
	r, _ := GetSoundMapByUserContributionIDList([]int{1})

	c.Check(r[1].UserContributionID, Equals, 1)
}

func (t *TestSound) TestAddSound(c *C) {
	AddSound(100, 1)

	r, _ := GetSoundByUserContributionID(100)

	c.Check(r.UserContributionID, Equals, 100)
}

func (t *TestSound) TestAddSoundDetailList(c *C) {
	char := GetCharacter{
		VoiceType: 1,
	}

	list := []GetBody{
		{
			Body:      "abc",
			TalkType:  models.TalkTypeImage,
			Priority:  1,
			Character: char,
		},
		{
			Body:      "def'a!-jhg",
			TalkType:  models.TalkTypeText,
			Priority:  2,
			Character: char,
		},
	}

	AddSoundDetailList(100, list)
	r, _ := GetSoundDetailListByUserContributionID(100)

	c.Check(r[1].UserContributionID, Equals, 100)
	c.Check(r[1].BodySound, Equals, "defajhg")
}

func (t *TestSound) TestSaveSoundDetailToBodySound(c *C) {
	err := SaveSoundDetailToBodySound(uint(1), "abcdef", 1)
	c.Check(err, Equals, nil)

	err = SaveSoundDetailToBodySound(uint(1), "abcdef", 2)
	c.Check(err, Not(Equals), nil)
}

func (t *TestSound) TestSaveSoundDetailTVoiceType(c *C) {
	err := SaveSoundDetailTVoiceType(uint(1), 1, 1)
	c.Check(err, Equals, nil)

	err = SaveSoundDetailTVoiceType(uint(1), 1, 2)
	c.Check(err, Not(Equals), nil)
}

func (t *TestSound) TestMakeSoundFile(c *C) {
	list := []models.UserContributionSoundDetail{
		{
			UserContributionID: 0,
			Priority:           1,
			VoiceType:          csvModels.VoiceTypeMeiNormal,
			BodySound:          "今日は雨だ",
		},
		{
			UserContributionID: 0,
			Priority:           2,
			VoiceType:          csvModels.VoiceTypeMeiAngry,
			BodySound:          "",
		},
		{
			UserContributionID: 0,
			Priority:           3,
			VoiceType:          0,
			BodySound:          "明日は晴れだ",
		},
		{
			UserContributionID: 0,
			Priority:           4,
			VoiceType:          csvModels.VoiceTypeMeiAngry,
			BodySound:          "でも、そのあと雨だ",
		},
		{
			UserContributionID: 0,
			Priority:           5,
			VoiceType:          csvModels.VoiceTypeMeiBashful,
			BodySound:          "明後日は曇りだ",
		},
		{
			UserContributionID: 0,
			Priority:           6,
			VoiceType:          csvModels.VoiceTypeMeiHappy,
			BodySound:          "3日後は晴れる",
		},
		{
			UserContributionID: 0,
			Priority:           7,
			VoiceType:          csvModels.VoiceTypeMeiSad,
			BodySound:          "来週も晴れるといいな",
		},
		{
			UserContributionID: 0,
			Priority:           8,
			VoiceType:          csvModels.VoiceTypeM100,
			BodySound:          "来月も晴れるといいな",
		},
	}

	r := MakeSoundFile(0, list)
	c.Check(r, Equals, nil)
}

func (t *TestSound) TestExistsSound(c *C) {
	r := ExistsSound(0)

	c.Check(r, Equals, true)
}

func (t *TestSound) TestUpdateSoundToMakeStatus(c *C) {
	r := UpdateSoundToMakeStatus(1, 1)

	c.Check(r, Equals, nil)
}

func (t *TestSound) TestUpdatesSoundToMakeStatusAndVoiceTypeByUserContributionID(c *C) {
	r := UpdatesSoundToMakeStatusAndVoiceTypeByUserContributionID(1, 1, 2)

	c.Check(r, Equals, nil)
}

func (t *TestSound) TestReplaceBodeySound(c *C) {
	r, _ := ReplaceBodeySound("映画通とガンダム00")

	c.Check(r, Equals, "えいがつうとガンダムダブルオー")
}

func (t *TestSound) TestGetSoundDetailListByMakeStatusMade(c *C) {
	r, _ := GetSoundDetailListByMakeStatusMade()

	c.Check(r[0].MakeStatus, Equals, models.MakeStatusMade)
}

func (t *TestSound) TestGetSoudDetailListBySpecifiedDays(c *C) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	utils.SetNow(time.Date(2015, 1, 3, 11, 00, 0, 0, loc))

	u, _ := GetSoundDetailListByMakeStatusMade()

	r := GetSoudDetailListBySpecifiedDays(u, 2)

	c.Check(len(r), Equals, 1)
}

func (t *TestSound) TestAddOrSaveSoundLength(c *C) {
	r := AddOrSaveSoundLength(1, 2, 3)
	c.Check(r, Equals, nil)

	r = AddOrSaveSoundLength(1, 2, 3)
	c.Check(r, Equals, nil)

	r = AddOrSaveSoundLength(2, 2, 3)
	c.Check(r, Equals, nil)
}
