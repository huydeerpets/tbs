package controllers

import "github.com/huydeerpets/tbs/utils/image"

// SetImageFileResize Specify the file to be resized
func (c *BaseController) SetImageFileResize(f string, p string, w uint, h uint) (int, error) {
	if isTest() {
		return 0, nil
	}

	tmpPath := "./static/files/tmp/" + p + "/_tmp_" + f
	if err := c.ToFile(tmpPath); err != nil {
		return ErrImageConversion, err
	}

	tmpRootPath := "./static/files/tmp/" + p + "/" + f

	if err := images.PngToJpeg(tmpPath, tmpRootPath); err != nil {
		return ErrImageConversion, err
	}

	outputPath := "./static/files/" + p + "/" + f

	if err := images.Resize(tmpRootPath, outputPath, w, h); err != nil {
		return ErrImageResize, err
	}

	return 0, nil
}

// ToFile Save file
func (c *BaseController) ToFile(path string) error {
	if isTest() {
		return nil
	}

	return c.SaveToFile("file", path)
}
