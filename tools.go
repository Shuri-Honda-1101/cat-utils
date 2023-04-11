//go:build tools
// +build tools

// NOTE:https://zenn.dev/kenghaya/articles/2f41736042e054
// go mod tidyで消されてしまう(消したくない)モジュールを書いておく

package tools

import (
	_ "ariga.io/ogent"
	_ "entgo.io/contrib/entoas"
	_ "entgo.io/ent/entc"
	_ "github.com/hedwigz/entviz"
)
