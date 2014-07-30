package gift

import(
  "math/rand"
)

// Noise adds random noise to the image
// It takes a parameter for how much noise should be added to the
// image as a percentage in the range (0, 100)
//
// Example:
//
//	g := gift.New(
//		gift.Noise(100),
//	)
//	dst := image.NewRGBA(src.Bounds())
//	g.Draw(dst, src)
//
func Noise(percentage float32) Filter {
	adjustAmount := minf32(maxf32(percentage, 0.0), 100.0) / 100
	return &colorFilter{
		fn: func(px pixel) pixel {
      offset := rand.Float32() * adjustAmount
			r := px.R + offset
			g := px.G + offset
			b := px.B + offset
			return pixel{r, g, b, px.A}
		},
	}
}
