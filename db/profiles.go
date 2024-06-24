package db

type Profile struct {
	ID   int
	Name string
}

func CreateProfile(name string) error {
	stmt, err := DB.Prepare("INSERT INTO profiles(name) VALUES(?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name)
	return err
}

func GetProfile(id int) (*Profile, error) {
	var profile Profile
	err := DB.QueryRow("SELECT id, name FROM profiles WHERE id = ?", id).Scan(&profile.ID, &profile.Name)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func GetProfiles() ([]*Profile, error) {
	var profiles []*Profile
	rows, err := DB.Query("SELECT id, name FROM profiles")
	if err != nil {
		return []*Profile{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var profile Profile
		err := rows.Scan(&profile.ID, &profile.Name)
		if err != nil {
			return []*Profile{}, err
		}
		profiles = append(profiles, &profile)
	}
	return profiles, nil
}
