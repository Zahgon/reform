package parse

import (
	"go/ast"
	"regexp"
)

var magicReformComment = regexp.MustCompile(`reform:([0-9A-Za-z_\.]+)`)

func fileGoType(x ast.Expr) string { _ = "STUB: not implemented"; return "" }

func commentText(g *ast.CommentGroup) string {
	_ = "STUB: not implemented"
	// this code used to just call g.Text(), but the behavior of this method changed in Go 1.15:
	// https://go-review.googlesource.com/c/go/+/224737
	return ""
}

func parseStructTypeSpec(ts *ast.TypeSpec, str *ast.StructType) (*StructInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// consider only fields with "reform:" tag

// strip quotes

// check for anonymous fields

// check for exported name

// parse tag and type

// File parses given file and returns found structs information.
func File(path string) ([]StructInfo, error) {
	_ = "STUB: not implemented"
	// parse file
	return nil, nil
}

// consider only top-level struct type declarations with magic comment

// ast.Print(fset, decl)

// ast.Print(fset, ts)

// magic comment may be attached to "type Foo struct" (TypeSpec)
// or to "type (" (GenDecl)

// ast.Print(fset, doc)

// ast.Print(fset, str)
