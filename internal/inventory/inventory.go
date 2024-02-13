package inventory

type Inventory struct {
	AccountId  int   `json:"accountId" db:"accountId"`
	ItemIdList []int `json:"temIdList" db:"temIdList"`
}
