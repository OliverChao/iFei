package forever

func GetArticleInfo() string {
	cmd := client.HGet("user:1:info", "age")
	s, e := cmd.Result()
	//return fmt.Sprintf("%v",cmd.Result())
	if e != nil {
		return ""
	}
	return s
}

//func SaveLoginState()
