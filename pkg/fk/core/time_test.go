package core_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/nolafw/faker/pkg/fk/common/util"
	"github.com/nolafw/faker/pkg/fk/core"
)

var _ = Describe("Tests for random time functions", func() {
	randTime := core.NewRandTime(util.RandSeed())

	past30Years := time.Now().Add(-30 * 365 * 24 * time.Hour)
	future30Years := time.Now().Add(30 * 365 * 24 * time.Hour)

	Describe("PastFuture", func() {
		It("should return any random past/future time", func() {
			r := randTime.PastFuture()

			Expect(r).To(BeTemporally(">=", past30Years))
			Expect(r).To(BeTemporally("<=", future30Years))
		})
	})
	// TODO:
	Describe("PastFrom", func() {})

	Describe("Past", func() {})

	Describe("FutureTo", func() {})

	Describe("Future", func() {})

	Describe("TimeRange", func() {})

	Describe("Duration", func() {})

	Describe("DurationMilliSec", func() {})

	Describe("DurationSec", func() {})

	Describe("DurationMin", func() {})

	Describe("DurationHour", func() {})

	Describe("DurationTo", func() {})

	Describe("DurationRange", func() {})
})
