package shop

type Shop struct {
	Id           int64
	Name         string
	ShopTypeId   int64
	ShopUpgrades []ShopUpgrade
	IsAvailable  bool

	Level         int32
	BuildPrice    int64
	HeartPrice    int64
	OperationCost int64

	Stock      int32
	Reputation int32
	Popularity int32

	Floor         int32
	FloorPosition int32
}

type ShopUpgrade struct {
	Id            int64
	ShopId        int64
	Name          string
	IsActive      bool
	HeartPrice    int64
	OperationCost int64
	ReputationAdd int32
	PopularityAdd int32
	StockAdd      int32
}
