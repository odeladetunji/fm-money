module money.com/kafkaservice

go 1.18

replace money.com/entity => ../entity

replace money.com/dto => ../dto

replace money.com/repository => ../repository

replace money.com/migration => ../migration

require (
	github.com/segmentio/kafka-go v0.4.47
	money.com/dto v0.0.0-00010101000000-000000000000
	money.com/entity v0.0.0-00010101000000-000000000000
	money.com/repository v0.0.0-00010101000000-000000000000
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	gorm.io/driver/postgres v1.5.7 // indirect
	gorm.io/gorm v1.25.10 // indirect
	money.com/migration v0.0.0-00010101000000-000000000000 // indirect
)
