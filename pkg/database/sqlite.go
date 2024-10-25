package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
	conn *sql.DB
}

// Init - инициализация базы данных и создание таблиц
func (db *SQLiteDB) Init() error {
	var err error
	db.conn, err = sql.Open("sqlite3", "management.db")
	if err != nil {
		return fmt.Errorf("ошибка подключения к SQLite: %w", err)
	}

	if err := db.createTables(); err != nil {
		return fmt.Errorf("ошибка создания таблиц: %w", err)
	}
	log.Println("SQLite: база данных инициализирована")
	return nil
}

// Close - закрытие соединения с базой данных
func (db *SQLiteDB) Close() error {
	if db.conn != nil {
		return db.conn.Close()
	}
	return nil
}

// Создание таблиц для SQLite
func (db *SQLiteDB) createTables() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS spheres (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            description TEXT
        );`,
		`CREATE TABLE IF NOT EXISTS services (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            description TEXT,
            price REAL,
            duration INTEGER,
            sphere_id INTEGER,
            FOREIGN KEY (sphere_id) REFERENCES spheres(id)
        );`,
		`CREATE TABLE IF NOT EXISTS records (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            date TEXT NOT NULL,
            time TEXT NOT NULL,
            service_id INTEGER,
            client_id INTEGER,
            price REAL,
            duration INTEGER,
            paid BOOLEAN,
            FOREIGN KEY (service_id) REFERENCES services(id)
        );`,
		`CREATE TABLE IF NOT EXISTS schedules (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            day_of_week INTEGER,
            start_time TEXT,
            end_time TEXT,
            is_holiday BOOLEAN
        );`,
	}
	for _, query := range queries {
		if _, err := db.conn.Exec(query); err != nil {
			return err
		}
	}
	return nil
}

// AddSphere - добавление записи в таблицу сфер деятельности
func (db *SQLiteDB) AddSphere(name, description string) error {
	_, err := db.conn.Exec("INSERT INTO spheres (name, description) VALUES (?, ?)", name, description)
	return err
}

// AddService - добавление услуги
func (db *SQLiteDB) AddService(name, description string, price float64, duration int, sphereID int) error {
	_, err := db.conn.Exec("INSERT INTO services (name, description, price, duration, sphere_id) VALUES (?, ?, ?, ?, ?)",
		name, description, price, duration, sphereID)
	return err
}

// AddRecord - добавление записи
func (db *SQLiteDB) AddRecord(date, time string, serviceID, clientID int, price float64, duration int, paid bool) error {
	_, err := db.conn.Exec("INSERT INTO records (date, time, service_id, client_id, price, duration, paid) VALUES (?, ?, ?, ?, ?, ?, ?)",
		date, time, serviceID, clientID, price, duration, paid)
	return err
}

// AddSchedule - добавление графика работы
func (db *SQLiteDB) AddSchedule(dayOfWeek int, startTime, endTime string, isHoliday bool) error {
	_, err := db.conn.Exec("INSERT INTO schedules (day_of_week, start_time, end_time, is_holiday) VALUES (?, ?, ?, ?)",
		dayOfWeek, startTime, endTime, isHoliday)
	return err
}
