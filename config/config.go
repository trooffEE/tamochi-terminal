package config

import "os"

/*
* TODO Временное решение для целей разработки. Посколько изначально разработка идёт frontend части, мы пока
* не думаем над реализацией авторизации, однако закладываем на это определенную проверку, а значит тестовые данные
* не будут лишними
 */
type config struct {
	DevLogin    string
	DevPassword string
}

func NewConfig() config {
	return config{
		DevLogin:    os.Getenv("DEV_LOGIN"),
		DevPassword: os.Getenv("DEV_PASSWORD"),
	}
}
