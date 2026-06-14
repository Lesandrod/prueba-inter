package handlers

import (
	"bytes"
	"encoding/json"
	"go-api/internal/services"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

type QRRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

type StatsRequest struct {
	Q [][]float64 `json:"q"`
	R [][]float64 `json:"r"`
}

type StatsResponse struct {
	Max               float64 `json:"max"`
	Min               float64 `json:"min"`
	Average           float64 `json:"average"`
	Total             float64 `json:"total"`
	HasDiagonalMatrix bool    `json:"hasDiagonalMatrix"`
}

func QRHandler(c *fiber.Ctx) error {

	var req QRRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "body invalido",
		})
	}

	result, err := services.FactorizeQR(req.Matrix)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	stats, err := fetchStatsFromNode(result.Q, result.R)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "error al consultar node",
		})
	}

	return c.JSON(fiber.Map{
		"q":     result.Q,
		"r":     result.R,
		"stats": stats,
	})
}

func fetchStatsFromNode(q [][]float64, r [][]float64) (StatsResponse, error) {

	payload := StatsRequest{
		Q: q,
		R: r,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return StatsResponse{}, err
	}

	nodeURL := os.Getenv("NODE_STATS_URL")
	if nodeURL == "" {
		// valor por defecto pensado para Docker compose
		nodeURL = "http://node-api:3000/stats"
	}

	resp, err := http.Post(
		nodeURL,
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return StatsResponse{}, err
	}

	defer resp.Body.Close()

	var stats StatsResponse

	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return StatsResponse{}, err
	}

	return stats, nil
}
