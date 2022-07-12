package domain

import "time"

type User struct {
	ID          int64     `json:"id"`
	Email       string    `json:"email"`
	Passsword   string    `json:"passsword"`
	PhoneNumber int       `json:"phone_number"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	DateBirth   time.Time `json:"date_birth"`
	TimeBirth   time.Time `json:"time_birth"`
	CityBirth   string    `json:"city_birth"`
	ZodiacSigh  string    `json:"zodiac_sigh"`
}

type RegistrationUser struct {
	Email     string `json:"email"`
	Passsword string `json:"passsword"`
}

type AuthorizationUser struct {
	Email     string `json:"email"`
	Passsword string `json:"passsword"`
}

//
//type ZodiacNatalMap struct {
//	Sun     string
//	Moon    string
//	Mercury string
//	Venus   string
//	Mars    string
//	Jupiter string
//	Saturn  string
//	Uranus  string
//	Neptune string
//	Pluto   string
//	Lilith  string
//	NNode   string
//}
//
//type PlacidusOrbNatalMap struct {
//	OneAsc string
//	Two    string
//	Three  string
//	Four   string
//	Five   string
//	Six    string
//	Seven  string
//	Eighth string
//	Nine   string
//	Ten    string
//	Eleven string
//	Twelve string
//}
//
//type HousesNatalMap struct {
//	Sun     string
//	Moon    string
//	Mercury string
//	Venus   string
//	Mars    string
//	Jupiter string
//	Saturn  string
//	Uranus  string
//	Neptune string
//	Pluto   string
//	Lilith  string
//	NNode   string
//}
//
//type TemperNatalMap struct {
//	Masculine string
//	Feminine  string
//	Cardinal  string
//	Fixed     string
//	Mutable   string
//	Fire      string
//	Earth     string
//	Air       string
//	Water     string
//}
