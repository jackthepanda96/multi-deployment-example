package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DBUSER string
	DBHOST string
	DBPASS string
	DBNAME string
	DBPORT uint16
	// RHOST     string
	// RPORT     uint16
	// RPASSWORD string
	// MUSER     string
	// MHOST     string
	// MPASS     string
	// MPORT     uint16
}

func InitConfig() *AppConfig {
	var response = new(AppConfig)
	response = readData()
	return response
}

func readData() *AppConfig {
	var data = new(AppConfig)

	data = readEnv()

	if data == nil {
		err := godotenv.Load(".env")
		data = readEnv()
		if err != nil || data == nil {
			return nil
		}
	}

	return data
}

func readEnv() *AppConfig {
	var data = new(AppConfig)
	var permit = true

	if val, found := os.LookupEnv("DBUSER"); found {
		data.DBUSER = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("DBPASS"); found {
		data.DBPASS = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("DBHOST"); found {
		data.DBHOST = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("DBPORT"); found {
		cnv, err := strconv.Atoi(val)
		if err != nil {
			permit = false
		}

		data.DBPORT = uint16(cnv)
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		data.DBNAME = val
	} else {
		permit = false
	}

	// if val, found := os.LookupEnv("RHOST"); found {
	// 	data.RHOST = val
	// } else {
	// 	permit = false
	// }

	// if val, found := os.LookupEnv("RPORT"); found {
	// 	cnv, err := strconv.Atoi(val)
	// 	if err != nil {
	// 		permit = false
	// 	}

	// 	data.RPORT = uint16(cnv)
	// } else {
	// 	permit = false
	// }

	// if val, found := os.LookupEnv("RPASSWORD"); found {
	// 	data.RPASSWORD = val
	// } else {
	// 	permit = false
	// }

	// if val, found := os.LookupEnv("MUSER"); found {
	// 	data.MUSER = val
	// } else {
	// 	permit = false
	// }

	// if val, found := os.LookupEnv("MPASS"); found {
	// 	data.MPASS = val
	// } else {
	// 	permit = false
	// }

	// if val, found := os.LookupEnv("MHOST"); found {
	// 	data.MHOST = val
	// } else {
	// 	permit = false
	// }

	// if val, found := os.LookupEnv("MPORT"); found {
	// 	cnv, err := strconv.Atoi(val)
	// 	if err != nil {
	// 		permit = false
	// 	}

	// 	data.MPORT = uint16(cnv)
	// } else {
	// 	permit = false
	// }

	if !permit {
		return nil
	}

	return data
}
