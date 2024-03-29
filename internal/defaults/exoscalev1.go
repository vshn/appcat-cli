package defaults

import (
	exoscalev1 "github.com/vshn/appcat/v4/apis/exoscale/v1"
)

func (d *Defaults) GetExoscalePostgreSQLDefault() *exoscalev1.ExoscalePostgreSQL {
	var postgreSQLdefault exoscalev1.ExoscalePostgreSQL

	postgreSQLdefault.Spec.Parameters.Service.MajorVersion = "14"
	postgreSQLdefault.Spec.Parameters.Service.Zone = "ch-dk-2"
	postgreSQLdefault.Spec.Parameters.Size.Plan = "hobbyist-2"
	postgreSQLdefault.Spec.Parameters.Backup.TimeOfDay = "12:00:00"
	postgreSQLdefault.Spec.Parameters.Maintenance.DayOfWeek = "sunday"
	postgreSQLdefault.Spec.Parameters.Maintenance.TimeOfDay = "00:00:00"

	return &postgreSQLdefault
}

func (d *Defaults) GetExoscaleRedisDefault() *exoscalev1.ExoscaleRedis {
	var redisDefault exoscalev1.ExoscaleRedis

	redisDefault.Spec.Parameters.Maintenance.DayOfWeek = "sunday"
	redisDefault.Spec.Parameters.Maintenance.TimeOfDay = "00:00:00"
	redisDefault.Spec.Parameters.Service.Zone = "ch-dk-2"

	return &redisDefault
}

func (d *Defaults) GetExoscaleKafkaDefault() *exoscalev1.ExoscaleKafka {
	var kafkaDefault exoscalev1.ExoscaleKafka

	kafkaDefault.Spec.Parameters.Service.Version = "3.4.0"
	kafkaDefault.Spec.Parameters.Service.Zone = "ch-dk-2"
	kafkaDefault.Spec.Parameters.Size.Plan = "startup-2"
	kafkaDefault.Spec.Parameters.Maintenance.DayOfWeek = "sunday"
	kafkaDefault.Spec.Parameters.Maintenance.TimeOfDay = "00:00:00"

	return &kafkaDefault
}

func (d *Defaults) GetExoscaleMySQLDefault() *exoscalev1.ExoscaleMySQL {
	var mySQLdefault exoscalev1.ExoscaleMySQL

	mySQLdefault.Spec.Parameters.Service.MajorVersion = "8"
	mySQLdefault.Spec.Parameters.Service.Zone = "ch-dk-2"
	mySQLdefault.Spec.Parameters.Size.Plan = "hobbyist-2"
	mySQLdefault.Spec.Parameters.Backup.TimeOfDay = "12:00:00"
	mySQLdefault.Spec.Parameters.Maintenance.DayOfWeek = "sunday"
	mySQLdefault.Spec.Parameters.Maintenance.TimeOfDay = "00:00:00"

	return &mySQLdefault
}

func (d *Defaults) GetExoscaleOpenSearchDefault() *exoscalev1.ExoscaleOpenSearch {
	var openSearchDefault exoscalev1.ExoscaleOpenSearch

	openSearchDefault.Spec.Parameters.Service.MajorVersion = "2"
	openSearchDefault.Spec.Parameters.Service.Zone = "ch-dk-2"
	openSearchDefault.Spec.Parameters.Size.Plan = "hobbyist-2"
	openSearchDefault.Spec.Parameters.Backup.TimeOfDay = "12:00:00"
	openSearchDefault.Spec.Parameters.Maintenance.DayOfWeek = "sunday"
	openSearchDefault.Spec.Parameters.Maintenance.TimeOfDay = "00:00:00"

	return &openSearchDefault
}
