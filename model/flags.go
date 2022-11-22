package model

import (
	"benchmark-cli/utils"
	"flag"
	"path/filepath"
)

var path = flag.String(
	"path",
	utils.Home()+"/StudioProjects/play-store-buy-me-a-coffee/benchmark",
	"Macrobenchmark root directory",
)

var Root = filepath.Dir(*path)

var Module = filepath.Base(*path)

var Origin = flag.String(
	"origin",
	"master",
	"Origin branch to be benchmarked",
)

var Head = flag.String(
	"head",
	"fix/fb-verify-dynamic-link",
	"Head branch to be benchmarked",
)
