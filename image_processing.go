package gamebot

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/go-vgo/robotgo"
	"gocv.io/x/gocv"
)

// (b *Bot) OpenImage will provide an image given a path. An error is returned if there
// is a problem reading the file.
func (b *Bot) OpenImage(path string) (*image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	return &img, err
}

// (b *Bot) CaptureWindow can be used to capture an image of the bot's set window.
func (b *Bot) CaptureWindow() *image.Image {
	wx, wy := b.config.window.position.x, b.config.window.position.y
	ww, wh := b.config.window.size.w, b.config.window.size.h
	bitRef := robotgo.CaptureScreen(wx, wy, ww, wh)
	// robotgo.MilliSleep(b.config.screenCaptureDelayMs)

	screenCap := robotgo.ToImage(bitRef)
	return &screenCap
}

// (b *Bot) DetectImage scan the bot's set window to detect images within the window.
// This function returns the minValue, maxValue, minLocation and maxLocation of the matched image.
// If an error occurs and error is returned.
//
// The `in` parameter represents the larger image where `tmpl` is the template image to search for.
//
// This function utilized opencv's TmCcoeffNormed algorithm by default. To change the algorithm use
// the `(b *Bot) SetCVMatchMode()`.
func (b *Bot) DetectImage(in *image.Image, tmpl *image.Image) (float32, float32, *image.Point, *image.Point, error) {
	inMat, err := gocv.ImageToMatRGB(*in)
	if err != nil {
		return 0, 0, nil, nil, fmt.Errorf("failed to convert img to gocv.Mat: %v", err)
	}
	tmplMat, err := gocv.ImageToMatRGB(*tmpl)
	if err != nil {
		return 0, 0, nil, nil, fmt.Errorf("failed to convert img to gocv.Mat: %v", err)
	}

	result, mask := gocv.NewMat(), gocv.NewMat()

	gocv.MatchTemplate(inMat, tmplMat, &result, b.config.cvMatchMode, mask)

	inMat.Close()
	tmplMat.Close()
	mask.Close()
	result.Close()

	mnv, mxv, mnl, mxl := gocv.MinMaxLoc(result)
	return mnv, mxv, &mnl, &mxl, nil
}

// (b *Bot) SetCVMatchMode returns the opencv template matching mode.
// Reference: [OpenCV Documentation](https://docs.opencv.org/4.6.0/df/dfb/group__imgproc__object.html) for more information.
func (b *Bot) CVMatchMode() string {
	b.config.botRWMut.Lock()
	defer b.config.botRWMut.Unlock()

	return b.config.cvMatchMode.String()
}

// (b *Bot) SetCVMatchMode set the opencv template matching mode.
// Reference: [OpenCV Documentation](https://docs.opencv.org/4.6.0/df/dfb/group__imgproc__object.html) for more information.
func (b *Bot) SetCVMatchMode(matchMode gocv.TemplateMatchMode) {
	b.config.botRWMut.Lock()
	defer b.config.botRWMut.Unlock()

	b.config.cvMatchMode = matchMode
}

// (b *Bot) ShowDetectedImage this function opens a window creating a rectangle around the min and max areas of the detected image.
// Pressing the 'q' key will close this window. This function is mostly used for debugging your bot.
//
// The `tmpl` is the template image to search for within the game window.
// This function will print the minValue, maxValue, MinLocation, and MaxLocation.
func (b *Bot) ShowDetectedImage(windowTitle string, tmpl *image.Image) error {
	w := gocv.NewWindow(windowTitle)
	imgX, imgY := (*tmpl).Bounds().Dx(), (*tmpl).Bounds().Dy()

	for {
		src := b.CaptureWindow()
		mnv, mxv, mnl, mxl, err := b.DetectImage(src, tmpl)
		if err != nil {
			return err
		}
		fmt.Println(mnv, mxv, mnl, mxl)

		srcMat, err := gocv.ImageToMatRGB(*src)
		if err != nil {
			return err
		}

		gocv.Rectangle(&srcMat, image.Rect(mxl.X, mxl.Y, mxl.X+imgX, mxl.Y+imgY), color.RGBA{255, 0, 0, 1}, 2)
		w.IMShow(srcMat)
		if w.WaitKey(1) == 113 {
			break
		}
	}

	return nil
}
