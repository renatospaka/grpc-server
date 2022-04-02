package main

import (
	"net/http"
	"server/model"
	"encoding/json"
)

var courseList = model.NewCourses()

func main() {
	http.HandleFunc("/course", CourseListHandler)
	http.ListenAndServe(":8000", nil)
}

func CourseListHandler(w http.ResponseWriter, r *http.Request) {
	courseJson, _ := json.Marshal(courseList)
	w.Write([]byte(courseJson))
}
