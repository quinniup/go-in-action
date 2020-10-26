package _4_flyweight

import "testing"

func TestFlyweight(t *testing.T) {
	imageViewer1 := NewImageViewer("image.png")
	imageViewer1.Display()
}

func TestImageFlyweight(t *testing.T) {
	imageViewer1 := NewImageViewer("image.png")
	imageViewer2 := NewImageViewer("image.png")
	if imageViewer1.ImageFlyweight != imageViewer2.ImageFlyweight {
		t.Fatal("The image view is difference.")
	}
}
