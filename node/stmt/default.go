package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Default struct {
	position *node.Position
	comments *[]comment.Comment
	Stmts    []node.Node
}

func NewDefault(Stmts []node.Node) node.Node {
	return &Default{
		nil,
		nil,
		Stmts,
	}
}

func (n Default) Attributes() map[string]interface{} {
	return nil
}

func (n Default) Position() *node.Position {
	return n.position
}

func (n Default) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Default) Comments() *[]comment.Comment {
	return n.comments
}

func (n Default) SetComments(c []comment.Comment) node.Node {
	n.comments = &c
	return n
}

func (n Default) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
