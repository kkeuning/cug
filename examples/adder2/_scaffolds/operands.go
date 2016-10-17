package main

import (
	"github.com/goadesign/goa"
	"github.com/kkeuning/cug/examples/adder2/_scaffolds/app"
)

// OperandsController implements the operands resource.
type OperandsController struct {
	*goa.Controller
}

// NewOperandsController creates a operands controller.
func NewOperandsController(service *goa.Service) *OperandsController {
	return &OperandsController{Controller: service.NewController("OperandsController")}
}

// Add runs the add action.
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	// OperandsController_Add: start_implement

	// Put your logic here

	// OperandsController_Add: end_implement
	return nil
}

// Multiply runs the multiply action.
func (c *OperandsController) Multiply(ctx *app.MultiplyOperandsContext) error {
	// OperandsController_Multiply: start_implement

	// Put your logic here

	// OperandsController_Multiply: end_implement
	return nil
}

// Subtract runs the subtract action.
func (c *OperandsController) Subtract(ctx *app.SubtractOperandsContext) error {
	// OperandsController_Subtract: start_implement

	// Put your logic here

	// OperandsController_Subtract: end_implement
	return nil
}
