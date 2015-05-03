package dockercli_test

import (
	"errors"
	"os/exec"

	"github.com/cloudfoundry/gunk/command_runner/fake_command_runner"
	. "github.com/cloudfoundry/gunk/command_runner/fake_command_runner/matchers"
	. "github.com/julz/garden-docker/dockercli"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Docker CLI Runner", func() {
	var innerRunner *fake_command_runner.FakeCommandRunner
	var runner *Runner

	BeforeEach(func() {
		innerRunner = fake_command_runner.New()
		runner = &Runner{innerRunner}
	})

	Describe("Inspect", func() {
		It("runs the inspect command with the given field and container id and returns the result", func() {
			innerRunner.WhenRunning(fake_command_runner.CommandSpec{}, func(cmd *exec.Cmd) error {
				cmd.Stdout.Write([]byte("some-field-value\n"))
				return nil
			})

			field, err := runner.Inspect(InspectCmd{
				ContainerID: "some-container",
				Field:       "some-field",
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(innerRunner).To(HaveExecutedSerially(fake_command_runner.CommandSpec{
				Path: "docker",
				Args: []string{
					"inspect",
					"--format='{{.some-field}}'",
					"some-container",
				},
			}))

			Expect(field).To(Equal("some-field-value"))
		})

		Context("when the command fails", func() {
			It("returns an error including the stderr stream", func() {
				innerRunner.WhenRunning(fake_command_runner.CommandSpec{}, func(cmd *exec.Cmd) error {
					cmd.Stderr.Write([]byte("no foo\n"))
					return errors.New("exit status 2")
				})

				cmd := InspectCmd{}
				_, err := runner.Inspect(cmd)
				Expect(err).To(MatchError("inspect: exit status 2: no foo"))
			})
		})
	})

	Describe("Run", func() {
		It("runs the docker run command", func() {
			cmd := RunCmd{}
			_, err := runner.Run(cmd)
			Expect(err).NotTo(HaveOccurred())

			Expect(innerRunner).To(HaveExecutedSerially(fake_command_runner.CommandSpec{
				Path: "docker",
			}))
		})

		It("responds with the printed container id", func() {
			innerRunner.WhenRunning(fake_command_runner.CommandSpec{}, func(cmd *exec.Cmd) error {
				cmd.Stdout.Write([]byte("container-id\n"))
				return nil
			})

			cmd := RunCmd{}
			containerID, err := runner.Run(cmd)
			Expect(err).NotTo(HaveOccurred())

			Expect(containerID).To(Equal("container-id"))
		})

		Context("when the run command fails", func() {
			It("returns an error including the stderr stream", func() {
				innerRunner.WhenRunning(fake_command_runner.CommandSpec{}, func(cmd *exec.Cmd) error {
					cmd.Stderr.Write([]byte("no foo\n"))
					return errors.New("exit status 2")
				})

				cmd := RunCmd{}
				_, err := runner.Run(cmd)
				Expect(err).To(MatchError("run: exit status 2: no foo"))
			})
		})
	})
})
