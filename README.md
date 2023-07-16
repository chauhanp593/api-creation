# Line Intersection API

The Line Intersection API is a RESTful web service that allows you to find intersecting lines between a given LineString and a set of lines. The API takes a LineString in GeoJSON format and a set of lines with start and end coordinates, and it returns the lines that intersect with the LineString.

## Endpoints

### POST /intersect

This endpoint finds the intersecting lines between the LineString and the provided set of lines.

#### Request Headers

- **Authorization**: The authorization token for accessing the API. (Replace "YOUR_AUTH_TOKEN" with your actual authorization token)
- **400 Bad Request**: If the request body is invalid or missing.
- **401 Unauthorized**: If the provided authorization token is incorrect or missing.

#### Getting Started

To run the Line Intersection API, follow these steps:

1. Clone this repository.
2. Install Go and any required dependencies.
3. Set the authorization token in the code (`YOUR_AUTH_TOKEN` in `handleRequest` function).
4. Build and run the Go program using `go run main.go`.
5. The API will be accessible at `http://localhost:8080`.

#### Dependencies

The Line Intersection API uses the following dependencies:

- Go (version 1.16 or later)
- Go packages: `github.com/gorilla/mux`, `github.com/gorilla/handlers`

#### Contributing

Contributions are welcome! If you find any issues or would like to add new features, please open an issue or submit a pull request.



#### Request Body

The request body should be a JSON object with the following structure:

```json
[{
  "lineString": {
    "type": "LineString",
    "coordinates": [
      [longitude1, latitude1],
      [longitude2, latitude2],
      ...
    ]
  },
  "lines": [
    {
      "id": "line1",
      "line": {
        "type": "LineString",
        "coordinates": [
          [startLongitude1, startLatitude1],
          [endLongitude1, endLatitude1]
        ]
      }
    },
    {
      "id": "line2",
      "line": {
        "type": "LineString",
        "coordinates": [
          [startLongitude2, startLatitude2],
          [endLongitude2, endLatitude2]
        ]
      }
    },
    ...
  ]
}

## Response
The API responds with a JSON array containing the line IDs and the coordinates where each line intersects with the LineString:
```json
[
  "line1, latitude1, longitude1",
  "line2, latitude2, longitude2",
  ...
]
