package core

type Restable interface {
	RestGetter
	RestPoster
	RestPutter
	RestDeleter
}

type RestGetter interface {
	Get() interface{}
}

type RestPoster interface {
	Post(obj interface{}) interface{}
}

type RestPutter interface {
	Put(obj interface{}) interface{}
}

type RestDeleter interface {
	Delete(obj interface{}) interface{}
}
