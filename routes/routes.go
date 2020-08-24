package routes

import (
	"html/template"
	"io"

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
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("html/*.html")),
	}
	e.Renderer = renderer

	e.Static("/", "html")

	//jwtGroup := e.Group("/jwt")
	//jwtGroup.Use(middleware.TokenJwt)
	//jwtGroup.GET("/main", controllers.MainJwt)
	//e.GET("/createJWT",controllers.CreateJWT)

	AdminGroup := e.Group("/admin")
	AdminGroup.Use(controllers.CheckCookieLogin)
	AdminGroup.POST("/dashboard", controllers.Dashboard)
	AdminGroup.GET("/dashboard", controllers.Dashboard)

	//--------------------------------------------------------
	AdminGroup.GET("/employee", controllers.EmployeePage)
	AdminGroup.POST("/employee", controllers.EmployeePage)
	AdminGroup.GET("/employee/:id", controllers.FindEmployee)
	AdminGroup.POST("/updateemployee", controllers.UpdateEmployee)
	AdminGroup.GET("/addemployee", controllers.AddEmployeePage)
	AdminGroup.POST("/addemp", controllers.AddEmployee)
	AdminGroup.POST("/delemp/:id", controllers.DeleteEmployee)
	//e.DELETE("/employee", controllers.DeleteEmployee)

	//-----------------------------------------------------------------
	e.POST("/register", controllers.RegisterEmployee)
	e.GET("/register", controllers.Register)
	e.PUT("/status/:id", controllers.SetStatusEmployee)
	e.GET("/viewstatus/:id", controllers.ViewStatusEmployee)
	e.GET("/laporanall", controllers.LaporanAll)
	e.GET("/laporanbyid/:id", controllers.LaporanById)

	//-----------------------------------------------------------------
	//e.GET("/employee", controllers.GetAllEmployee)
	//e.GET("/employee/:id", controllers.FindEmployee)
	//e.POST("/employee", controllers.AddEmployee)
	//e.PUT("/employee", controllers.UpdateEmployee)
	//e.DELETE("/delemployee/:id", controllers.DeleteEmployee)

	//-----------------------------------------------------------
	AdminGroup.GET("/attendance", controllers.AttendancePage)
	AdminGroup.GET("/attendance/:id", controllers.FindAttendance)
	AdminGroup.POST("/attendance", controllers.AddAttendance)
	AdminGroup.PUT("/attendance", controllers.EditAttendance)
	AdminGroup.DELETE("/delattendance/:id", controllers.DeleteAttendance)

	//------------------------------------------------------------
	AdminGroup.GET("/salary", controllers.SalaryPage)
	AdminGroup.GET("/salary/:id", controllers.FindSalary)
	AdminGroup.POST("/salary", controllers.AddSalary)
	AdminGroup.PUT("/salary", controllers.EditSalary)
	AdminGroup.DELETE("/delsalary/:id", controllers.DeleteSalary)

	//----------------------------------------------------------------------
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/Checklogin", controllers.CheckLogin)
	e.GET("/login", controllers.Login)
	e.POST("/login", controllers.Login)
	e.POST("/logout", controllers.LogOut)

	return e
}
