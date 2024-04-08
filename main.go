package main

import (
	"encoding/json"
	"fmt"
	"os"
	"podval-bans/code/stucts/server"
	"strings"

	"github.com/LittleDrongo/fmn-lib/console/color"
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

	//читаю конфиг сервера 3
	file3, err := os.Open("data/input/3.json")
	errors.Println(err, "Ошибка при открытии файла:")
	defer file3.Close()
	var configThree server.ServerConfig
	err = json.NewDecoder(file3).Decode(&configThree)
	errors.Println(err, "Ошибка при декодировании JSON файла:")

	//создаю общий банлист
	combinedBans := MergeBans(configOne.Bans, configTwo.Bans, configThree.Bans)

	fmt.Println(color.BG_GREEN, "Объединенный список забаненных игроков:", color.BG_RESET)
	fmt.Println(color.GREEN)
	PrintJson(combinedBans)
	fmt.Println(color.RESET)
	//Заливаю общий бан лист в конфиги
	configOne.Bans = combinedBans
	configTwo.Bans = combinedBans
	configThree.Bans = combinedBans

	//Сохраняю конфиг в файл json файл
	WriteJsonWithFormatHMTL(configOne, "data/output/1.json")
	WriteJsonWithFormatHMTL(configTwo, "data/output/2.json")
	WriteJsonWithFormatHMTL(configThree, "data/output/3.json")
	WriteJsonWithFormatHMTL(combinedBans, "data/output/banlist.json")

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

func WriteJsonWithFormatHMTL(data interface{}, filepath string) {
	files.MakeDirIfIsNotExist(filepath)

	jsonData, err := json.MarshalIndent(data, "", "    ")
	errors.Println(err, "Ошибка при создании объекта данных JSON")

	// Заменяю символы "<" и ">" на их HTML-эквиваленты
	jsonString := string(jsonData)
	jsonString = strings.Replace(jsonString, "\\u003c", "<", -1)
	jsonString = strings.Replace(jsonString, "\\u003e", ">", -1)

	// jsonString = strings.Replace(jsonString, `\`, `\\`, -1)
	// jsonString = strings.Replace(jsonString, `"`, `\"`, -1)
	// jsonString = strings.Replace(jsonString, "\n", `\n`, -1)
	// jsonString = strings.Replace(jsonString, "\r", `\r`, -1)
	// jsonString = strings.Replace(jsonString, "\t", `\t`, -1)

	err = os.WriteFile(filepath, []byte(jsonString), 0644)
	errors.Println(err, "Ошибка сохранения файла JSON")
}

func PrintJson(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	errors.Fatalln(err, "Ошибка при создании объекта данных JSON:")
	fmt.Println(string(jsonData))
}
