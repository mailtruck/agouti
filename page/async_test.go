package page_test

import (
	. "github.com/sclevine/agouti/page"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti/mocks"
	"github.com/sclevine/agouti/webdriver"
	"time"
)

var _ = Describe("Async", func() {
	var (
		async   FinalSelection
		failer  *mocks.Failer
		driver  *mocks.Driver
		element *mocks.Element
	)

	BeforeEach(func() {
		driver = &mocks.Driver{}
		failer = &mocks.Failer{}
		element = &mocks.Element{}
		async = NewPage(driver, failer).Within("#selector").ShouldEventually(500*time.Millisecond, 100*time.Millisecond)
	})

	Describe("#Selector", func() {
		It("returns the selector", func() {
			Expect(async.Selector()).To(Equal("#selector"))
		})
	})

	Describe("#ContainText", func() {
		BeforeEach(func() {
			driver.GetElementsCall.ReturnElements = []webdriver.Element{element}
			element.GetTextCall.ReturnText = "no match"
		})

		Context("if #ContainText eventually passes", func() {
			It("passes the test", func(done Done) {
				go func() {
					defer GinkgoRecover()
					Expect(func() { async.ContainText("text") }).NotTo(Panic())
					Expect(failer.DownCount).To(Equal(21))
					Expect(failer.UpCount).To(Equal(6))
					Expect(failer.AsyncCalled).To(BeTrue())
					Expect(failer.SyncCalled).To(BeTrue())
					Expect(failer.IsAsync).To(BeFalse())
					Expect(failer.ResetCalled).To(BeTrue())
					close(done)
				}()
				time.Sleep(400 * time.Millisecond)
				element.GetTextCall.ReturnText = "text"
			})
		})

		Context("if #ContainText never passes", func() {
			It("fails the test", func(done Done) {
				go func() {
					defer GinkgoRecover()
					Expect(func() { async.ContainText("text") }).To(Panic())
					Expect(failer.Message).To(Equal("After 500ms:\n FAILED"))
					Expect(failer.DownCount).To(Equal(25))
					Expect(failer.UpCount).To(Equal(6))
					Expect(failer.AsyncCalled).To(BeTrue())
					Expect(failer.SyncCalled).To(BeTrue())
					Expect(failer.IsAsync).To(BeFalse())
					Expect(failer.ResetCalled).To(BeFalse())
					close(done)
				}()
				time.Sleep(600 * time.Millisecond)
			})
		})
	})

	Describe("#HaveAttribute", func() {
		BeforeEach(func() {
			driver.GetElementsCall.ReturnElements = []webdriver.Element{element}
			element.GetAttributeCall.ReturnValue = "some other value"
		})

		Context("if #ContainText eventually passes", func() {
			It("passes the test", func(done Done) {
				go func() {
					defer GinkgoRecover()
					Expect(func() { async.HaveAttribute("some-attribute", "some value") }).NotTo(Panic())
					Expect(failer.DownCount).To(Equal(21))
					Expect(failer.UpCount).To(Equal(6))
					Expect(failer.AsyncCalled).To(BeTrue())
					Expect(failer.SyncCalled).To(BeTrue())
					Expect(failer.IsAsync).To(BeFalse())
					Expect(failer.ResetCalled).To(BeTrue())
					close(done)
				}()
				time.Sleep(400 * time.Millisecond)
				element.GetAttributeCall.ReturnValue = "some value"
			})
		})

		Context("if #ContainText never passes", func() {
			It("fails the test", func(done Done) {
				go func() {
					defer GinkgoRecover()
					Expect(func() { async.HaveAttribute("some-attribute", "some value") }).To(Panic())
					Expect(failer.Message).To(Equal("After 500ms:\n FAILED"))
					Expect(failer.DownCount).To(Equal(25))
					Expect(failer.UpCount).To(Equal(6))
					Expect(failer.AsyncCalled).To(BeTrue())
					Expect(failer.SyncCalled).To(BeTrue())
					Expect(failer.IsAsync).To(BeFalse())
					Expect(failer.ResetCalled).To(BeFalse())
					close(done)
				}()
				time.Sleep(600 * time.Millisecond)
			})
		})
	})
})
