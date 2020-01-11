package main

import (
	"html/template"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"

	//"gitlab.com/username/online-service-and-customer-care/comment/Crepository"
	//"gitlab.com/username/online-service-and-customer-care/comment/cservice"
	"gitlab.com/username/online-service-and-customer-care/delivery/http/handler"
	"gitlab.com/username/online-service-and-customer-care/employee/employee_repository"
	"gitlab.com/username/online-service-and-customer-care/employee/employee_service"
)
func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	var tmpl = template.Must(template.ParseGlob("../../ui/templates/*"))
	tmpl.ExecuteTemplate(w,"index.layout",nil)
}



func main() {
	tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))
	dbconn, err := gorm.Open("postgres", "postgres://postgres:lealem-g@localhost/customercare?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	emplRepo := employee_repository.NewEmployeeGormRepo(dbconn)
	emplSrv := employee_service.NewEmployeeServiceGorm(emplRepo)
	//comRepo := Crepository.NewCommentGormRepo(dbconn)
	//comserv := cservice.NewCommentService(comRepo)
	employeeHandler := handler.NewEmployeeHandler(tmpl, emplSrv)
	//comhan := handler.NewAdminCommentHandler(comserv)

	//jobr := employee_repository.NewJobGormRepo(dbconn)
	//jobs := employee_service.NewJobServiceGorm(jobr)

	//jobh := handler.NewJobHandler(jobs)
	//mux := http.NewServeMux()
	router := httprouter.New()

	router.GET("/v2/admin/comments/:id", employeeHandler.GetSingleEmployee)
	router.GET("/v2/admin/comments", employeeHandler.GetEmployees)
	router.PUT("/v1/admin/comments/:id", employeeHandler.PutEmployee)
	router.POST("/v2/Add", employeeHandler.PostEmployee)
	router.DELETE("/v1/admin/comments/:id", employeeHandler.DeleteEmployee)
	//router.POST("/v1/admin/comments", employeeHandler.PostEmployee)
	router.GET("/Index",index)
	//router.GET("/Add",employeeHandler.Add)

	http.ListenAndServe(":8181", router)

}
