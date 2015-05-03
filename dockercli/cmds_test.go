package dockercli_test

import (
	. "github.com/julz/garden-docker/dockercli"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cmds", func() {
	Describe("Create", func() {
		Context("with no volumes", func() {
			It("serializes to a docker cli command", func() {
				cmd := (&RunCmd{
					Program:     "foo",
					ProgramArgs: []string{"bar", "baz"},
					Image:       "some-image",
				}).Cmd()

				Expect(cmd.Path).To(ContainSubstring("docker"))
				Expect(cmd.Args).To(Equal([]string{
					"docker", "run", "some-image", "foo", "bar", "baz",
				}))
			})
		})

		Context("with volumes", func() {
			It("serializes to a docker cli command", func() {
				cmd := (&RunCmd{
					Program:     "foo",
					ProgramArgs: []string{"bar", "baz"},
					Image:       "some-image",
					Volumes:     []Volume{{"host", "container"}, {"host2", "container2"}},
				}).Cmd()

				Expect(cmd.Path).To(ContainSubstring("docker"))
				Expect(cmd.Args).To(Equal([]string{
					"docker", "run", "-v", "host:container", "-v", "host2:container2", "some-image", "foo", "bar", "baz",
				}))
			})
		})

		Context("with the detached flag", func() {
			It("adds the -d flag", func() {
				cmd := (&RunCmd{
					Program: "foo",
					Image:   "some-image",
					Detach:  true,
				}).Cmd()

				Expect(cmd.Path).To(ContainSubstring("docker"))
				Expect(cmd.Args).To(Equal([]string{
					"docker", "run", "-d", "some-image", "foo",
				}))
			})
		})
	})
})
