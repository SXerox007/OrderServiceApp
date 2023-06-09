package constants

//status and messages
const (
	MSG_SUCCESS                  = "Success"
	MSG_FAILURE                  = "Error"
	MSG_SOMETING_WENT_WRONG      = "Something Went Wrong"
	ERR_MSG_INTERNAL_SERVER      = "Internal Server Error:"
	ERR_MSG_INVALID_ACCESS_TOKEN = "Invalid Access Token"
	MAX_QUANTITY                 = 10
	DiscountedProductCount       = 3
	DiscountPercentage           = 10
)

// order constant
const (
	ORDER_PLACED     = "Placed"
	ORDER_DISPATCHED = "Dispatched"
	ORDER_COMPLETED  = "Completed"
	ORDER_CANCELLED  = "Cancelled"
)

//product constant
const (
	PREMIUM = "Premium"
	REGULAR = "Regular"
	BUDGET  = "Budget"
)
