package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"leo-garay/go-tasks/data"
	"leo-garay/go-tasks/models"
	"net/http"
)
 

func Create(w http.ResponseWriter, r *http.Request) {
	var job models.Job
 
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Data")
	}

	json.Unmarshal(reqBody, &job)
	data.JobsData = append(data.JobsData, job) 

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(job)

}

func CreateError(w http.ResponseWriter, r *http.Request) {
	 	 http.Error(w, "Conflict", http.StatusConflict) 	   
}
func GetJobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.JobsData)
}