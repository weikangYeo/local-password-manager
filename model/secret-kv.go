package model

type SecretKV struct {
	Key         string
	Value       string
	Description string
}

func (kv *SecretKV) GetUnmaskedValue(secretKey string) string {

}
