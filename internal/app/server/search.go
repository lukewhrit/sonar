/*
 * Copyright 2021 Luke Whrit

 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 *     http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lukewhrit/sonar/internal/pkg/domain"
)

// IResponse is the response object for the GET /search route
type IResponse struct {
	Query            string           `json:"query"`
	EstimatedResults int              `json:"estimated_results"`
	Results          []domain.IResult `json:"results"`
}

// GET /search?q=string
func searchRoute(c *fiber.Ctx) error {
	query := c.Query("q")

	return c.JSON(map[string]interface{}{
		"type": "search",
		"payload": IResponse{
			Query:            query,
			EstimatedResults: 0,
			Results:          []domain.IResult{},
		},
	})
}
