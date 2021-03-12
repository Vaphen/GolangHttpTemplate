package profile

type ProfileService struct { }

func NewService() *ProfileService {
	return &ProfileService{}
}

func (p ProfileService) Get() interface{} {
	return struct{A string}{"Get"}
}

func (p ProfileService) Post(obj interface{}) interface{} {
	return struct{A string}{"Post"}
}

func (p ProfileService) Delete(obj interface{}) interface{} {
	return struct{A string}{"Delete"}
}

func (p ProfileService) Put(obj interface{}) interface{} {
	return struct{A string}{"Put"}
}
