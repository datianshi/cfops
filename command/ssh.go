package command

import (
	"bytes"
	"io"

	"golang.org/x/crypto/ssh"
)

// Config for the SSH connection
type Config struct {
	Username string
	Password string
	Host     string
	Port     int
}

// Copier copies from an io.Reader to an io.Writer
type Copier interface {
	Copy(dest io.Writer, src io.Reader) error
}

// DefaultCopier is an SSH copier
type DefaultCopier struct {
	client ClientInterface
}

type ClientInterface interface {
	NewSession() (SSHSession, error)
}

//Wrapper of ssh client to match client interface signature, since client.NewSession() does not use an interface
type SshClientWrapper struct {
	sshclient *ssh.Client
}

func NewClientWrapper(client *ssh.Client) *SshClientWrapper {
	return &SshClientWrapper{
		sshclient: client,
	}
}

func (c *SshClientWrapper) NewSession() (SSHSession, error) {
	return c.sshclient.NewSession()
}

func NewCopier(client ClientInterface) (copier *DefaultCopier) {
	copier = &DefaultCopier{
		client: client,
	}
	return
}

type SSHSession interface {
	Start(cmd string) error
	Wait() error
	StdoutPipe() (io.Reader, error)
	Close() error
}

func (copier *DefaultCopier) Copy(dest io.Writer, src io.Reader) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(src)
	s := buf.String()
	return copier.Execute(dest, s)
}

// Copy the output from a command to the specified io.Writer
func (copier *DefaultCopier) Execute(dest io.Writer, command string) (err error) {
	session, err := copier.client.NewSession()
	defer session.Close()
	if err != nil {
		return
	}
	stdoutReader, err := session.StdoutPipe()
	if err != nil {
		return
	}
	err = session.Start(command)
	if err != nil {
		return
	}
	_, err = io.Copy(dest, stdoutReader)
	if err != nil {
		return
	}
	err = session.Wait()
	return
}
