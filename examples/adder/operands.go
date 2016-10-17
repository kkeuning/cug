//go:generate goagen bootstrap -d github.com/kkeuning/goa-react/examples/adder/design
package main

import (
	"strconv"

	"github.com/goadesign/goa"
	"github.com/kkeuning/goa-react/examples/adder/app"
)

// OperandsController implements the operands resource.
type OperandsController struct {
	*goa.Controller
}

// NewOperandsController creates a operands controller.
func NewOperandsController(service *goa.Service) *OperandsController {
	return &OperandsController{Controller: service.NewController("operands")}
}

// Add runs the add action.
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	ctx.OK([]byte(strconv.Itoa(ctx.Left + ctx.Right)))
	return nil
}
