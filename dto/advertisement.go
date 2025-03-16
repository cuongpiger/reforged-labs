package dto

import (
	lsutil "github.com/cuongpiger/reforged-labs/utils"
)

// CreateAdvertisementRequestDTO ...

type CreateAdvertisementRequestDTO struct {
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Genre          string   `json:"genre"`
	TargetAudience []string `json:"targetAudience"`
	VisualElements []string `json:"visualElements"`
	CallToAction   string   `json:"callToAction"`
	Duration       int      `json:"duration"`
	Priority       int      `json:"priority"`
}

type CreateAdvertisementResponseDTO struct {
	AdvertisementID string           `json:"adId"`
	Status          string           `json:"status"`
	Priority        int              `json:"priority"`
	CreateAt        lsutil.Timestamp `json:"createAt"`
}

// GetAdvertisementRequestDTO ...
type GetAdvertisementRequestDTO struct {
	AdvertisementId string `uri:"ad_id" binding:"required"`
}

type (
	GetAdvertisementResponseDTO struct {
		AdvertisementID string              `json:"adId"`
		Status          string              `json:"status"`
		Analysis        AnalysisResponseDTO `json:"analysis"`
		CreatedAt       lsutil.Timestamp    `json:"createdAt"`
		CompletedAt     *lsutil.Timestamp   `json:"completedAt"`
	}

	AnalysisResponseDTO struct {
		EffectivenessScore     float64  `json:"effectivenessScore"`
		Strengths              []string `json:"strengths"`
		ImprovementSuggestions []string `json:"improvementSuggestions"`
	}
)
