package context

// CreateUser creates new user and stores into database
func (conn *Database) CreateUser(u Users) (nUser Users, err error) {
	if err = conn.DB.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// Migrate creates the schema of the tables
func (conn *Database) Migrate() {
	conn.DB.AutoMigrate(&Users{}, &Addresses{}, &UserAddresses{})

}

// GetUser gives us the user id by specifying username
func (conn *Database) GetUser(username string) (user Users, err error) {

	if err = conn.DB.Where("username = ? or email = ?", username, username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// GetAddress gives us default address of the user specifying username
func (conn *Database) GetAddress(username string) (address Addresses, err error) {
	user, err := conn.GetUser(username)
	if err != nil {
		return address, err
	}
	userAddresses := &UserAddresses{}
	if err = conn.DB.Where("user_id = ? and is_default='YES'", user.ID).First(userAddresses).Error; err != nil {
		return address, err
	}

	if err := conn.DB.Where("id = ?", userAddresses.AddressID).First(&address).Error; err != nil {
		return address, err
	}
	return address, nil
}

// SetAddress insert address for the user by specifying username
func (conn *Database) SetAddress(username string, address Addresses, isDefault string) (newAddress Addresses, err error) {
	user, err := conn.GetUser(username)
	if err != nil {
		return newAddress, err
	}

	if err := conn.DB.Create(&address).Error; err != nil {
		return newAddress, err
	}

	err = conn.DB.Create(&UserAddresses{
		UserID:    user.ID,
		AddressID: address.ID,
		IsDefault: isDefault,
	}).Error

	if err != nil {
		return address, err
	}

	return address, nil

}

// SetDefaultAddress set the default address for the user
func (conn *Database) SetDefaultAddress(id uint64, username string) (string, error) {
	user, err := conn.GetUser(username)
	if err != nil {
		return "failed", err
	}
	if err := conn.DB.Model(&UserAddresses{UserID: user.ID}).Update("is_default", "NO").Error; err != nil {
		return "failed", err
	}
	if err := conn.DB.Model(&UserAddresses{AddressID: id}).Update("is_default", "YES").Error; err != nil {
		return "failed", err
	}

	return "success", nil
}
