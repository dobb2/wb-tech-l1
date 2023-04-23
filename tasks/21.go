package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.

Адаптер — это структурный паттерн, который позволяет подружить несовместимые объекты.

Шаблон проектирования «Адаптер» позволяет использовать интерфейс существующего класса как другой интерфейс.
Этот шаблон часто применяется для обеспечения работы одних классов с другими без изменения их исходного кода.
*/

// Структура, которая получив данные от хранилища, выполняет какую-то бизнес-логику
type Client struct {
}

func (c Client) GetDataInto(r Repository) {
	fmt.Println("Get for client")
	r.WriteData()
}

// Интерфейс(контракт), по котрому клиент может работать с различными хранилищами
type Repository interface {
	WriteData()
}

// Хранилище кэш, которое реализует интерфейс Storage, с ним клиент может работать без адаптера
type Cache struct {
}

func (c Cache) WriteData() {
	fmt.Println("all data from cache received")
}

// Хранилище, с которым напрямую клиент работать не может(так как он не реализует интерфейс хранилища), но очень хочет
type PostgresDB struct{}

func (w PostgresDB) SelectData() {
	fmt.Println("all data from postgres received")
}

// Адаптер, реализует интерфейс Storage и при этом в него встроена ссылка на структуру PostgresDB,
// тем самым мы можем вызывать в методах адаптера методы PostgresDB(обертка)
type StorageAdapter struct {
	PostgreStorage *PostgresDB
}

func (s StorageAdapter) WriteData() {
	fmt.Println("adapter allows you to receive data from postgres")
	s.PostgreStorage.SelectData()
}

func main() {
	client := &Client{}
	cache := &Cache{}

	client.GetDataInto(cache)
	fmt.Println()

	postgrestorage := &PostgresDB{}
	postgresAdapter := &StorageAdapter{
		PostgreStorage: postgrestorage,
	}
	//client.GetDataInto(postgrestorage) так мы не сможем получить данные с PostgresDB
	client.GetDataInto(postgresAdapter)
}
