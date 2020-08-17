package routes

import (
	"github.com/inggit_prakasa/Employee/controllers"
	"github.com/labstack/echo"
	"html/template"
	"io"
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
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("html/*.html")),
	}
	e.Renderer = renderer

	e.Static("/", "html")
	e.POST("/", controllers.Dashboard)
	e.GET("/", controllers.Dashboard)

	//--------------------------------------------------------
	e.GET("/employee",controllers.EmployeePage)
	e.GET("/employee/:id",controllers.FindEmployee)
	e.POST("/employee",controllers.AddEmployee)
	e.PUT("/employee",controllers.UpdateEmployee)
	e.DELETE("/delemployee/:id",controllers.DeleteEmployee)

	//-----------------------------------------------------------
	e.GET("/attendance", controllers.AttendancePage)
	e.GET("/attendance/:id", controllers.FindAttendance)
	e.POST("/attendance",controllers.AddAttendance)
	e.PUT("/attendance",controllers.EditAttendance)
	e.DELETE("/delattendance/:id",controllers.DeleteAttendance)

	//------------------------------------------------------------
	e.GET("/salary", controllers.SalaryPage)
	e.GET("/salary/:id", controllers.FindSalary)
	e.POST("/salary",controllers.AddSalary)
	e.PUT("/salary",controllers.EditSalary)
	e.DELETE("/delsalary/:id",controllers.DeleteSalary)

	//----------------------------------------------------------------------
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)
	e.GET("/login", controllers.Login)

	return e
}
