package utils

import (
	"flag"
	"path/filepath"
)

var path = flag.String(
	"path",
	Home()+"/StudioProjects/play-store-buy-me-a-coffee/benchmark",
	"Macrobenchmark root directory",
)

var Root = filepath.Dir(*path)

var Module = filepath.Base(*path)

var Origin = flag.String(
	"origin",
	"master",
	"Origin branch to be benchmarked",
)

/**
*
* INSTALL_FAILED_VERSION_DOWNGRADE
* Head branch is not allowed to have an older version code.
*
**/

var Head = flag.String(
	"head",
	"master",
	"Head branch to be benchmarked",
)
