package redis

//var c *redis.Client
//
//func init() {
//	client := redis.NewClient(&redis.Options{
//		Addr:     config.C().Redis.Addr(),
//		Password: config.C().Redis.Pass,
//	})
//	if err := client.Ping().Err(); err != nil {
//		logrus.Fatal(err)
//	}
//	c = client
//}
//
//func CLI() *redis.Client {
//	return c
//}
