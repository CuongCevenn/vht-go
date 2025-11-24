package orderrpcclient

type FoodRPCClient struct {
	foodServiceURI string
}

func NewFoodRPCClient(foodServiceURI string) *FoodRPCClient {
	return &FoodRPCClient{foodServiceURI: foodServiceURI}
}

type UserRPCClient struct {
	userServiceURI string
}

func NewUserRPCClient(userServiceURI string) *UserRPCClient {
	return &UserRPCClient{userServiceURI: userServiceURI}
}

