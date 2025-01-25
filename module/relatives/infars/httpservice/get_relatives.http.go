package relativeshttpservice

import (
	"github.com/PhuPhuoc/curanest-patient-service/common"
	relativesqueries "github.com/PhuPhuoc/curanest-patient-service/module/relatives/usecase/queries"
	"github.com/gin-gonic/gin"
)

// @Summary		get relatives accounts with filter option
// @Description	get relatives accounts with filter option
// @Tags			relative
// @Accept			json
// @Produce		json
// @Param			create	form		body					relativesqueries.FilterAccountQuery	true	"account creation data"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/api/v1/relatives/filter [post]
// @Security		ApiKeyAuth
func (s *relativesHttpService) handleGetRelativesAccounts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto relativesqueries.FilterAccountQuery
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		result, err := s.query.GetAccWithFilter.Handle(ctx.Request.Context(), &dto)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseGetWithPagination(ctx, result, dto.Paging, dto.Filter)
	}
}
