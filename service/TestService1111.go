package service

//
//func Method1(c *gin.Context) {
//	var issues = make([]dto.IssueInfo, 0)
//	Issues := models.Select()
//	for _, ishida := range Issues {
//		issue := dto.IssueInfo{
//			Id:      ishida.ID,
//			Key:     ishida.Key,
//			ImageID: int(ishida.IssueID),
//		}
//		issues = append(issues, issue)
//	}
//	jsondata, _ := json.Marshal(issues)
//	token := token(c)
//	utils.Red.Set(c, "Result:"+token, jsondata, 12*time.Hour)
//	utils.RespOk(c.Writer, Issues, "返回四条石田测试题")
//}
//
//func GetColor(c *gin.Context) {
//	Issue := models.SearchColor()
//	Issue.Key = rands()
//	token := token(c)
//	id, _ := c.Get("userInfoId")
//	userstr, _ := utils.Red.Get(c, "UserInfo:"+string(id.(int))+":").Result()
//	//var height = 0
//	user := dto.UserInfo{}
//	//if userstr != ""{
//	json.Unmarshal([]byte(userstr), &user)
//	//	height = user.Heightest
//	//}
//	//user := dto.UserInfo{
//	//	ID:        id.(int),
//	//	Blood:     3,
//	//	Heightest: 0,
//	//}
//	user.ID = id.(int)
//	user.Blood = 3
//	user.Score = 0
//	userJson, _ := json.Marshal(user)
//	utils.Red.Set(c, "UserInfo:"+string(user.ID)+":", userJson, -1)
//	utils.Red.Set(c, token, Issue.Key, 12*time.Hour)
//	utils.RespOk(c.Writer, Issue, "返回两张相似色调色块")
//}
//
//func Judge_c(c *gin.Context) {
//	token := token(c)
//	key := c.Query("key")
//	id, _ := c.Get("userInfoId")
//	userstr, _ := utils.Red.Get(c, "UserInfo:"+string(id.(int))+":").Result()
//	user := dto.UserInfo{}
//	json.Unmarshal([]byte(userstr), &user)
//	if user.Blood > 0 {
//		cacheKey, err := utils.Red.Get(c, token).Result()
//		if err != nil {
//			utils.RespFail(c.Writer, "redis获取键值对错误。")
//			return
//		}
//		if key != cacheKey {
//			utils.RespFail(c.Writer, "本题回答错误。")
//			user.Blood--
//			userJson, _ := json.Marshal(user)
//			utils.Red.Set(c, "UserInfo:"+string(user.ID)+":", userJson, -1)
//			return
//
//		} else {
//			utils.RespOk(c.Writer, nil, "本题回答正确。")
//			user.Score++
//			if user.Score > user.Heightest {
//				user.Heightest = user.Score
//			}
//			userJson, _ := json.Marshal(user)
//			utils.Red.Set(c, "UserInfo:"+string(user.ID)+":", userJson, -1)
//			return
//		}
//	}
//	utils.RespFail(c.Writer, "血量不足！")
//	return
//}
//
//// 将Issue的list存入redis并从redis去出查看答案
//func Judge_m(c *gin.Context) {
//	options := c.Request.FormValue("options")
//	Issues, _ := utils.Red.Get(c, "Result:"+token(c)).Result()
//	var str string
//	var issuesCache = make([]dto.IssueInfo, 0)
//	var results = make([]dto.ResultInfo, 0)
//	var rets = make([]int, 3)
//	json.Unmarshal([]byte(Issues), &issuesCache)
//	cnt := 0
//	for i, issueCache := range issuesCache {
//		if i < len(options) {
//			image := dto.SeachImage(issueCache.ImageID)
//			if issueCache.Key == string(options[i]) {
//				result := dto.ResultInfo{
//					Key:   issueCache.Key,
//					Mykey: string(options[i]),
//					Point: Point(image.C_type),
//					Image: image.Image,
//					Flag:  true,
//				}
//				results = append(results, result)
//				cnt++
//				//str += "第" + strconv.Itoa(i+1) + "题回答正确\n"
//			} else {
//				result := dto.ResultInfo{
//					Key:   issueCache.Key,
//					Mykey: string(options[i]),
//					Point: Point(image.C_type),
//					Image: image.Image,
//					Flag:  false,
//				}
//				rets[image.C_type-1]++
//				results = append(results, result)
//				//str += "第" + strconv.Ita(i+1) + "题回答错误\n"
//			}
//		}
//	}
//	str = ret(rets)
//	id, _ := c.Get("userInfoId")
//	userJson,_ := utils.Red.Get(c,"user:"+token(c)).Result()
//	user := dto.UserInfo{}
//	json.Unmarshal([]byte(userJson),user)
//	if user == nil {
//		user.ID =
//	}
//	utils.Red.Set(c,"user:"+token(c),)
//	utils.RespOk(c.Writer, results, "共有4道题，回答正确"+strconv.Itoa(cnt)+"道题;"+str)
//	return
//}
//
//// 返回结果？？？
//func Point(t int) string {
//	switch t {
//	case 1:
//		{
//			return "分析：红色色盲"
//		}
//	case 2:
//		{
//			return "分析：红色色盲"
//		}
//	case 3:
//		{
//			return "分析：蓝紫色盲"
//		}
//	}
//	return ""
//}
//func ret(t []int) string {
//	var str string
//	if t[0] != 0 {
//		str += "有一定程度红色认知困难 "
//	}
//	if t[1] != 0 {
//		str += "有一定程度绿色认知困难 "
//	}
//	if t[2] != 0 {
//		str += "有一定程度蓝紫色认知困难 "
//	}
//	return str
//}
