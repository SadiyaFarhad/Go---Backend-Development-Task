package service

import "time"

func CalculateAge(dob time.Time) int {
    today := time.Now()
    age := today.Year() - dob.Year()
    if today.YearDay() < dob.YearDay() {
        age--
    }
    return age
}
