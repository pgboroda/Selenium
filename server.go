package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

const userFilename = "users.json"

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, asdld!")
	})
	//добавить пользователя
	e.POST("/users", createUser)

	// найти пользователя по Id
	e.GET("/users", findUser)

	// изменить пользователя
	e.PUT("/users", updateUser)

	// удалить пользователя
	// e.DELETE("/users", deleteUser)

	//получить всех пользователей
	//e.GET("/users")

	e.Logger.Fatal(e.Start(":8080"))

}

// Пуш запрос на добавление юзера
func createUser(c echo.Context) (err error) {
	// ловим запрос в Json
	u := new(User)
	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// вызываем файл и преобразуем
	var users = readJSON()
	// ищем пользователя по ID  из полученного запроса
	for i := 0; i < len(users); i++ {
		if u.Id == users[i].Id {
			return c.JSON(http.StatusBadRequest, "пользователь уже существует")
		}
	}
	//  Если пользователя нет -  то добавляем к файлу в конец
	x := User{
		Id:   u.Id,
		Name: u.Name,
	}
	if x.Name != "" {
		result := append(users, x)
		rawDataOut, err := json.MarshalIndent(&result, "", "  ")
		if err != nil {
			log.Fatal("JSON marshaling failed:", err)
		}

		err = ioutil.WriteFile(userFilename, rawDataOut, 0)
		if err != nil {
			log.Fatal("Cannot write updated settings file:", err)
		}
	}
	// отбивка что пользователь добавлен
	return c.JSON(http.StatusOK, u)
}

// Ищем пользователя по Id
func findUser(c echo.Context) (err error) {

	// ловим запрос в Json
	u := new(User)
	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// вызываем файл и преобразуем
	var users = readJSON()
	// ищем и выводим имя пользователя по ID  из полученного запроса
	for i := 0; i < len(users); i++ {
		if u.Id == users[i].Id {
			return c.JSON(http.StatusOK, users[i])
		}
	}
	return c.JSON(http.StatusBadRequest, "пользователь не существует")
}

// изменить пользователя
func updateUser(c echo.Context) (err error) {
	// ловим запрос в Json
	u := new(User)
	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// вызываем файл и преобразуем
	var users = readJSON()
	// ищем пользователя по ID  из полученного запроса
	for i := 0; i < len(users); i++ {
		if u.Id == users[i].Id {
			x := User{
				Id:   u.Id,
				Name: u.Name,
			}
			// если нашли и полученное имя не пустое - то меняем
			if x.Name != "" {
				users[i].Name = x.Name
				rawDataOut, err := json.MarshalIndent(&users, "", "  ")
				if err != nil {
					log.Fatal("JSON marshaling failed:", err)
				}

				err = ioutil.WriteFile(userFilename, rawDataOut, 0)
				if err != nil {
					log.Fatal("Cannot write updated settings file:", err)
				}
				return c.JSON(http.StatusOK, u)
			} else {
				return c.JSON(http.StatusBadRequest, "Пустое Имя")
			}

		}

	}
	return c.JSON(http.StatusBadRequest, "Нет такого пользователя")
}

// func deleteUser(c echo.Context) error {
// }

// структура файла User
type (
	User struct {
		Id   int    `json:"id"`
		Name string `json:"name" `
	}
)

//вытаскиваем и читаем список имён
func readJSON() []User {
	jsonFile, err := os.Open(userFilename)
	if err != nil {
		fmt.Print(err)
	}
	defer jsonFile.Close()
	data, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(err)
	}
	var result []User

	jsonErr := json.Unmarshal(data, &result)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return result

}
