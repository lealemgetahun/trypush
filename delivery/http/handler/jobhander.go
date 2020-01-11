package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gitlab.com/username/online-service-and-customer-care/employee"
)

// JobHandler handles comment related http requests
type JobHandler struct {
	jobService employee.JobService
	//tmpl        *template.Template
}

// NewJobHandler returns new AdminCommentHandler object
func NewJobHandler(jobService employee.JobService) *JobHandler {
	return &JobHandler{jobService: jobService}
}

// GetJob handles GET /v1/admin/comments request
func (jh *JobHandler) GetJob(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	job, errs := jh.jobService.Job()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(job, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
