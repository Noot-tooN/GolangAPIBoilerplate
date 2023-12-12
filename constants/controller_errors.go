package constants

import "github.com/gin-gonic/gin"

var (
	InvalidInputError = gin.H{
		"error": "Invalid input provided",
		"message": "Invalid inputs. Please check your inputs",
	}
	CreateError = gin.H{
		"error": "Could not create the requested object",
		"message": "Please check the logs to figure out what went wrong",
	}
	UnauthorizedError = gin.H{
		"error": "UnauthorizedError",
		"message": "Please check your credentials",
	}
	RequestFinished = gin.H{
		"message": "Request has finished successfully",
	}
)