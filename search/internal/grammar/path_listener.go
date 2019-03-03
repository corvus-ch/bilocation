// Code generated from search/internal/grammar/Path.g4 by ANTLR 4.7.2. DO NOT EDIT.

package grammar // Path
import "github.com/antlr/antlr4/runtime/Go/antlr"

// PathListener is a complete listener for a parse tree produced by PathParser.
type PathListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterNoneExpr is called when entering the noneExpr production.
	EnterNoneExpr(c *NoneExprContext)

	// EnterAnyExpr is called when entering the anyExpr production.
	EnterAnyExpr(c *AnyExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterTag is called when entering the tag production.
	EnterTag(c *TagContext)

	// EnterClassifierTagExpr is called when entering the classifierTagExpr production.
	EnterClassifierTagExpr(c *ClassifierTagExprContext)

	// EnterNameTagExpr is called when entering the nameTagExpr production.
	EnterNameTagExpr(c *NameTagExprContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitNoneExpr is called when exiting the noneExpr production.
	ExitNoneExpr(c *NoneExprContext)

	// ExitAnyExpr is called when exiting the anyExpr production.
	ExitAnyExpr(c *AnyExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)

	// ExitClassifierTagExpr is called when exiting the classifierTagExpr production.
	ExitClassifierTagExpr(c *ClassifierTagExprContext)

	// ExitNameTagExpr is called when exiting the nameTagExpr production.
	ExitNameTagExpr(c *NameTagExprContext)
}
