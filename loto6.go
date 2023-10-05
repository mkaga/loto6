package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"

	"time"

	"math/rand"
)

func main() {
	e := echo.New()

	e.GET("/", loto6)

	e.Logger.Fatal(e.Start(":1323"))
}

func loto6(c echo.Context) error {
	var rokuyo [6]string = [6]string{"大安", "赤口", "先勝", "友引", "先負", "仏滅"}
	var number [43]int = [43]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
		31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43}

	location, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(location)
	rokuyoInx := (int(now.Month()) + now.Day()) % 6
	var randomArray [7]string
	rand.Shuffle(len(number), func(i, j int) {
		number[i], number[j] = number[j], number[i]
	})

	for i := 0; i < 7; i++ {
		if i < 6 {
			randomArray[i] = strconv.Itoa(number[i+rokuyoInx])
		} else {
			randomArray[i] = "(" + strconv.Itoa(number[i+rokuyoInx]) + ")"
		}
	}
	return c.String(http.StatusOK, "本日は"+rokuyo[rokuyoInx]+"です。 : "+strings.Join(randomArray[:], ","))
}
