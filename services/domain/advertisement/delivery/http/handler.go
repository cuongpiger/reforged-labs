package http

import (
	lgin "github.com/gin-gonic/gin"
	lzap "go.uber.org/zap"
	lhttp "net/http"

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

		// Response to client
		logger.Info("Created advertisement successfully, response to client")
		dataResponse := &lsdto.CreateAdvertisementResponseDTO{
			AdvertisementID: "123",
			Status:          "queued",
			Priority:        body.Priority,
			CreateAt:        lsutil.Now(),
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

		// Response to client
		logger.Info("Get advertisement successfully, response to client")
		dataResponse := &lsdto.GetAdvertisementResponseDTO{
			AdvertisementID: "123",
			Status:          "completed",
			Analysis: lsdto.AnalysisResponseDTO{
				EffectivenessScore:     0.8,
				Strengths:              []string{"Strong call to action with clear incentive", "Appeals to target audience's desire for progression"},
				ImprovementSuggestions: []string{"Consider adding social proof elements", "Highlight immediate gameplay satisfaction"},
			},
			CreatedAt:   lsutil.Now(),
			CompletedAt: lsutil.Now(),
		}

		pctx.PureJSON(lhttp.StatusOK, lsutil.NewResponse().
			SetStringMessage("Get advertisement successfully").
			SetData(dataResponse).
			GetResponse())
	})
}
