package helper

import (
	"github.com/farhaniupr/dating-api/package/library"
	"go.mongodb.org/mongo-driver/bson"
)

type CommonHelper struct {
	env library.Env
}

func ModuleCommonHelper(env library.Env) CommonHelper {
	return CommonHelper{
		env: env,
	}
}

func (u CommonHelper) Helper(id string) {
	// helper
}

func (u CommonHelper) StructToBsonD(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
