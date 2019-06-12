package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

func main()  {

	twoScale("scale/1.jpg", "scale/out.jpg", 2)

}
//简单的缩放算法
func simpleScale(filename string, outfile string, scale float64)  {
	jp,_ := os.Open(filename)
	jpImage,_ := jpeg.Decode(jp)
	bounds := jpImage.Bounds()

	rect := image.Rect(bounds.Min.X,bounds.Min.Y,
		int(float64(bounds.Max.X) * scale) , int(float64(bounds.Max.Y) * scale))

	srcImg := image.NewRGBA(rect)
	white := color.RGBA{255,255,255,255}

	draw.Draw(srcImg, rect, image.NewUniform(white),image.ZP, draw.Src)

	for x := rect.Min.X; x< rect.Max.X ; x++  {
		for y := rect.Min.Y; y< rect.Max.Y ; y++  {
			color := jpImage.At( int(float64(x) * 1 / scale), int(float64(y) * 1 / scale))
			srcImg.Set(x, y, color)
		}
	}
	srcImage,_ := os.Create(outfile)
	if err := jpeg.Encode(srcImage,srcImg, nil); err !=nil {
		panic(err)
	}
}

//简单的缩放算法
func twoScale(filename string, outfile string, scale float64)  {
	jp,_ := os.Open(filename)
	jpImage,_ := jpeg.Decode(jp)
	bounds := jpImage.Bounds()

	rect := image.Rect(bounds.Min.X,bounds.Min.Y,
		int(float64(bounds.Max.X) * scale) , int(float64(bounds.Max.Y) * scale))

	srcImg := image.NewRGBA(rect)
	white := color.RGBA{255,255,255,255}

	draw.Draw(srcImg, rect, image.NewUniform(white),image.ZP, draw.Src)

	for x := rect.Min.X; x< rect.Max.X ; x++  {
		for y := rect.Min.Y; y< rect.Max.Y ; y++  {
			sx1 := int(float64(x) * 1 / scale)
			sy1 := int(float64(y) * 1 / scale)
			sx2 := sx1 + 1
			sy2 := sy1 + 1
			s1 := y - sy1
			s2 := sx2 - x
			s3 := 1.0 - s1
			s4 := 1.0 - s2
			v1 := jpImage.At(sx1,sy1)
			vr1,vg1,vb1,va1 := v1.RGBA()
			v2 := jpImage.At(sx2, sy1)
			vr2,vg2,vb2,va2 := v2.RGBA()
			v3 := jpImage.At(sx1, sy2)
			vr3,vg3,vb3,va3 := v3.RGBA()
			v4 := jpImage.At(sx2, sy2)
			vr4,vg4,vb4,va4 := v4.RGBA()

			vr0 := float64(vr1) *float64(s1) *float64(s4) + float64(vr2) * float64(s1) * float64(s2) +float64(vr3) *float64(s2) *float64(s3) + float64(vr4) *float64(s3) *float64(s4)

			vg0 := float64(vg1) *float64(s1) *float64(s4) + float64(vg2) * float64(s1) * float64(s2) +float64(vg3) *float64(s2) *float64(s3) + float64(vg4) *float64(s3) *float64(s4)

			vb0 := float64(vb1) *float64(s1) *float64(s4) + float64(vb2) * float64(s1) * float64(s2) +float64(vb3) *float64(s2) *float64(s3) + float64(vb4) *float64(s3) *float64(s4)

			va0 := float64(va1) *float64(s1) *float64(s4) + float64(va2) * float64(s1) * float64(s2) +float64(va3) *float64(s2) *float64(s3) + float64(va4) *float64(s3) *float64(s4)

			c := color.RGBA{uint8(vr0), uint8(vg0), uint8(vb0), uint8(va0)}

			//color := jpImage.At( int(float64(x) * 1 / scale), int(float64(y) * 1 / scale))
			srcImg.Set(x, y, c)
		}
	}
	srcImage,_ := os.Create(outfile)
	if err := jpeg.Encode(srcImage,srcImg, nil); err !=nil {
		panic(err)
	}
}