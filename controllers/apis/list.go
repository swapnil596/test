package apis

//
//import (
//	"api-registration-backend/common"
//	"api-registration-backend/config"
//	"github.com/gin-gonic/gin"
//	"log"
//	"net/http"
//	"strconv"
//)
//
//type CatalogueController struct {
//
//}
//
//func (c CatalogueController) List(ctx *gin.Context)  {
//
//	// Query string parameters are parsed using the existing underlying request object
//	// The request responds to a url matching /workflow/dags?start=<number>&end=<number>
//	start := ctx.DefaultQuery("start", "0")
//	end := ctx.DefaultQuery("end", "10")
//
//	base, err := strconv.ParseInt(start, 10, 64)
//	if err != nil {
//		log.Println(err.Error())
//		common.FailResponse(ctx, http.StatusBadRequest, "Error",
//			gin.H{"errors": "Invalid data passed in `start`"})
//		return
//	}
//
//	limit, err := strconv.ParseInt(end, 10, 64)
//	if err != nil {
//		log.Println(err.Error())
//		common.FailResponse(ctx, http.StatusBadRequest, "Error",
//			gin.H{"errors": "Invalid data passed in `end`"})
//		return
//	}
//
//	if limit - base > 10 {
//		log.Println("Range greater than ")
//		common.FailResponse(ctx, http.StatusNotAcceptable, "Error",
//			gin.H{"errors": ""})
//		return
//	}
//
//	// fetch required configurations
//	conf := config.GetConfigurations()
//
//
//}
