package categories

import (
	"github.com/shion0625/grpc/msa/internal/pkg/errs"

	"github.com/google/uuid"
)

// 商品カテゴリを表すEntity
type Category struct {
	id   *categoryId   // カテゴリ番号
	name *categoryName // カテゴリ名
}

// コンストラクタ
func NewCategory(name *categoryName) (*Category, *errs.DomainError) {
	if uid, err := uuid.NewRandom(); err != nil { // UUIDを生成する
		return nil, errs.NewDomainError(err.Error())
	} else {
		if id, err := NewCategoryId(uid.String()); err != nil {
			return nil, errs.NewDomainError(err.Error())
		} else {
			return &Category{
				id:   id,
				name: name,
			}, nil
		}
	}
}

// ゲッター
func (c *Category) Id() *categoryId {
	return c.id
}
func (c *Category) Name() *categoryName {
	return c.name
}

// 同一性検証
func (c *Category) Equals(obj *Category) (bool, *errs.DomainError) {
	if obj == nil {
		return false, errs.NewDomainError("引数でnilが指定されました。")
	}
	result := c.id.Equals(obj.Id())
	return result, nil
}
