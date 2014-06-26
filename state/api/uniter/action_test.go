// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package uniter_test

import (
	"github.com/juju/juju/state/api/params"
	"github.com/juju/juju/state/api/uniter"
	"github.com/juju/names"
	gc "launchpad.net/gocheck"
)

type actionSuite struct {
	uniterSuite
}

var _ = gc.Suite(&actionSuite{})
var basicParams = map[string]interface{}{"outfile": "foo.txt"}

func (s *actionSuite) TestAction(c *gc.C) {

	var actionTests = []struct {
		description string
		action      params.Action
	}{{
		description: "A simple Action.",
		action: params.Action{
			Name:   "snapshot",
			Params: basicParams,
		},
	}, {
		description: "An Action with nested parameters.",
		action: params.Action{
			Name: "backup",
			Params: map[string]interface{}{
				"outfile": "foo.bz2",
				"compression": map[string]interface{}{
					"kind":    "bzip",
					"quality": float64(5.0),
				},
			},
		},
	}}

	for i, actionTest := range actionTests {
		c.Logf("test %d: %s", i, actionTest.description)
		actionId, err := s.uniterSuite.wordpressUnit.AddAction(
			actionTest.action.Name,
			actionTest.action.Params)
		c.Assert(err, gc.IsNil)

		actionTag := names.NewActionTag(s.uniterSuite.wordpressUnit.UnitTag(), i)
		c.Assert(actionTag.Id(), gc.Equals, actionId)

		retrievedAction, err := s.uniter.Action(actionTag)
		c.Assert(err, gc.IsNil)

		c.Assert(retrievedAction.Name(), gc.DeepEquals, actionTest.action.Name)
		c.Assert(retrievedAction.Params(), gc.DeepEquals, actionTest.action.Params)
	}
}

func (s *actionSuite) TestActionNotFound(c *gc.C) {
	actionTag := names.NewActionTag(names.NewUnitTag("wordpress/0"), 0)
	_, err := s.uniter.Action(actionTag)
	c.Assert(err, gc.NotNil)
	c.Assert(err.Error(), gc.Equals, "action \"wordpress/0_a_0\" not found")
}

func (s *actionSuite) TestNewActionAndAccessors(c *gc.C) {
	testAction, err := uniter.NewAction("snapshot", basicParams)
	c.Assert(err, gc.IsNil)
	testName := testAction.Name()
	testParams := testAction.Params()
	c.Assert(testName, gc.Equals, "snapshot")
	c.Assert(testParams, gc.DeepEquals, basicParams)
}
