// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package controller_test

import (
	"github.com/juju/errors"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/cmd/juju/controller"
	"github.com/juju/juju/environs/configstore"
	"github.com/juju/juju/testing"
)

type ListSuite struct {
	testing.FakeJujuHomeSuite
	store configstore.Storage
}

var _ = gc.Suite(&ListSuite{})

type errorStore struct {
	err error
}

func (errorStore) CreateInfo(envName string) configstore.EnvironInfo {
	panic("CreateInfo not implemented")
}

func (errorStore) List() ([]string, error) {
	panic("List not implemented")
}

func (e errorStore) ListSystems() ([]string, error) {
	return nil, e.err
}

func (errorStore) ReadInfo(envName string) (configstore.EnvironInfo, error) {
	panic("ReadInfo not implemented")
}

func (s *ListSuite) SetUpTest(c *gc.C) {
	s.FakeJujuHomeSuite.SetUpTest(c)
	s.store = configstore.NewMem()

	var envList = []struct {
		name       string
		serverUUID string
		modelUUID  string
	}{
		{
			name:       "test1",
			serverUUID: "test1-uuid",
			modelUUID:  "test1-uuid",
		}, {
			name:       "test2",
			serverUUID: "test1-uuid",
			modelUUID:  "test2-uuid",
		}, {
			name:      "test3",
			modelUUID: "test3-uuid",
		},
	}
	for _, env := range envList {
		info := s.store.CreateInfo(env.name)
		info.SetAPIEndpoint(configstore.APIEndpoint{
			Addresses:  []string{"localhost"},
			CACert:     testing.CACert,
			ModelUUID:  env.modelUUID,
			ServerUUID: env.serverUUID,
		})
		err := info.Write()
		c.Assert(err, jc.ErrorIsNil)
	}
}

func (s *ListSuite) TestControllerList(c *gc.C) {
	context, err := testing.RunCommand(c, controller.NewListCommandForTest(s.store))
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(testing.Stdout(context), gc.Equals, "test1\ntest3\n")
}

func (s *ListSuite) TestUnrecognizedArg(c *gc.C) {
	_, err := testing.RunCommand(c, controller.NewListCommandForTest(s.store), "whoops")
	c.Assert(err, gc.ErrorMatches, `unrecognized args: \["whoops"\]`)
}

func (s *ListSuite) TestListSystemsError(c *gc.C) {
	s.store = errorStore{err: errors.New("cannot read info")}
	_, err := testing.RunCommand(c, controller.NewListCommandForTest(s.store))
	c.Assert(err, gc.ErrorMatches, "failed to list controllers in config store: cannot read info")
}
