package dal

func (n *News) mgo() *Mgo {
	return NewMgo("news")
}

// 添加新闻
/*func (n *News) AddNews() {

}*/
