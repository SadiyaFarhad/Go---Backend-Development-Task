package handler

import (
	"go-backend-user/internal/models"
	"go-backend-user/internal/service"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// In-memory storage for demo purposes
var users = []models.User{}
var idCounter = 1

// ---------------- CREATE USER ----------------
func CreateUser(c *fiber.Ctx) error {
	var req struct {
		Name string `json:"name"`
		DOB  string `json:"dob"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Parse DOB in "YYYY-MM-DD" format
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid date format, use YYYY-MM-DD",
		})
	}

	u := models.User{
		ID:   idCounter,
		Name: req.Name,
		DOB:  dob,
	}
	idCounter++
	users = append(users, u)

	return c.Status(fiber.StatusCreated).JSON(u)
}

// ---------------- GET USER BY ID ----------------
func GetUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	for _, u := range users {
		if u.ID == id {
			resp := models.UserResponse{
				ID:   u.ID,
				Name: u.Name,
				DOB:  u.DOB.Format("2006-01-02"),
				Age:  service.CalculateAge(u.DOB),
			}
			return c.JSON(resp)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

// ---------------- LIST ALL USERS ----------------
func ListUsers(c *fiber.Ctx) error {
	var resp []models.UserResponse
	for _, u := range users {
		resp = append(resp, models.UserResponse{
			ID:   u.ID,
			Name: u.Name,
			DOB:  u.DOB.Format("2006-01-02"),
			Age:  service.CalculateAge(u.DOB),
		})
	}
	return c.JSON(resp)
}

// ---------------- UPDATE USER ----------------
func UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	var req struct {
		Name string `json:"name"`
		DOB  string `json:"dob"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid date format, use YYYY-MM-DD",
		})
	}

	for i, u := range users {
		if u.ID == id {
			users[i].Name = req.Name
			users[i].DOB = dob
			resp := models.UserResponse{
				ID:   users[i].ID,
				Name: users[i].Name,
				DOB:  users[i].DOB.Format("2006-01-02"),
				Age:  service.CalculateAge(users[i].DOB),
			}
			return c.JSON(resp)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

// ---------------- DELETE USER ----------------
func DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}
