package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model

	Name     string `json:"name" gorm:"unique" validate:"required"`
	Location string `json:"location" validate:"required"`
	Field    string `json:"field" validate:"required"`
}

type Job struct {
	gorm.Model
	Company         Company         `json:"-" gorm:"ForeignKey:cid"`
	Cid             uint            `json:"cid"`
	Name            string          `json:"name" validate:"required"`
	Budget          string          `json:"budget" validate:"required"`
	MinNoticePeriod string          `json:"min_notice_period" validate:"required"`
	MaxNoticePeriod string          `json:"max_notice_period" validate:"required"`
	JobLocation     []Location      `json:"job_location" gorm:"many2many:job_location"`
	Technology      []Technology    `json:"technology" gorm:"many2many:job_technology"`
	Description     string          `json:"description" validate:"required"`
	MinExp          string          `json:"min_exp" validate:"required"`
	MaxExp          string          `json:"max_exp" validate:"required"`
	Qualifications  []Qualification `json:"qualification" gorm:"many2many:job_qualification"`
	Shift           []Shift         `json:"shift" gorm:"many2many:job_shift"`
	JobType         string          `json:"job_type" validate:"required"`
}

type Location struct {
	gorm.Model
	Place string `json:"place_name"`
}

type Technology struct {
	gorm.Model
	TechnologyName string `json:"technology_name"`
}

type Qualification struct {
	gorm.Model
	QualificationName string `json:"qualification_name"`
}

type Shift struct {
	gorm.Model
	ShiftType string `json:"shift_type"`
}

type UserApplication struct {
	Name    string  `json:"name"`
	College string  `json:"college"`
	Jid     uint    `json:"jid"`
	Job     NeweJob `json:"newjob"`
}

type NeweJob struct {
	Company         Company         `json:"-" gorm:"ForeignKey:cid"`
	Cid             uint            `json:"cid"`
	Name            string          `json:"name" validate:"required"`
	Budget          string          `json:"budget" validate:"required"`
	MinNoticePeriod string          `json:"min_notice_period" validate:"required"`
	MaxNoticePeriod string          `json:"max_notice_period" validate:"required"`
	JobLocation     []Location      `json:"job_location" gorm:"many2many:job_location"`
	Technology      []Technology    `json:"technology" gorm:"many2many:job_technology"`
	Description     string          `json:"description" validate:"required"`
	MinExp          string          `json:"min_exp" validate:"required"`
	MaxExp          string          `json:"max_exp" validate:"required"`
	Qualifications  []Qualification `json:"qualification" gorm:"many2many:job_qualification"`
	Shift           []Shift         `json:"shift" gorm:"many2many:job_shift"`
	JobType         string          `json:"job_type" validate:"required"`
}
