package database

import "errors"

// Database - интерфейс для взаимодействия с базой данных
type Database interface {
	Init() error
	Close() error
	AddSphere(name, description string) error
	AddService(name, description string, price float64, duration int, sphereID int) error
	AddRecord(date, time string, serviceID, clientID int, price float64, duration int, paid bool) error
	AddSchedule(dayOfWeek int, startTime, endTime string, isHoliday bool) error
}

var ErrNotImplemented = errors.New("method not implemented")

// NewDatabase - создает экземпляр нужной базы данных
func NewDatabase(dbType string) (Database, error) {
	switch dbType {
	case "sqlite":
		return &SQLiteDB{}, nil
	//case "postgres":
	//	return &PostgresDB{}, nil
	default:
		return nil, errors.New("неизвестный тип базы данных")
	}
}
