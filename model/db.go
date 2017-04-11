package model

// DB is the interface to the database.
// The idea here is to have multiple different DB implementations.
type DB interface {
	// Get gets the value for the given stock code.
	// returns the value and nil if it was found, nil and error if it
	// was not found and another appropriate error otherwise
	Get(code string) (kd, error)

	// Set sets the value for the given stock data. returns stock code if the data
	// was successfully set, error if the data already exist, and another
	// appropriate error otherwise
	Set(data kd) (string, error)

	// Save sets or creates the data. returns true and nil if the key was created
	// on this call, false and nil if the key was not created but still successfully
	// updated, and false and the appropriate error otherwise
	Save(data kd) (bool, error)
}
