package http

import (
	lfmt "fmt"
	lhttp "net/http"

	lgin "github.com/gin-gonic/gin"
	lzap "go.uber.org/zap"

	lsdto "github.com/cuongpiger/reforged-labs/dto"
	lsadsuc "github.com/cuongpiger/reforged-labs/services/domain/advertisement/usecase"
	lsutil "github.com/cuongpiger/reforged-labs/utils"
)

type advertisementHandler struct {
	advertisementUseCase lsadsuc.IAdvertisementUseCase
}

func NewAdvertisementHandler(padsUc lsadsuc.IAdvertisementUseCase) *advertisementHandler {
	return &advertisementHandler{advertisementUseCase: padsUc}
}

func (s *advertisementHandler) createAdvertisement() lgin.HandlerFunc {
	return lsutil.WithContext(func(pctx *lsutil.Context) {
		var (
			logger = lsutil.GetLogger(pctx)
		)

		logger.Info("Receiving request to create advertisement")
		body := new(lsdto.CreateAdvertisementRequestDTO)
		if err := pctx.ShouldBindJSON(body); err != nil {
			logger.Error("Invalid body request", lzap.Error(err), lzap.Any("body", body))
			pctx.PureJSON(lhttp.StatusBadRequest, lsutil.NewResponse().
				SetStringMessage("Invalid body request").
				SetStatus(lsutil.ResponseStatusFailed).
				GetResponse())
			return
		}

		// Call use case
		result, err := s.advertisementUseCase.CreateAdvertisement(pctx, body)
		if err != nil {
			logger.Error("Create advertisement failed", lzap.Error(err))
			pctx.PureJSON(lhttp.StatusInternalServerError, lsutil.NewResponse().
				SetStringMessage("Create advertisement failed").
				SetStatus(lsutil.ResponseStatusFailed).
				GetResponse())
			return
		}

		// Response to client
		logger.Info("Created advertisement successfully, response to client")
		dataResponse := &lsdto.CreateAdvertisementResponseDTO{
			AdvertisementID: result.Id,
			Status:          result.Status,
			Priority:        result.Priority,
			CreateAt:        lsutil.TimestampFrom(result.CreateAt),
		}

		pctx.PureJSON(lhttp.StatusCreated, lsutil.NewResponse().
			SetData(dataResponse).
			SetStringMessage("Created advertisement successfully").
			GetResponse())
	})
}

func (s *advertisementHandler) getAdvertisement() lgin.HandlerFunc {
	return lsutil.WithContext(func(pctx *lsutil.Context) {
		var (
			logger = lsutil.GetLogger(pctx)
		)

		logger.Info("Receiving request to get advertisement")
		uri := new(lsdto.GetAdvertisementRequestDTO)
		if err := pctx.ShouldBindUri(uri); err != nil {
			logger.Error("Invalid URI request", lzap.Error(err))
			pctx.PureJSON(lhttp.StatusBadRequest, lsutil.NewResponse().
				SetStringMessage("Invalid URI request").
				SetStatus(lsutil.ResponseStatusFailed).
				GetResponse())
			return
		}

		// Call use case
		result, err := s.advertisementUseCase.GetAdvertisement(pctx, uri.AdvertisementId)
		if err != nil {
			logger.Error("Get advertisement failed", lzap.Error(err))
			pctx.PureJSON(lhttp.StatusNotFound, lsutil.NewResponse().
				SetStringMessage(lfmt.Sprintf("Advertisement with ID '%s' not found", uri.AdvertisementId)).
				SetStatus(lsutil.ResponseStatusFailed).
				GetResponse())
			return
		}

		// Response to client
		logger.Info("Get advertisement successfully, response to client")
		dataResponse := &lsdto.GetAdvertisementResponseDTO{
			AdvertisementID: result.Id,
			Status:          result.Status,
			Analysis: lsdto.AnalysisResponseDTO{
				EffectivenessScore:     result.Analysis.EffectivenessScore,
				Strengths:              result.Analysis.Strengths,
				ImprovementSuggestions: result.Analysis.ImprovementSuggestions,
			},
			CreatedAt: lsutil.TimestampFrom(result.CreateAt),
		}

		if result.CompleteAt != nil {
			ts := lsutil.TimestampFrom(*result.CompleteAt)
			dataResponse.CompletedAt = &ts
		}

		pctx.PureJSON(lhttp.StatusOK, lsutil.NewResponse().
			SetStringMessage("Get advertisement successfully").
			SetData(dataResponse).
			GetResponse())
	})
}
