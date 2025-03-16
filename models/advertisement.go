package models

import ltime "time"

type Advertisement struct {
	Id         string      `json:"id" gorm:"primaryKey"`
	Status     string      `json:"status"`
	Priority   int         `json:"priority"`
	Analysis   Analysis    `json:"analysis" gorm:"serializer:json"` // Store as JSON in DB
	CreateAt   ltime.Time  `json:"createAt"`
	CompleteAt *ltime.Time `json:"completeAt"`
}

type Analysis struct {
	EffectivenessScore     float64  `json:"effectivenessScore"`
	Strengths              []string `json:"strengths"`
	ImprovementSuggestions []string `json:"improvementSuggestions"`
}
