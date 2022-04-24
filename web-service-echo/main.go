package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type course struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type allCourses []course

var courses = allCourses{
	{
		Id:   1,
		Name: "Curso Práctico de Go",
	},
	{
		Id:   2,
		Name: "Curso de Docker",
	},
	{
		Id:   3,
		Name: "Curso de Git",
	},
}

func main() {
	e := echo.New()

	e.GET("/courses", getAllCourses)
	e.GET("/courses/:id", getCourseById)
	e.POST("/courses", createCourse)
	e.PUT("/courses/:id", modifyCourse)
	e.DELETE("/courses/:id", deleteCourseById)

	e.Start(":2000")
}

func getAllCourses(c echo.Context) error {
	return c.JSON(http.StatusOK, courses)
}

func getCourseById(c echo.Context) error {
	id := c.Param("id")

	for _, courseitem := range courses {
		if strconv.Itoa(courseitem.Id) == id {
			return c.JSON(http.StatusOK, courseitem)
		}
	}
	return c.String(http.StatusBadRequest, "The indicated course doesn't exist.")
}

func createCourse(c echo.Context) error {
	new_course := new(course)
	if err := c.Bind(new_course); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	courses = append(courses, *new_course)

	return c.JSON(http.StatusOK, courses)
}

func modifyCourse(c echo.Context) error {
	updated_course := new(course)
	id := c.Param("id")

	if err := c.Bind(updated_course); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	for i, course := range courses {
		if strconv.Itoa(course.Id) == id {
			courses = append(courses[:i], courses[i+1:]...)
			courses = append(courses, *updated_course)

			return c.JSON(http.StatusOK, courses)
		}
	}
	return c.String(http.StatusBadRequest, "The indicated course doesn't exist.")
}

func deleteCourseById(c echo.Context) error {
	id := c.Param("id")

	for i, courseitem := range courses {
		if strconv.Itoa(courseitem.Id) == id {
			courses = append(courses[:i], courses[i+1:]...)
			return c.JSON(http.StatusOK, courses)
		}
	}
	return c.String(http.StatusBadRequest, "The indicated course doesn't exist.")
}
