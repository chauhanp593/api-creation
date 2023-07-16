package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LineString struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

type Line struct {
	ID   string     `json:"id"`
	Line LineString `json:"line"`
}

type CombinedPayload struct {
	LineString LineString `json:"lineString"`
	Lines      []Line     `json:"lines"`
}

func doSegmentsIntersect(a, b, c, d []float64) bool {
	// Calculate the orientation of three points
	orientation1 := calculateOrientation(a, b, c)
	orientation2 := calculateOrientation(a, b, d)
	orientation3 := calculateOrientation(c, d, a)
	orientation4 := calculateOrientation(c, d, b)

	// Check if the orientations are different (indicating intersection)
	if orientation1 != orientation2 && orientation3 != orientation4 {
		return true
	}

	// Special cases for collinear segments
	if orientation1 == 0 && isOnSegment(a, b, c) {
		return true
	}
	if orientation2 == 0 && isOnSegment(a, b, d) {
		return true
	}
	if orientation3 == 0 && isOnSegment(c, d, a) {
		return true
	}
	if orientation4 == 0 && isOnSegment(c, d, b) {
		return true
	}

	return false
}

func calculateOrientation(a, b, c []float64) int {
	// Calculate the orientation of three points using cross product
	val := (b[1]-a[1])*(c[0]-b[0]) - (b[0]-a[0])*(c[1]-b[1])
	if val == 0 {
		return 0 // Collinear
	} else if val > 0 {
		return 1 // Clockwise
	} else {
		return 2 // Counterclockwise
	}
}

func isOnSegment(a, b, p []float64) bool {
	// Check if point p lies on segment ab
	if p[0] >= min(a[0], b[0]) && p[0] <= max(a[0], b[0]) &&
		p[1] >= min(a[1], b[1]) && p[1] <= max(a[1], b[1]) {
		return true
	}
	return false
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Check authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader != "YOUR_AUTH_TOKEN" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Decode the payload
	var combinedPayload CombinedPayload
	err := json.NewDecoder(r.Body).Decode(&combinedPayload)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Extract the lineString and lines from the payload
	lineString := combinedPayload.LineString
	lines := combinedPayload.Lines

	// Find the intersecting lines
	intersectingLines := make([]string, 0)
	for _, line := range lines {
		lineID := line.ID
		lineCoords := line.Line.Coordinates

		for i := 0; i < len(lineCoords)-1; i++ {
			lineA := lineCoords[i]
			lineB := lineCoords[i+1]
			for _, coord := range lineString.Coordinates {
				if doSegmentsIntersect(lineA, lineB, coord, coord) {
					newString := fmt.Sprintf("%s, %.4f, %.4f", lineID, coord[0], coord[1])
					intersectingLines = append(intersectingLines, newString)
					break
				}
			}
		}
	}

	// Prepare the response
	responseJSON, err := json.Marshal(intersectingLines)
	if err != nil {
		http.Error(w, "Failed to generate response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func main() {
	http.HandleFunc("/", handleRequest)

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
