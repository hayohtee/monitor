package htmx

import "embed"

// PageFS holds the embed filesystem
// for the pages directory which contains
// html pages.
//
//go:embed pages
var PageFS embed.FS