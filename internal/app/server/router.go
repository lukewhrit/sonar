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
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Start begins the Sonar server app
func Start(port int, hostname string) error {
	app := fiber.New()

	// Register API app midleware
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(etag.New())
	app.Use(logger.New())

	// Register API Routes
	app.Get("/search", searchRoute)

	return app.Listen(hostname + fmt.Sprintf(":%d", port))
}
