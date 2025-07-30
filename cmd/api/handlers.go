package main

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Hello, world! ", app.Domain)

	// Json Data
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Movies up and running",
		Version: "1.0.0",
	}

	// out, err := json.Marshal(payload)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(out)

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About, here!")
}

// ฟังก์ชันสำหรับแสดงรายชื่อหนังทั้งหมดโดยการทดสอบ mock data
func (app *application) AllDemoMovies(w http.ResponseWriter, r *http.Request) {

	// สร้างตัวแปรไว้เก็บข้อมูลหนัง
	var movies []models.Movie

	// กำหนดตัวแปรรูปแบบวันที่ yyyy-mm-dd
	rd, _ := time.Parse("2006-01-02", "1981-06-12")

	// สร้างข้อมูลหนัง
	highlander := models.Movie{
		ID:          1,
		Title:       "Highlander",
		ReleaseDate: rd,
		RunTime:     116,
		MPAARating:  "R",
		Description: "A very nice movie",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// เพิ่มข้อมูลหนังลงใน slice
	movies = append(movies, highlander)

	rd, _ = time.Parse("2006-01-02", "1982-06-07")

	rotla := models.Movie{
		ID:          2,
		Title:       "Raiders of the Lost Ark",
		ReleaseDate: rd,
		MPAARating:  "PG-13",
		RunTime:     115,
		Description: "Another very nice movie",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, rotla)

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}

// ฟังก์ชันสำหรับแสดงรายชื่อหนังทั้งหมดโดยเรียกจาก database
func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	// ดึงข้อมูลหนังทั้งหมดจาก database โดยใช้เมธอด AllMovies จาก app.DB.
	movies, err := app.DB.AllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, movies)
}
