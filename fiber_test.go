package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var engine = mustache.New("./template", ".mustache")
var app = fiber.New(fiber.Config{
	Views: engine,
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		c.Status(http.StatusInternalServerError)
		return c.SendString("Error: " + err.Error())
	},
})

func TestRoutingHelloWorld(t *testing.T) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	request := httptest.NewRequest("GET", "/", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello, World!", string(bytes))
}

func TestCtx(t *testing.T) {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Query("name", "Guest")
		return c.SendString("Hello " + name)
	})

	request := httptest.NewRequest("GET", "/hello?name=Dika", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Dika", string(bytes))

	request = httptest.NewRequest("GET", "/hello", nil)
	response, err = app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err = io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Guest", string(bytes))
}

func TestHttpRequest(t *testing.T) {
	app.Get("/request", func(c *fiber.Ctx) error {
		first := c.Get("firstname")
		last := c.Cookies("lastname")
		return c.SendString("Hello " + first + " " + last)
	})

	request := httptest.NewRequest("GET", "/request", nil)
	request.Header.Set("firstname", "Dika")
	request.AddCookie(&http.Cookie{Name: "lastname", Value: "Koko"})
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Dika Koko", string(bytes))
}

func TestRouteParameter(t *testing.T) {
	app.Get("/users/:userId/orders/:orderId", func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		orderId := c.Params("orderId")
		return c.SendString("Get Order " + orderId + " from user " + userId)
	})

	request := httptest.NewRequest("GET", "/users/dika/orders/10", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Get Order 10 from user dika", string(bytes))
}

func TestFormRequest(t *testing.T) {
	app.Post("/hello", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		return c.SendString("Hello " + name)
	})

	body := strings.NewReader("name=Dika")
	request := httptest.NewRequest("POST", "/hello", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Dika", string(bytes))
}

//go:embed source/sample.txt
var sampleFile []byte

func TestFormUpload(t *testing.T) {
	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		err = c.SaveFile(file, "./target/"+file.Filename)
		if err != nil {
			return err
		}

		return c.SendString("Upload file to target " + file.Filename + " successfully")
	})

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	file, err := writer.CreateFormFile("file", "sample.txt")
	assert.Nil(t, err)
	file.Write(sampleFile)
	writer.Close()
	request := httptest.NewRequest("POST", "/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Upload file to target sample.txt successfully", string(bytes))
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestRequestBody(t *testing.T) {
	app.Post("/login", func(c *fiber.Ctx) error {
		body := c.Body()

		request := new(LoginRequest)
		err := json.Unmarshal(body, request)
		if err != nil {
			return err
		}
		return c.SendString("Hello " + request.Username)
	})

	body := strings.NewReader(`{"username":"Dika","password":"password111221"}`)
	request := httptest.NewRequest("POST", "/login", body)
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Dika", string(bytes))
}

type RegisterRequest struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password"  xml:"password" form:"password"`
	Name     string `json:"name"  xml:"name" form:"name"`
}

func TestBodyParser(t *testing.T) {
	app.Post("/register", func(c *fiber.Ctx) error {
		request := new(RegisterRequest)
		err := c.BodyParser(request)
		if err != nil {
			return err
		}
		return c.SendString("Register " + request.Username + " successfully")
	})
}

func TestBodyParserJSON(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(`{"username":"Dika","password":"password111221","name":"Dika Koko"}`)
	request := httptest.NewRequest("POST", "/register", body)
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Dika successfully", string(bytes))
}

func TestBodyParserFORM(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(`username=Dika&password=password111221&name=Dika+Koko`)
	request := httptest.NewRequest("POST", "/register", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Dika successfully", string(bytes))
}

func TestBodyParserXML(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(`
			<RegisterRequest>
				<username>Dika</username>
				<password>password111221</password>
				<name>Dika Koko</name>
			</RegisterRequest>
		`)
	request := httptest.NewRequest("POST", "/register", body)
	request.Header.Set("Content-Type", "application/xml")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Dika successfully", string(bytes))
}

func TestResponseJSON(t *testing.T) {
	app.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"username": "Dika",
			"name":     "Dika koko",
		})
	})

	request := httptest.NewRequest("GET", "/user", nil)
	request.Header.Set("Accept", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"name":"Dika koko","username":"Dika"}`, string(bytes))
}

func TestDownloadFile(t *testing.T) {
	app.Get("/download", func(c *fiber.Ctx) error {
		return c.Download("./source/sample.txt", "sample.txt")
	})

	request := httptest.NewRequest("GET", "/download", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)
	assert.Equal(t, `attachment; filename="sample.txt"`, response.Header.Get("Content-Disposition"))

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Sample file upload!\n", string(bytes))
}

func TestRoutingGroup(t *testing.T) {
	helloWorld := func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	}

	api := app.Group("/api")
	api.Get("/hello", helloWorld) // /api/hello
	api.Get("/world", helloWorld) // /api/world

	web := app.Group("/web")
	web.Get("/hello", helloWorld) // /web/hello
	web.Get("/world", helloWorld) // /web/world

	request := httptest.NewRequest("GET", "/api/hello", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World", string(bytes))
}

func TestStatic(t *testing.T) {
	app.Static("/public", "./source")

	request := httptest.NewRequest("GET", "/public/sample.txt", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Sample file upload!\n", string(bytes))
}

func TestErrorHandler(t *testing.T) {
	app.Get("/error", func(c *fiber.Ctx) error {
		return errors.New("ups")
	})

	request := httptest.NewRequest("GET", "/error", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 500, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Error: ups", string(bytes))
}

func TestView(t *testing.T) {
	app.Get("/view", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title":   "Hello Title",
			"header":  "Hello Header",
			"content": "Hello Content",
		})
	})

	request := httptest.NewRequest("GET", "/view", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(bytes), "Hello Title")
	assert.Contains(t, string(bytes), "Hello Header")
	assert.Contains(t, string(bytes), "Hello Content")
}

func TestClient(t *testing.T) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	agent := client.Get("https://example.com")
	status, response, errors := agent.String()
	fmt.Println(status, response)
	assert.Nil(t, errors)
	assert.Equal(t, 200, status)
	assert.Contains(t, response, "Example Domain")
}
