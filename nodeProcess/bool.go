package nodeProcess

import (
	"github.com/blue-bird1/ConfusedPHP/nodetype"
	"github.com/blue-bird1/ConfusedPHP/util"
	"github.com/z7zmey/php-parser/node"
)

type BoolProcess struct {
	BasePrecess
}

func NewBoolProcess(name string, f func(n node.Node) ([]node.Node, node.Node)) *BoolProcess {
	precess := &BoolProcess{
		BasePrecess: BasePrecess{
			name: name,
			precess: func(n node.Node) (append []node.Node, replace node.Node) {
				return f(n)
			},
		},
	}
	return precess
}

func (b BoolProcess) Check(n node.Node, preNode util.EnterNode) bool {
	return nodetype.IsRetBoolType(n)
}
