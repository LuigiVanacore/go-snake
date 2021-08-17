package Scenes

type SceneNode struct {
	Children []*SceneNode
	Parent   *SceneNode
}

func (s *SceneNode) attachChild(child *SceneNode) {
	child.Parent = s
	s.Children = append(s.Children, child)
}

func (s *SceneNode) detachChild(node *SceneNode) *SceneNode {
	found, i := s.findChildren(node)
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

func (s *SceneNode) findChildren(node *SceneNode) (bool, int) {
	for i, child := range s.Children {
		if &child == &node {
			return true, i
		}
	}
	return false, 0
}
