package chaincode

type VoteData struct {
	User User `yaml:"User" json:"User"`
	Vote string `yaml:"Vote" json:"Vote"`
	Gateway Gateway `yaml:"Gateway" json:"Gateway"`
}

type User struct {
	ID string `yaml:"UserId" json:"UserId"`
	FirstName        string   `yaml:"FirstName" json:"FirstName"`
	LastName         string   `yaml:"LastName" json:"LastName"`
	DriversLicenseID string   `yaml:"DriversLicenseID" json:"DriversLicenseID"`
	ImageURIs        []string `yaml:"ImageURIs" json:"ImageURIs"`
	VideoURIs        []string `yaml:"VideoURIs" json:"VideoURIs"`
	SoundURIs        []string `yaml:"SoundURIs" json:"SoundURIs"`
}

type Gateway struct {
	ID       string `yaml:"GatewayId" json:"GatewayId"`
	Location string `yaml:"Location" json:"Location"`
}
