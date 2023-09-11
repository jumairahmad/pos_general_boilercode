package Routers

import (
	Controllers "POS/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("book", Controllers.ListBook)
		v1.POST("book", Controllers.AddNewBook)
		v1.GET("book/:id", Controllers.GetOneBook)
		v1.PUT("book/:id", Controllers.PutOneBook)
		v1.DELETE("book/:id", Controllers.DeleteBook)

		v1.GET("order", Controllers.ListOrders)
		v1.POST("order", Controllers.AddNewOrders)
		v1.GET("order/:id", Controllers.GetOneOrders)
		v1.PUT("order/:id", Controllers.PutOneOrders)
		v1.DELETE("order/:id", Controllers.DeleteOrders)

		v1.GET("product", Controllers.ListProduct)
		v1.POST("product", Controllers.AddNewProduct)
		v1.GET("product/:id", Controllers.GetOneProduct)
		v1.PUT("product/:id", Controllers.PutOneProduct)
		v1.DELETE("product/:id", Controllers.DeleteProduct)

		v1.GET("quantity", Controllers.ListQuantity)
		v1.POST("quantity", Controllers.AddNewQuantity)
		v1.GET("quantity/:id", Controllers.GetOneQuantity)
		v1.PUT("quantity/:id", Controllers.PutOneQuantity)
		v1.DELETE("quantity/:id", Controllers.DeleteQuantity)

		v1.GET("roles", Controllers.ListRoles)
		v1.POST("roles", Controllers.AddNewRoles)
		v1.GET("roles/:id", Controllers.GetOneRoles)
		v1.PUT("roles/:id", Controllers.PutOneRoles)
		v1.DELETE("roles/:id", Controllers.DeleteRoles)

		v1.GET("stock", Controllers.ListStock)
		v1.POST("stock", Controllers.AddNewStock)
		v1.GET("stock/:id", Controllers.GetOneStock)
		v1.PUT("stock/:id", Controllers.PutOneStock)
		v1.DELETE("stock/:id", Controllers.DeleteStock)

		v1.GET("supplier", Controllers.ListSupplier)
		v1.POST("supplier", Controllers.AddNewSupplier)
		v1.GET("supplier/:id", Controllers.GetOneSupplier)
		v1.PUT("supplier/:id", Controllers.PutOneSupplier)
		v1.DELETE("supplier/:id", Controllers.DeleteSupplier)

		v1.GET("user", Controllers.ListUser)
		v1.POST("user", Controllers.AddNewUser)
		v1.GET("user/:id", Controllers.GetOneUser)
		v1.PUT("user/:id", Controllers.PutOneUser)
		v1.DELETE("user/:id", Controllers.DeleteUser)

		v1.GET("warehouse", Controllers.ListWarehouse)
		v1.POST("warehouse", Controllers.AddNewWarehouse)
		v1.GET("warehouse/:id", Controllers.GetOneWarehouse)
		v1.PUT("warehouse/:id", Controllers.PutOneWarehouse)
		v1.DELETE("warehouse/:id", Controllers.DeleteWarehouse)

	}

	return r
}
