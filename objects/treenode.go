package objects

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TokenIterator struct {
	i      int
	tokens []string
}

func (it *TokenIterator) next() string {
	if it.i == len(it.tokens) {
		return "null"
	}
	res := it.tokens[it.i]
	it.i++
	return res
}

func ParseTreeFromStringMust(lcNotation string) *TreeNode {
	root, err := ParseTreeFromString(lcNotation)
	if err != nil {
		panic(err)
	}
	return root
}

func ParseTreeFromString(lcNotation string) (*TreeNode, error) {
	tokens, err := getTreeTokens(lcNotation)
	if err != nil {
		return nil, err
	}
	if len(tokens) == 0 {
		return nil, nil
	}

	it := TokenIterator{0, tokens}
	root, err := parseChild(it.next())
	if err != nil {
		return nil, err
	}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		currNode := nodes[0]
		nodes = nodes[1:]
		if currNode == nil {
			continue
		}

		var lErr, rErr error
		currNode.Left, lErr = parseChild(it.next())
		currNode.Right, rErr = parseChild(it.next())
		if lErr != nil || rErr != nil {
			return nil, errors.Join(lErr, rErr)
		}
		nodes = append(nodes, currNode.Left, currNode.Right)
	}
	return root, nil
}

func getTreeTokens(row string) ([]string, error) {
	row = strings.Trim(row, " ")
	if len(row) < 2 || row[0] != '[' || row[len(row)-1] != ']' {
		return nil, fmt.Errorf("fuck these trees man: %q", row)
	}

	rawTokens := strings.Split(row[1:len(row)-1], ",")
	var tokens []string
	for _, rawToken := range rawTokens {
		tokens = append(tokens, strings.Trim(rawToken, " "))
	}
	return tokens, nil

}

func parseChild(token string) (*TreeNode, error) {
	if token == "null" {
		return nil, nil
	}
	val, err := strconv.Atoi(token)
	if err != nil {
		return nil, err
	}
	return &TreeNode{val, nil, nil}, nil
}
