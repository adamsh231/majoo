package domain

import (
	"database/sql"
	"github.com/adamsh231/majoo/packages/helper"
	"github.com/adamsh231/majoo/packages/jwe"
	"github.com/adamsh231/majoo/packages/jwt"
	"github.com/adamsh231/majoo/packages/postgres"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
	jwtFiber "github.com/gofiber/jwt/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	PostgresDB    *sql.DB
	Validator     *validator.Validate
	JweCredential jwe.Credential
	JwtCredential jwt.JwtCredential
	JwtConfig     jwtFiber.Config
}

var (
	ValidatorDriver *validator.Validate
	Uni             *ut.UniversalTranslator
	Translator      ut.Translator
)

func LoadConfig() (res Config, err error) {
	err = godotenv.Load("../../.env")
	helper.Log(err, "Load Env")

	// postgres connection
	PostgresInfo := postgresql.Connection{
		Host:                    os.Getenv("DB_HOST"),
		DbName:                  os.Getenv("DB_NAME"),
		User:                    os.Getenv("DB_USERNAME"),
		Password:                os.Getenv("DB_PASSWORD"),
		Port:                    os.Getenv("DB_PORT"),
		SslMode:                 os.Getenv("DB_SSL_MODE"),
		DBMaxConnection:         helper.StringToInt(os.Getenv("DB_MAX_CONNECTION")),
		DBMAxIdleConnection:     helper.StringToInt(os.Getenv("DB_MAX_IDLE_CONNECTION")),
		DBMaxLifeTimeConnection: helper.StringToInt(os.Getenv("DB_MAX_LIFETIME_CONNECTION")),
	}
	res.PostgresDB, err = PostgresInfo.DbConnect()
	if err != nil {
		log.Fatal(err.Error())
	}

	//jwe credential
	res.JweCredential = jwe.Credential{
		KeyLocation: os.Getenv("JWE_PRIVATE_KEY"),
		Passphrase:  os.Getenv("JWE_PRIVATE_KEY_PASSPHRASE"),
	}

	//jwt credential
	res.JwtCredential = jwt.JwtCredential{
		TokenSecret:         os.Getenv("SECRET"),
		ExpiredToken:        helper.StringToInt(os.Getenv("TOKEN_EXP_TIME")),
		RefreshTokenSecret:  os.Getenv("SECRET_REFRESH_TOKEN"),
		ExpiredRefreshToken: helper.StringToInt(os.Getenv("REFRESH_TOKEN_EXP_TIME")),
	}

	//jwt config
	res.JwtConfig = jwtFiber.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		Claims:     &jwt.CustomClaims{},
	}

	// validator
	res.Validator = ValidatorDriver

	return res, err
}

func ValidatorInit() {
	en := en.New()
	id := id.New()
	Uni = ut.New(en, id)

	transEN, _ := Uni.GetTranslator("en")
	transID, _ := Uni.GetTranslator("id")

	ValidatorDriver = validator.New()

	enTranslations.RegisterDefaultTranslations(ValidatorDriver, transEN)
	idTranslations.RegisterDefaultTranslations(ValidatorDriver, transID)

	switch os.Getenv("APP_LOCALE") {
	case "en":
		Translator = transEN
	case "id":
		Translator = transID
	}
}
