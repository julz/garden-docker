package gardendocker_test

import (
	"github.com/cloudfoundry-incubator/garden"
	"github.com/julz/garden-docker"
	"github.com/julz/garden-docker/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Backend", func() {
	var backend *gardendocker.Backend
	var repo gardendocker.Repo
	var fakeCreator *fakes.FakeCreator

	var createdContainer *gardendocker.Container

	BeforeEach(func() {
		fakeCreator = new(fakes.FakeCreator)
		repo = gardendocker.NewRepo()
		backend = &gardendocker.Backend{
			Creator: fakeCreator,
			Repo:    repo,
		}

		createdContainer = &gardendocker.Container{
			InfoHandler: &gardendocker.InfoHandler{Spec: garden.ContainerSpec{
				Handle: "was-created",
			}},
		}

		fakeCreator.CreateReturns(createdContainer, nil)
	})

	Describe("Create", func() {
		It("Creates a container using the container creator", func() {
			spec := garden.ContainerSpec{RootFSPath: "something"}
			backend.Create(spec)

			Expect(fakeCreator.CreateCallCount()).To(Equal(1))
			Expect(fakeCreator.CreateArgsForCall(0)).To(Equal(spec))
		})

		Context("after creation", func() {
			BeforeEach(func() {
				spec := garden.ContainerSpec{RootFSPath: "something", Handle: "ahandle"}
				backend.Create(spec)
			})

			It("adds the container to the repository", func() {
				Expect(repo.FindByHandle("was-created")).To(Equal(createdContainer))
			})
		})
	})
})
