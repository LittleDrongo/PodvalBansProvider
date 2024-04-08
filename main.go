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

	file, err := os.Open("data/1.json")

	errors.Println(err, "Ошибка при открытии файла:")
	defer file.Close()
	var configOne server.ServerConfig
	err = json.NewDecoder(file).Decode(&configOne)
	errors.Println(err, "Ошибка при декодировании JSON файла:")

	file2, err := os.Open("data/2.json")
	errors.Println(err, "Ошибка при открытии файла:")
	defer file2.Close()
	var configTwo server.ServerConfig
	err = json.NewDecoder(file2).Decode(&configTwo)
	errors.Println(err, "Ошибка при декодировании JSON файла:")

	combinedBans := MergeBans(configOne.Bans, configTwo.Bans)
	fmt.Println("Объединенный список забаненных игроков:")
	Print(combinedBans)

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

func Write(data interface{}, filepath string) {

	files.MakeDirIfIsNotExist(filepath)

	file, err := json.MarshalIndent(data, "", "	")
	errors.Println(err, "Ошибка при создании объекта данных JSON")

	err = os.WriteFile(filepath, file, 0644)
	errors.Println(err, "Ошибка сохранения файла JSON")
}

func Print(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	errors.Fatalln(err, "Ошибка при создании объекта данных JSON:")
	fmt.Println(string(jsonData))
}
