package cmd

import "github.com/joho/godotenv"

func loadEnv() error {
	// .env
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}
