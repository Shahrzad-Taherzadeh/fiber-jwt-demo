package race

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

var counter int
var mu sync.Mutex

func RaceHandler(c *fiber.Ctx) error {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	return c.JSON(fiber.Map{"counter": counter})
}