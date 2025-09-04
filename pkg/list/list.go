package list

import "github.com/DiogoJunqueiraGeraldo/jcf/pkg"

type List[T comparable] interface {
	jcf.Coll[T]
	jcf.IterableColl[T]
	jcf.ExtensibleColl[T]
	jcf.RemovableColl[T]
}
