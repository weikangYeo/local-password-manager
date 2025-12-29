package model

type Profile struct {
	Name  string `json:"name"`
	Hints string `json:"hints"`
	// when app start, check if app key is generated. If not gen.
	// todo: Later key to sign secret key and store.
	// this signature would be used to verify secret key.
	KeySignature string `json:"keySignature"`
}
