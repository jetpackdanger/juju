// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package cmd_test

import (
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"
	charmresource "gopkg.in/juju/charm.v6-unstable/resource"

	"github.com/juju/juju/resource/cmd"
)

var _ = gc.Suite(&FormatterSuite{})

type FormatterSuite struct {
	testing.IsolationSuite
}

func (s *FormatterSuite) TestFormatInfoOkay(c *gc.C) {
	data := []byte("spamspamspam")
	fp, err := charmresource.GenerateFingerprint(data)
	c.Assert(err, jc.ErrorIsNil)
	fingerprint := string(fp.Bytes())
	info := cmd.NewInfo(c, "spam", ".tgz", "X", fingerprint)
	formatted := cmd.FormatInfo(info)

	c.Check(formatted, jc.DeepEquals, cmd.FormattedInfo{
		Name:        "spam",
		Type:        "file",
		Path:        "spam.tgz",
		Comment:     "X",
		Revision:    0,
		Fingerprint: "fac09f7d67d1bd30f41135c16fc132e5b635988c162f353786ca28ec0605d4d8ed55799d5d02e9275473574bf3754975",
		Origin:      "upload",
	})
}
