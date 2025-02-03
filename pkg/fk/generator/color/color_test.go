package color

import (
	"testing"

	"github.com/nolafw/faker/pkg/fk/common/util"
	"github.com/nolafw/faker/pkg/fk/core"
	"github.com/nolafw/faker/pkg/fk/provider"
	"github.com/nolafw/faker/pkg/fk/provider/global"
	"github.com/yrichika/gest/pkg/gt"
)

func TestColor(testingT *testing.T) {
	coreRand := core.NewRand(util.RandSeed())
	global := &provider.Global{
		Colors: global.CreateColors(),
	}

	color := New(coreRand, global)

	tScn := gt.CreateTest(testingT)
	tScn.Describe("SafeName", func() {
		tScn.It("should return random safe color name", func() {
			colorName := color.SafeName()
			gt.Expect(tScn, &colorName).ToBeIn(global.Colors.SafeColorNames)
		})
	})

	tCn := gt.CreateTest(testingT)
	tCn.Describe("Name", func() {
		tCn.It("should return random color name", func() {
			colorName := color.Name()
			gt.Expect(tCn, &colorName).ToBeIn(global.Colors.AllColorNames)
		})
	})

	tHc := gt.CreateTest(testingT)
	tHc.Describe("Hex", func() {
		tHc.It("should return random hex color string", func() {
			r := color.Hex()
			gt.Expect(tHc, &r).ToMatchRegex(`^#[0-9a-f]{6}$`)
		})

	})

	tShc := gt.CreateTest(testingT)
	tShc.Describe("SafeHex", func() {
		tShc.It("should return random safe hex color string", func() {
			r := color.SafeHex()
			// FIXME: this regex is not exactly right.
			gt.Expect(tShc, &r).ToMatchRegex(`^#[0-9a-f]{6}$`)
		})
	})

	tRCN := gt.CreateTest(testingT)
	tRCN.Describe("RgbAsNum", func() {
		tRCN.It("should return random rgb color as numbers", func() {
			r, g, b := color.RgbAsNum()
			gt.Expect(tRCN, &r).ToBe_(gt.Between(0, 255))
			gt.Expect(tRCN, &g).ToBe_(gt.Between(0, 255))
			gt.Expect(tRCN, &b).ToBe_(gt.Between(0, 255))
		})
	})

	regex255 := `([01]?[0-9]?[0-9]|2[0-4][0-9]|25[0-5])`
	regexAlpha := `([01]?(\.\d+)?)`

	tRCS := gt.CreateTest(testingT)
	tRCS.Describe("RgbAsStr", func() {
		tRCS.It("should return random rgb color string", func() {
			r := color.RgbAsStr()
			gt.Expect(tRCS, &r).ToMatchRegex("^" + regex255 + "," + regex255 + "," + regex255 + "$")
		})
	})

	tRCA := gt.CreateTest(testingT)
	tRCA.Describe("RgbAsArr", func() {
		tRCA.It("should return random rgb color as array", func() {
			r := color.RgbAsArr()
			gt.Expect(tRCA, &r[0]).ToBe_(gt.Between(0, 255))
			gt.Expect(tRCA, &r[1]).ToBe_(gt.Between(0, 255))
			gt.Expect(tRCA, &r[2]).ToBe_(gt.Between(0, 255))
		})
	})

	tRgbCss := gt.CreateTest(testingT)
	tRgbCss.Describe("RgbCss", func() {
		tRgbCss.It("should return random rgb css color string", func() {
			r := color.RgbCss()
			gt.Expect(tRgbCss, &r).ToMatchRegex(`^rgb\(` + regex255 + `,` + regex255 + `,` + regex255 + `\)$`)
		})
	})

	tRgbaCss := gt.CreateTest(testingT)
	tRgbaCss.Describe("RgbaCss", func() {
		tRgbaCss.It("should return random rgba css color string", func() {
			r := color.RgbaCss()
			gt.Expect(tRgbaCss, &r).ToMatchRegex(`^rgba\(` + regex255 + `,` + regex255 + `,` + regex255 + `,` + regexAlpha + `\)$`)
		})
	})

	tHCN := gt.CreateTest(testingT)
	tHCN.Describe("HslAsNum", func() {
		tHCN.It("should return random hsl color as numbers", func() {
			h, s, l := color.HslAsNum()
			gt.Expect(tHCN, &h).ToBe_(gt.Between(0, 360))
			gt.Expect(tHCN, &s).ToBe_(gt.Between(0, 100))
			gt.Expect(tHCN, &l).ToBe_(gt.Between(0, 100))
		})
	})

	regex360 := `(?:36[0]|3[0-5][0-9]|[12][0-9][0-9]|[1-9]?[0-9])`
	regex100 := `(?:100|[1-9]?[0-9])`

	tHCS := gt.CreateTest(testingT)
	tHCS.Describe("HslAsStr", func() {
		tHCS.It("should return random hsl color string", func() {
			r := color.HslAsStr()
			gt.Expect(tHCS, &r).ToMatchRegex("^" + regex360 + "," + regex100 + "," + regex100 + "$")
		})
	})

	tHCA := gt.CreateTest(testingT)
	tHCA.Describe("HslAsArr", func() {
		tHCA.It("should return random hsl color as array", func() {
			r := color.HslAsArr()
			gt.Expect(tHCA, &r[0]).ToBe_(gt.Between(0, 360))
			gt.Expect(tHCA, &r[1]).ToBe_(gt.Between(0, 100))
			gt.Expect(tHCA, &r[2]).ToBe_(gt.Between(0, 100))
		})
	})
}
