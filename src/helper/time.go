package helper

import (
	"time"
)

var (
	loc      *time.Location
	Now      = now
	months   = []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	weekdays = []string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}
)

func now() time.Time {
	return time.Now().In(loc)
}

func InitTime() (err error) {
	loc, err = time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return err
	}
	return nil
}

func GetLocation() *time.Location {
	return loc
}
