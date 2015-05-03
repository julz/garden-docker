package gardendocker_test

import (
	. "github.com/julz/garden-docker"
	"github.com/julz/garden-docker/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Net", func() {
	var container *NetHandler
	var fakeChain *fakes.FakeChain

	BeforeEach(func() {
		fakeChain = new(fakes.FakeChain)
		container = &NetHandler{
			Chain: fakeChain,
		}
	})

	Describe("NetIn", func() {
		It("forwards ports using iptables", func() {
			container.NetIn(123, 456)
			Expect(fakeChain.ForwardCallCount()).Should(Equal(1))
		})

		Context("when the host port is empty", func() {
			It("selects a unique host port", func() {
				container.NetIn(0, 456)
				container.NetIn(0, 456)
				Expect(fakeChain.ForwardCallCount()).Should(Equal(2))

				_, _, hostPort1, _, _, _ := fakeChain.ForwardArgsForCall(0)
				_, _, hostPort2, _, _, _ := fakeChain.ForwardArgsForCall(1)
				Expect(hostPort1).NotTo(Equal(hostPort2))
			})
		})
	})
})
