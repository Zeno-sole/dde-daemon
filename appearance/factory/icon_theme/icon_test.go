/**
 * Copyright (c) 2011 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package icon_theme

import (
	C "launchpad.net/gocheck"
	. "pkg.linuxdeepin.com/dde-daemon/appearance/utils"
	"testing"
)

type testWrapper struct{}

func init() {
	C.Suite(&testWrapper{})
}

func Test(t *testing.T) {
	C.TestingT(t)
}

func (*testWrapper) TestInfoList(c *C.C) {
	userDirs := []PathInfo{
		{
			BaseName: "",
			FilePath: "testdata",
			FileFlag: FileFlagUserOwned,
		},
	}

	list := getThemeList(nil, userDirs)
	c.Check(len(list), C.Equals, 1)
}

func (*testWrapper) TestWriteConfig(c *C.C) {
	c.Check(setGtk2Theme("testdata/gtkrc-2.0", "Deepin"), C.Not(C.NotNil))
	c.Check(setGtk3Theme("testdata/settings.ini", "Deepin"), C.Not(C.NotNil))
}

func (*testWrapper) TestIconHidden(c *C.C) {
	type testEnable struct {
		filename string
		ret      bool
	}

	var infos = []testEnable{
		{
			filename: "testdata/index.theme",
			ret:      true,
		},
		{
			filename: "testdata/index2.theme",
			ret:      false,
		},
		{
			filename: "testdata/index3.theme",
			ret:      false,
		},
		{
			filename: "testdata/xxxxxxx.theme",
			ret:      true,
		},
	}

	for _, info := range infos {
		c.Check(isIconHidden(info.filename), C.Equals, info.ret)
	}
}
