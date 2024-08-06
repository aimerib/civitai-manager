package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DBTransactionMiddleware wraps the request in a database transaction
func DBTransactionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		txHandle := db.Begin()
		c.Set("db_trx", txHandle)
		c.Set("db", txHandle) // For backward compatibility if you're using "db" key elsewhere

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		c.Next()

		if len(c.Errors) > 0 {
			txHandle.Rollback()
		} else {
			if err := txHandle.Commit().Error; err != nil {
				txHandle.Rollback()
			}
		}
	}
}

// GetTrx retrieves the current transaction from the Gin context
func GetTrx(c *gin.Context) *gorm.DB {
	return c.MustGet("db_trx").(*gorm.DB)
}
