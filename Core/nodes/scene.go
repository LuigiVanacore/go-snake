package nodes

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene struct {
	Parent   *Scene
	Children []*Scene
	Nodes    []Node
}

func (s *Scene) Update() {
	for _, node := range s.Nodes {
		node.Update()
	}
}

func (s *Scene) Draw(screen *ebiten.Image) {
	for _, node := range s.Nodes {
		node.Draw(screen)
	}
}

func (s *Scene) AddNode(node *Node) {
	s.Nodes = append(s.Nodes, *node)
}

func (s *Scene) FindNodeByName(name string) (error, *Node) {
	for _, n := range s.Nodes {
		if n.GetName() == name {
			return nil, &n
		}
	}
	return fmt.Errorf("node not found: %s", name), nil
}

func (s *Scene) DeleteNodeByName(name string) error {
	for i, n := range s.Nodes {
		if n.GetName() == name {
			if i != len(s.Nodes)-1 {
				s.Nodes[i] = s.Nodes[len(s.Nodes)-1]
			}
			s.Nodes = s.Nodes[:len(s.Nodes)-1]
			return nil
		}
	}
	return fmt.Errorf("node not found: %s", name)
}

func (s *Scene) AttachChild(child *Scene) {
	child.Parent = s
	s.Children = append(s.Children, child)
}

func (s *Scene) DetachChild(scene *Scene) *Scene {
	found, i := s.FindChildren(scene)
	if found {
		result := s.Children[i]
		result.Parent = nil
		if i != len(s.Children)-1 {
			s.Children[i] = s.Children[len(s.Children)-1]
		}
		s.Children = s.Children[:len(s.Children)-1]
		return result
	}
	return nil
}

func (s *Scene) FindChildren(node *Scene) (bool, int) {
	for i, child := range s.Children {
		if &child == &node {
			return true, i
		}
	}
	return false, 0
}
