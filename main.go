package main

import (
	"encoding/json"
	"fmt"
	"os"
	"podval-bans/code/stucts/server"

	"github.com/LittleDrongo/fmn-lib/errors"
	"github.com/LittleDrongo/fmn-lib/utils/files"
	// fmn-lib/utils/json/
)

func main() {

	//читаю конфиг сервера 1
	file, err := os.Open("data/input/1.json")
	errors.Println(err, "Ошибка при открытии файла:")
	defer file.Close()
	var configOne server.ServerConfig
	err = json.NewDecoder(file).Decode(&configOne)
	errors.Println(err, "Ошибка при декодировании JSON файла:")

	//читаю конфиг сервера 2
	file2, err := os.Open("data/input/2.json")
	errors.Println(err, "Ошибка при открытии файла:")
	defer file2.Close()
	var configTwo server.ServerConfig
	err = json.NewDecoder(file2).Decode(&configTwo)
	errors.Println(err, "Ошибка при декодировании JSON файла:")

	//создаю общий банлист
	combinedBans := MergeBans(configOne.Bans, configTwo.Bans)
	fmt.Println("Объединенный список забаненных игроков:")
	PrintJson(combinedBans)

	//Заливаю общий бан лист в конфиги
	configOne.Bans = combinedBans
	configTwo.Bans = combinedBans

	//Сохраняю конфиг в файл json файл
	WriteJson(configOne, "data/output/1.json")
	WriteJson(configTwo, "data/output/2.json")

}

func MergeBans(bans ...server.Bans) server.Bans {

	combinedBans := make(server.Bans)
	for _, ban := range bans {
		for k, v := range ban {
			combinedBans[k] = v
		}
	}
	return combinedBans
}

func WriteJson(data interface{}, filepath string) {

	files.MakeDirIfIsNotExist(filepath)

	file, err := json.MarshalIndent(data, "", "	")
	errors.Println(err, "Ошибка при создании объекта данных JSON")

	err = os.WriteFile(filepath, file, 0644)
	errors.Println(err, "Ошибка сохранения файла JSON")
}

func PrintJson(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	errors.Fatalln(err, "Ошибка при создании объекта данных JSON:")
	fmt.Println(string(jsonData))
}
