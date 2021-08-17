package Utils

//
//func GetImage(name string, sprite *Resources.Sprite) (*ebiten.Image, error) {
//	var rect image.Rectangle
//	var found bool
//	for _, img := range obj.Specs.Images {
//		if img.Name == name {
//			rect = image.Rect(img.X, img.Y, img.X+img.W, img.Y+img.H)
//			found = true
//			break
//		}
//	}
//	if !found {
//		return nil, fmt.Errorf("not found in %s", obj.Name)
//	}
//	img := obj.Image.SubImage(rect).(*ebiten.Image)
//	return img, nil
//}
//
//func newImage(src []byte) *ebiten.Image {
//	img, _, err := image.Decode(bytes.NewReader(src))
//	if err != nil {
//		log.Fatal(err)
//	}
//}
