package controllers

import (
	postgresqlclient "golangapi/databases/postgre_sql_client"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckPostgre(ctx *gin.Context) {
	pgclient := postgresqlclient.GetDefaultPostgreClient()

	err := pgclient.Ping()

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "Postgre pinged")
}

func CheckServer(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Ping Pong")
}
