package routes

import (
	"html/template"
	"io"
	"net/http"

	"github.com/inggit_prakasa/Employee/controllers"

	"github.com/labstack/echo"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("html/*.html")),
	}
	e.Renderer = renderer

	e.Static("/", "html")

	e.GET("/employee", controllers.GetAllEmployee)
	e.GET("/employee/:id", controllers.FindEmployee)
	e.POST("/employee", controllers.AddEmployee)
	e.PUT("/employee", controllers.UpdateEmployee)
	e.DELETE("/employee", controllers.DeleteEmployee)

	//-----------------------------------------------------------------
	e.POST("/register", controllers.RegisterEmployee)
	e.PUT("/status/:id", controllers.SetStatusEmployee)
	e.GET("/viewstatus/:id", controllers.ViewStatusEmployee)
	e.GET("/laporanall", controllers.LaporanAll)
	e.GET("/laporanbyid/:id", controllers.LaporanById)
	//-----------------------------------------------------------------

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)
	e.GET("/login", controllers.Login)

	return e
}
