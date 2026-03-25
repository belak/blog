package blog

import "embed"

//go:embed blog
var BlogFS embed.FS

//go:embed pages
var PagesFS embed.FS

//go:embed assets
var AssetsFS embed.FS
