package lead

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   int
}

func GetLead(ctx *fiber.Ctx)    {}
func GetLeads(ctx *fiber.Ctx)   {}
func NewLead(ctx *fiber.Ctx)    {}
func DeleteLead(ctx *fiber.Ctx) {}
